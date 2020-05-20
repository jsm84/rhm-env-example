# Red Hat Marketplace - Environment Variable Example Operators

This repo contains example operator code and metadata (for deployment using the Operator Lifecycle Manager or OLM) used to showcase the override of an operator image source using environment variables.
Refactoring your operator code to use env vars for the operand(s) image source is not just to ensure that your operator will work in disconnected/offline or proxied OpenShift environments.
This is a hard technical requirement that is used to identify all images associated with your operator, so that they can be mirrored to the marketplace registry.

## Required Environment Variables for RHM

For compatibility with the Red Hat Marketplace, the following naming convention is required:

* All env vars should start with the prefix `RELATED_IMAGE_`
* The convention follows the rough form of `RELATED_IMAGE_<IDENTIFIER>`
* Any friendly name can be used for the latter part of the environment variable, however any alpha characters should be upper case (eg: `RELATED_IMAGE_DATABASE`) per bourne/bash shell conventions
* Alphanumeric or purely numeric identifiers can also be used (eg: `RELATED_IMAGE_DB0` or simply `RELATED_IMAGE_0`)
* Numerous environment variables can be defined and utilized as needed to allow any number of operand types (unique container images)
* Extended identifiers can be used by including additional underscores in the identifier (eg: `RELATED_IMAGE_ISTIO_SIDECAR`)

## Example Implementation - Quick Links

See [ubinoop_controller.go](https://github.com/jsm84/rhm-env-example/blob/master/ubi-noop-go/pkg/controller/ubinoop/ubinoop_controller.go#L139-L176) for an example go source implementation.

Also see [clusterserviceversion.yaml](https://github.com/jsm84/rhm-env-example/blob/master/ubi-noop-go/deploy/olm-catalog/ubi-noop-go/0.0.1/ubi-noop-go.v0.0.1.clusterserviceversion.yaml#L86-L88) for an example of setting these environment variables in the operator metadata.
