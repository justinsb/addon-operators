package controllers

import (
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/status"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	api "sigs.k8s.io/addon-operators/api/v1alpha1"
)

var _ reconcile.Reconciler = &NodeLocalDNSReconciler{}

// NodeLocalDNSReconciler reconciles a NodeLocalDNS object
type NodeLocalDNSReconciler struct {
	client.Client
	Log logr.Logger

	declarative.Reconciler
}

// +kubebuilder:rbac:groups=addons.k8s.io,resources=nodelocaldns,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=addons.k8s.io,resources=nodelocaldns/status,verbs=get;update;patch

func (r *NodeLocalDNSReconciler) SetupWithManager(mgr ctrl.Manager) error {
	addon.Init()

	labels := map[string]string{
		"k8s-app": "nodelocaldns",
	}

	watchLabels := declarative.SourceLabel(mgr.GetScheme())

	if err := r.Reconciler.Init(mgr, &api.NodeLocalDNS{},
		declarative.WithObjectTransform(declarative.AddLabels(labels)),
		declarative.WithOwner(declarative.SourceAsOwner),
		declarative.WithLabels(watchLabels),
		declarative.WithStatus(status.NewBasic(mgr.GetClient())),
		// TODO: add an application to your manifest:  declarative.WithObjectTransform(addon.TransformApplicationFromStatus),
		// TODO: add an application to your manifest:  declarative.WithManagedApplication(watchLabels),
		declarative.WithObjectTransform(addon.ApplyPatches),
	); err != nil {
		return err
	}

	c, err := controller.New("nodelocaldns-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to NodeLocalDNS
	err = c.Watch(&source.Kind{Type: &api.NodeLocalDNS{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to deployed objects
	_, err = declarative.WatchAll(mgr.GetConfig(), c, r, watchLabels)
	if err != nil {
		return err
	}

	return nil
}
