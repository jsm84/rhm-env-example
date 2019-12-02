package ubinoop

import (
	"context"
	noopv1alpha1 "github.com/jsm84/om-env-example/pkg/apis/noop/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"strconv"
)

var log = logf.Log.WithName("controller_ubinoop")

// Add creates a new UBINoOp Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileUBINoOp{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("ubinoop-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource UBINoOp
	err = c.Watch(&source.Kind{Type: &noopv1alpha1.UBINoOp{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Pods and requeue the owner UBINoOp
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &noopv1alpha1.UBINoOp{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileUBINoOp implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileUBINoOp{}

// ReconcileUBINoOp reconciles a UBINoOp object
type ReconcileUBINoOp struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a UBINoOp object and makes changes based on the state read
// and what is in the UBINoOp.Spec
// This example creates a desired number of Pods according to cr.Spec.Size
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileUBINoOp) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling UBINoOp")

	// Fetch the UBINoOp instance
	instance := &noopv1alpha1.UBINoOp{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define a new Pod object
	pods := newPodsForCR(instance)

	// Check if each Pod already exists
	found := &corev1.Pod{}

	for _, pod := range pods {
		// Set UBINoOp instance as the owner and controller
		if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
			return reconcile.Result{}, err
		}
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
		if err != nil && errors.IsNotFound(err) {
			reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
			err = r.client.Create(context.TODO(), pod)
			if err != nil {
				return reconcile.Result{}, err
			}

			// Pod created successfully - don't requeue
			return reconcile.Result{}, nil
		} else if err != nil {
			return reconcile.Result{}, err
		}

		// Pod already exists - don't requeue
		reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	}

	// Update CR Status once deployed
	if instance.Status.Deployed != true {
		instance.Status.Deployed = true
		r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update pod deploy status")
			return reconcile.Result{}, err
		}
	}
	return reconcile.Result{}, nil
}

// newPodsForCR returns a ubi-minimal pod with the same name/namespace as the cr
func newPodsForCR(cr *noopv1alpha1.UBINoOp) []*corev1.Pod {
	// get the container image location from the expected env var
	// default to ubi8 minimal image from registry.redhat.io
	ubiImg := os.Getenv("RELATED_IMAGE_UBI_MINIMAL")
	if ubiImg == "" {
		ubiImg = "registry.redhat.io/ubi8/ubi-minimal:latest"
	}

	// podLst is a slice of pod whose len matches cr.Spec.Size
	podLst := make([]*corev1.Pod, cr.Spec.Size)

	labels := map[string]string{
		"app": cr.Name,
	}

	for idx, _ := range podLst {
		i := strconv.Itoa(idx)
		podLst[idx] = &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      cr.Name + "-pod" + i,
				Namespace: cr.Namespace,
				Labels:    labels,
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:    "ubi-minimal",
						Image:   ubiImg,
						Command: []string{"sleep", "3600"},
					},
				},
			},
		}
	}

	return podLst
}
