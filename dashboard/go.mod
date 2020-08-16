module sigs.k8s.io/cluster-addons/dashboard

go 1.13

require (
	github.com/docker/docker v0.7.3-0.20190327010347-be7ac8be2ae0
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.10.1
	github.com/onsi/gomega v1.7.0
	k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
	k8s.io/klog v1.0.0
	sigs.k8s.io/controller-runtime v0.4.0
	sigs.k8s.io/kubebuilder-declarative-pattern v0.0.0-20200226054827-748a6481b2a4
	sigs.k8s.io/kubebuilder-intercept v0.0.0-00010101000000-000000000000
)

replace sigs.k8s.io/kubebuilder-intercept => ../../kubebuilder-intercept
