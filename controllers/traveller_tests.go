// +kubebuilder:docs-gen:collapse=Apache License
package controllers

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	routev1 "github.com/openshift/api/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"

	mydomainv1alpha1 "my.domain/hello-operator2/api/v1alpha1"
)

// +kubebuilder:docs-gen:collapse=Imports
var _ = Describe("Arcade Controller", func() {
	ctx := context.Background()
	By("Verify deployment was created")
	dep := &appsv1.Deployment{}
	Eventually(func() bool {
		err := k8sClient.Get(ctx, key, dep)
		if err != nil {
			return false
		}
		return true
	}, timeout, interval).Should(BeTrue())
	Expect(dep.Spec.Template.Spec.Containers[0].Name).To(Equal("arcade"))
	Expect(dep.Spec.Template.Spec.Containers[0].Image).To(ContainSubstring("arcade"))

})
