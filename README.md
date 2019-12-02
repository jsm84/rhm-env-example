# Open Marketplace - Environment Variable Example Operators

This repo contains example operator code and metadata (for deployment using the Operator Lifecycle Manager or OLM) used to showcase the override of an operator image source using environment variables.
Refactoring your operator code to use env vars for the operand(s) image source will help ensure that your operator will work in disconnected/offline or proxied Kubernetes environments such as OpenShift Container Platform (when installed in such an environment).
You'll see the term "air gapped" used elsewhere to refer to such environments, but that term is misleading so it's not used here.

## Proposed Naming Convention for Environment Variables

For compatibility with OpenShift Container Platform and the K8s platform-agnostic Operator Lifecycle Manager, the following convention is suggested when naming your environment variables:

* All env vars should start with the prefix `RELATED_IMAGE_`
* The convention follows the rough form of `RELATED_IMAGE_<identifier>`
* Any friendly name can be used for the latter part of the environment variable, however any alpha characters should be upper case (eg: `RELATED_IMAGE_DATABASE`) per bourne/bash shell conventions
* Alphanumeric or purely numeric identifiers can also be used (eg: `RELATED_IMAGE_DB0` or simply `RELATED_IMAGE_0`)
* Numerous environment variables can be defined and utilized as needed to allow any number of operand types (unique container images)
* Extended identifiers can be used by including additional underscores in the identifier (eg: `RELATED_IMAGE_ISTIO_SIDECAR`)

## Example Implementation - Quick Links

See <TODO: link source for controller.go> for an example go source implementation.

Also see <TODO: link to ClusterServiceVersion> for an example of setting these environment variables in the operator metadata.
