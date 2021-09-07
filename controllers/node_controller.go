/*
Copyright 2021.

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

package controllers

import (
	"context"
	"regexp"

	configv1alpha1 "github.com/snapp-incubator/node-config-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// NodeReconciler reconciles a Node object
type NodeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=nodes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=core,resources=nodes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Node object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.9.2/pkg/reconcile
func (r *NodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Lookup the route instance for this reconcile request
	node := &corev1.Node{}
	err := r.Get(ctx, req.NamespacedName, node)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			logger.Info("Resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		logger.Error(err, "Failed to get Objet")
		return ctrl.Result{}, err
	}

	// get the list of NodeConfigs
	ncList := &configv1alpha1.NodeConfigList{}
	err = r.List(ctx, ncList, &client.ListOptions{})
	if err != nil {
		logger.Error(err, "Failed to list NodeConfigs")
		return ctrl.Result{}, err
	}
	// for NodeConfig in list, check if node matches any NodeConfig
	for _, nc := range ncList.Items {
		if nodeMatchNodeConfig(*node, nc.Spec.Match) {
			// Update the Node with the NodeConfig
			updateNode, match := nodeMergeNodeConfig(*node, nc.Spec.Merge)
			// if they are already the same, continue
			if match {
				logger.Info("Node has desired NodeConfig", "nodeconfig.Name", nc.Name)
				continue
			}
			logger.Info("Updating Node with NodeConfig", "nodeconfig.Name", nc.Name)
			err = r.Update(ctx, &updateNode)
			if err != nil {
				logger.Error(err, "Failed to update Node", "nodeconfig.Name", nc.Name)
				return ctrl.Result{}, err
			}

		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Node{}).
		Complete(r)
}

func nodeMatchNodeConfig(node corev1.Node, m configv1alpha1.Match) bool {
	for _, nodeNamePattern := range m.NodeNamePatterns {
		pattern := "^" + nodeNamePattern + "$"
		match, _ := regexp.MatchString(pattern, node.Name)
		// Invalid regular expression, moving on to next rule
		if match {
			return true
		}
	}
	return false
}

func nodeMergeNodeConfig(node corev1.Node, m configv1alpha1.Merge) (updatedNode corev1.Node, match bool) {
	match = true
	if m.Labels != nil {
		for k, v := range m.Labels {
			if rv, ok := node.Labels[k]; !ok || rv != v {
				match = false
				node.Labels[k] = v
			}
		}
	}
	if m.Annotations != nil {
		for k, v := range m.Annotations {
			if rv, ok := node.Annotations[k]; !ok || rv != v {
				match = false
				node.Annotations[k] = v
			}
		}
	}

	if m.Taints != nil {
		newTaints := node.Spec.Taints
		for _, taint := range m.Taints {
			found := false
			for i, nodeTaint := range node.Spec.Taints {
				if nodeTaint.Key == taint.Key {
					found = true
					if nodeTaint.Value != taint.Value || nodeTaint.Effect != taint.Effect {
						match = false
						newTaints[i] = taint
					}
				}
			}
			if !found {
				match = false
				newTaints = append(newTaints, taint)
			}
		}
		node.Spec.Taints = newTaints
	}

	return node, match
}
