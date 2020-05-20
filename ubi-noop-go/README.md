# UBI NoOp Go Operator
### Example Use Only
This is a simple operator to be used as a development example only. It is, however a working example that deploys the desired number of pods, which each sleep for 1 hour (hence no-op).

To test, create a Custom Resource of `kind: UBINoOp` with `spec.size` set to the desired number of pods.

The pods' image can be customized by setting the `RELATED_IMAGE_UBI_MINIMAL` environment variable, which should contain the full path to the desired image/tag.

This operator sets owner references on all pods for garbage collection, and also updates the CR status with a simple boolean value upon successful deployment (`status.deployed`).

For quick links to example code and associated OLM metadata bundle, go to https://github.com/jsm84/rhm-env-example.

### Known bugs and lacking feature examples

See above for the "academic use" disclaimer, but the operator currently only handles scale-out (increase number of pods) but not scale-in (terminating excess pods if `cr.spec.size` is set to a lesser number of pods).
Expect inverse scaling in a later version.

