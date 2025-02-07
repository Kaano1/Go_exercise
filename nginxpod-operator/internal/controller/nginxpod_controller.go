/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	nginxpodcomv1 "nginxpod/api/v1"
)

// NginxPodReconciler reconciles a NginxPod object
type NginxPodReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

// +kubebuilder:rbac:groups=nginxpod.com,resources=nginxpods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=nginxpod.com,resources=nginxpods/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=nginxpod.com,resources=nginxpods/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the NginxPod object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.0/pkg/reconcile
func (r *NginxPodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	ctx = context.Background()
	Log := r.Log.WithValues("nginxpod", req.NamespacedName)

	var nginxPod nginxpodcomv1.NginxPod
	if err := r.Get(ctx, req.NamespacedName, &nginxPod); err != nil {
		Log.Error(err, "unable to fetch NginxPod")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	podSpec, err := r.createNginx(nginxPod)
	if err != nil {
		Log.Error(err, "unable to create Nginx Pod")
		return ctrl.Result{}, err
	}

	if err := r.Create(ctx, podSpec); err != nil {
		Log.Error(err, "unable to create Nginx Pod")
		return ctrl.Result{}, err
	}
	nodePort, err := r.createService(nginxPod)
	if err != nil {
		Log.Error(err, "unable to create Nginx Service")
		return ctrl.Result{}, err
	}

	if err := r.Create(ctx, nodePort); err != nil {
		Log.Error(err, "unable to create Nginx Service")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *NginxPodReconciler) createService(nginxpod nginxpodcomv1.NginxPod) (*corev1.Service, error) {
	// Create a new Service

	service := corev1.Service{
		TypeMeta:   metav1.TypeMeta{APIVersion: corev1.SchemeGroupVersion.String(), Kind: "Service"},
		ObjectMeta: metav1.ObjectMeta{Name: nginxpod.Spec.ContainerName + "-service", Namespace: nginxpod.Namespace},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": nginxpod.Spec.ContainerName},
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					Port:       80,
					TargetPort: intstr.FromInt(80), // why don't define string in front of TargetPort? because it is an intstr.IntOrString type and it can be either an int or a string.
				},
			},
			Type: corev1.ServiceTypeNodePort,
		},
	}
	return &service, nil
}

func (r *NginxPodReconciler) createNginx(nginxpod nginxpodcomv1.NginxPod) (*corev1.Pod, error) {
	// Create a new Pod
	image := "nginx:latest"

	nginxPod := corev1.Pod{
		TypeMeta:   metav1.TypeMeta{APIVersion: corev1.SchemeGroupVersion.String(), Kind: "Pod"},
		ObjectMeta: metav1.ObjectMeta{Name: nginxpod.Spec.ContainerName + "-pod", Namespace: nginxpod.Namespace, Labels: map[string]string{"app": nginxpod.Spec.ContainerName}},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{ // we need to copilot support for this part
				{
					Name:  nginxpod.Spec.ContainerName,
					Image: image,
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 80,
							Protocol:      corev1.ProtocolTCP,
						},
					},
				},
			},
		},
	}

	return &nginxPod, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NginxPodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&nginxpodcomv1.NginxPod{}).
		Named("nginxpod").
		Complete(r)
}
