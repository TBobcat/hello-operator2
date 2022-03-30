// +kubebuilder:docs-gen:collapse=Apache License
package controllers

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "k8s.io/api/apps/v1"

	// corev1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//intstr "k8s.io/apimachinery/pkg/util/intstr"

	"k8s.io/apimachinery/pkg/types"
	//ctrl "sigs.k8s.io/controller-runtime"
	//mydomainv1alpha1 "my.domain/hello-operator2/api/v1alpha1"
)

var (
	cfg       *rest.Config
	k8sClient client.Client
)

// +kubebuilder:docs-gen:collapse=Imports
var _ = Describe("Arcade Controller", func() {
	const (
		ArcadeName            = "test-traveller"
		ArcadeNamespace       = "default"
		ArcadePort      int32 = 8080
		timeout               = time.Second * 10
		interval              = time.Millisecond * 250
	)

	ctx := context.Background()
	var err error

	// Create client for testing
	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).ToNot(HaveOccurred())
	Expect(k8sClient).ToNot(BeNil())

	key := types.NamespacedName{Name: ArcadeName, Namespace: ArcadeNamespace}
	//createdArcade := &mydomainv1alpha1.Traveller{}

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
