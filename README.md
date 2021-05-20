## Hello, OpenShift for Developers! ##
This is a example that will serve an HTTP response of "Hello OpenShift for Developers!" written in Golang. It is also
intended to be build and used with the [Golang Source-to-Image builder image](https://github.com/sclorg/golang-container).  This builder image is included in current versions of OpenShift (4.7 at the time).

### Deploying on OpenShift
You can deploy this application on OpenShift using the following command:

```shell
oc new-app golang~https://github.com/openshift-for-developers/hello.git
```

Note: this example app is derived from the [example at sclorg/golang-ex](https://github.com/sclorg/golang-ex).

The response message can be set by using the RESPONSE environment
variable.  You will need to edit the pod definition and add an
environment variable to the container definition and run the new pod.
