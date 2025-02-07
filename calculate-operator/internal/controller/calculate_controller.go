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
	"encoding/base64"
	"fmt"

	culatecomv1 "culate.com/api/v1"
	"github.com/go-logr/logr"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strconv"
)

// CalculateReconciler reconciles a Calculate object
type CalculateReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

// +kubebuilder:rbac:groups=culate.com,resources=calculates,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=culate.com,resources=calculates/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=culate.com,resources=calculates/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Calculate object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.0/pkg/reconcile
func (r *CalculateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	ctx = context.Background()                               // ctx ensures that controllers are created with a background context before being passed to the reconciler
	log := r.Log.WithValues("calculate", req.NamespacedName) // r.Log.WithValues() returns a new Logger with the given values appended to those in the Logger's Values in this way we can add the name of the Calculate object to the log

	var calculate culatecomv1.Calculate                                // we can use this to get the Calculate object from the request
	if err := r.Get(ctx, req.NamespacedName, &calculate); err != nil { // r.Get() is used to get the Calculate object from the request and store it in the calculate variable. we use ctx because it is a background context and req.NamespacedName is the name of the Calculate object
		log.Error(err, "unable to fetch Calculate") // if we are unable to get the Calculate object we log an error
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	jobSpec, err := r.createJob(calculate) // we can use this to create a job from the Calculate object
	if err != nil {
		log.Error(err, "unable to create job")
		return ctrl.Result{}, err
	}

	if err := r.Create(ctx, jobSpec); err != nil { //r.Create() is used to create a job from the jobSpec
		log.Error(err, "unable to create job")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *CalculateReconciler) createJob(calculate culatecomv1.Calculate) (*batchv1.Job, error) { //batchv1.Job is the type of the jobSpec that we will return from this function
	// we can use this to create a job from the Calculate object
	image := "knsit/pandoc"                                                                                          // this image is used to convert the base64 text to markdown
	base64text := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(calculate.Spec.Num1 + calculate.Spec.Num2))) // base64text has the base64 encoded text of the Calculate object

	// corev1 is the core api group, it contains the core types of kubernetes like pods, services, etc.
	// batchv1 is the batch api group, it contains the batch types of kubernetes like jobs, cronjobs, etc.
	// metav1 is the meta api group, it contains the meta types of kubernetes like objectmeta, listmeta, etc.
	j := batchv1.Job{ // we can use this to create a job from the Calculate object
		TypeMeta: metav1.TypeMeta{APIVersion: batchv1.SchemeGroupVersion.String(), Kind: "Job"}, // SchemeGroupVersion.String() returns the group version as a string if you want to create a pod we should define core.SchemeGroupVersion.String()
		ObjectMeta: metav1.ObjectMeta{
			Name:      calculate.Name + "-job",
			Namespace: calculate.Namespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{ // template is a pod template that will be used to create the pods
				Spec: corev1.PodSpec{
					RestartPolicy: corev1.RestartPolicyOnFailure, // RestartPolicyOnFailure means that the pod will be restarted if it fails
					InitContainers: []corev1.Container{ // initContainers are containers that are run before the main container
						{
							Name:    "store-result",
							Image:   "alpine",
							Command: []string{"/bin/sh"},                                                                // the command that will be run in the container
							Args:    []string{"-c", fmt.Sprintf("echo %s | base64 -d >> /data/result.txt", base64text)}, // the args that will be passed to the command and %s will be replaced by the base64text. SprintF is used to format the string
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "data",
									MountPath: "/data",
								},
							},
						},
						{
							Name:    "convert",
							Image:   image,
							Command: []string{"sh", "-c"},
							Args:    []string{"pandoc", "/data/result.txt", "-o", "/data/text.md"}, // this cod ensures that the base64 text is converted to markdown
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "data",
									MountPath: "/data",
								},
							},
						},
					},
					Containers: []corev1.Container{
						{
							Name:    "main",
							Image:   "alpine",
							Command: []string{"sh", "-c", "sleep 3600"}, // we sleep 3600 seconds to keep the pod running
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "data",
									MountPath: "/data",
								},
							},
						},
					},
					Volumes: []corev1.Volume{ // volumes are used to share data between containers
						{
							Name: "data",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{},
							},
						},
					},
				},
			},
		},
	}

	return &j, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CalculateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&culatecomv1.Calculate{}).
		Named("calculate").
		Complete(r)
}
