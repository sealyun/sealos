/*
Copyright 2023.

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
	"fmt"
	meteringcommonv1 "github.com/labring/sealos/controllers/common/metering/api/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"os"
	"strconv"
	"time"

	"github.com/go-logr/logr"
	accountv1 "github.com/labring/sealos/controllers/account/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DebtReconciler reconciles a Debt object
type DebtReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	logr.Logger
	accountSystemNamespace string
}

var DebtConfig = accountv1.DefaultDebtConfig

//+kubebuilder:rbac:groups=account.sealos.io,resources=debts,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=account.sealos.io,resources=debts/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=account.sealos.io,resources=debts/finalizers,verbs=update
//+kubebuilder:rbac:groups=account.sealos.io,resources=accounts,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=metering.common.sealos.io,resources=extensionresourceprices,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app,resources=daemonSets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app,resources=deployments,verbs=get;list;watch;create;update;patch;delete

func (r *DebtReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	account := &accountv1.Account{}
	if err := r.Get(ctx, req.NamespacedName, account); err != nil {
		r.Logger.Error(err, err.Error())
		return ctrl.Result{}, err
	}

	debt := &accountv1.Debt{}
	if err := r.Get(ctx, client.ObjectKey{Name: GetDebtName(account.Name), Namespace: r.accountSystemNamespace}, debt); client.IgnoreNotFound(err) != nil {
		return ctrl.Result{}, err
	} else if err != nil && client.IgnoreNotFound(err) == nil {
		if err := r.syncDebt(ctx, account, debt); err != nil {
			return ctrl.Result{}, err
		} else {
			r.Logger.Info("create or update debt success", "debt", debt)
			return ctrl.Result{Requeue: true, RequeueAfter: time.Second}, nil
		}
	}

	r.Logger.Info("debt info", "debt", debt)
	if debt.Status.AccountDebtStatus == "" {
		debt.Status.AccountDebtStatus = accountv1.DebtStatusNormal
	}

	if err := r.reconcileDebtStatus(ctx, debt, account); err != nil {
		r.Logger.Error(err, "reconcile debt status error")
		return ctrl.Result{}, err
	}
	return ctrl.Result{Requeue: true, RequeueAfter: time.Minute}, nil
}

func (r *DebtReconciler) reconcileDebtStatus(ctx context.Context, debt *accountv1.Debt, account *accountv1.Account) error {
	oweamount := account.Status.Balance - account.Status.DeductionBalance
	Updataflag := false
	if oweamount > 0 && debt.Status.AccountDebtStatus == accountv1.DebtStatusSmall || debt.Status.AccountDebtStatus == accountv1.DebtStatusMedium || debt.Status.AccountDebtStatus == accountv1.DebtStatusLarge {
		debt.Status.AccountDebtStatus = accountv1.DebtStatusNormal
		debt.Status.LastUpdateTimestamp = time.Now().Unix()
		Updataflag = true
		if err := r.change2Normal(ctx, *account); err != nil {
			return err
		}
	}

	normalPrice, ok := DebtConfig[accountv1.DebtStatusNormal]
	if !ok {
		r.Error(fmt.Errorf("get normal price error"), "")
	}
	if debt.Status.AccountDebtStatus == accountv1.DebtStatusNormal && oweamount < normalPrice {
		debt.Status.AccountDebtStatus = accountv1.DebtStatusSmall
		debt.Status.LastUpdateTimestamp = time.Now().Unix()
		Updataflag = true
		if err := r.sendSmallNotice(ctx, account.Name); err != nil {
			return err
		}
		if err := r.change2Small(ctx, *account); err != nil {
			return err
		}
	}

	smallBlockTimeSecond, ok := DebtConfig[accountv1.DebtStatusSmall]
	if !ok {
		return fmt.Errorf("get smallBlockTimeSecond, error")
	}
	if debt.Status.AccountDebtStatus == accountv1.DebtStatusSmall && (time.Now().Unix()-debt.Status.LastUpdateTimestamp) > smallBlockTimeSecond {
		debt.Status.AccountDebtStatus = accountv1.DebtStatusMedium
		debt.Status.LastUpdateTimestamp = time.Now().Unix()
		Updataflag = true
		if err := r.sendMediumNotice(ctx, account.Name); err != nil {
			return err
		}
		if err := r.change2Medium(ctx, *account); err != nil {
			return err
		}
	}

	mediumBlockTimeSecond, ok := DebtConfig[accountv1.DebtStatusMedium]
	if !ok {
		return fmt.Errorf("get mediumBlockTimeSecond, error")
	}
	if debt.Status.AccountDebtStatus == accountv1.DebtStatusMedium && (time.Now().Unix()-debt.Status.LastUpdateTimestamp) > mediumBlockTimeSecond {
		debt.Status.AccountDebtStatus = accountv1.DebtStatusLarge
		debt.Status.LastUpdateTimestamp = time.Now().Unix()
		Updataflag = true
		if err := r.sendLargeNotice(ctx, account.Name); err != nil {
			return err
		}
		if err := r.change2Large(ctx, *account); err != nil {
			return err
		}
	}

	if Updataflag {
		if err := r.Status().Update(ctx, debt); err != nil {
			return err
		}
	}
	return nil
}

func (r *DebtReconciler) change2Normal(_ context.Context, account accountv1.Account) error {
	r.Logger.Info("change status to normal", "account", account.Name)
	return nil
}

func (r *DebtReconciler) change2Small(_ context.Context, account accountv1.Account) error {
	r.Logger.Info("change status to Small", "account", account.Name)
	return nil
}

func (r *DebtReconciler) change2Medium(_ context.Context, account accountv1.Account) error {
	r.Logger.Info("change status to medium", "account", account.Name)
	return nil
}

func (r *DebtReconciler) change2Large(ctx context.Context, account accountv1.Account) error {
	r.Logger.Info("change status to large", "account", account.Name)
	return r.deleteUserResource(ctx, GetUserNamespace(account.Name))
}

func (r *DebtReconciler) syncDebt(ctx context.Context, account *accountv1.Account, debt *accountv1.Debt) error {
	debt.Name = GetDebtName(account.Name)
	debt.Namespace = r.accountSystemNamespace
	if _, err := controllerutil.CreateOrUpdate(ctx, r.Client, debt, func() error {
		debt.Spec.UserName = account.Name
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func GetDebtName(AccountName string) string {
	return fmt.Sprintf("%s%s", accountv1.DebtPrefix, AccountName)
}

func GetUserNamespace(AccountName string) string {
	return "ns-" + AccountName
}

func (r *DebtReconciler) sendNotice(_ context.Context, _ string, _ any) error {
	return nil
}

func (r *DebtReconciler) sendSmallNotice(ctx context.Context, accountName string) error {
	return r.sendNotice(ctx, accountName, nil)
}

func (r *DebtReconciler) sendMediumNotice(ctx context.Context, accountName string) error {
	return r.sendNotice(ctx, accountName, nil)
}
func (r *DebtReconciler) sendLargeNotice(ctx context.Context, accountName string) error {
	return r.sendNotice(ctx, accountName, nil)
}

func (r *DebtReconciler) deleteUserResource(ctx context.Context, namespace string) error {
	r.Logger.Info("enter delete user resource", "namespace", namespace)
	// delete register metering resource
	extensonResources := meteringcommonv1.ExtensionResourcePriceList{}
	if err := r.List(ctx, &extensonResources); client.IgnoreNotFound(err) != nil {
		return err
	}

	for _, extensonResource := range extensonResources.Items {
		for _, groupVersionKind := range extensonResource.Spec.GroupVersionKinds {
			u := unstructured.UnstructuredList{}
			u.SetGroupVersionKind(schema.GroupVersionKind{
				Group:   groupVersionKind.Group,
				Version: groupVersionKind.Version,
				Kind:    groupVersionKind.Kind,
			})
			if err := r.List(ctx, &u, client.InNamespace(namespace)); client.IgnoreNotFound(err) != nil {
				return err
			}
			for _, item := range u.Items {
				r.Logger.Info("delete resource", "resource name:", item.GetName(), "get GVK", item.GroupVersionKind())
				if err := r.Delete(ctx, &item); client.IgnoreNotFound(err) != nil {
					return err
				}
			}
		}
	}

	// delete all pod
	podlist := v1.PodList{}
	if err := r.List(ctx, &podlist, client.InNamespace(namespace)); client.IgnoreNotFound(err) != nil {
		return err
	}
	for _, pod := range podlist.Items {
		if err := r.Delete(ctx, &pod); err != nil {
			r.Logger.Info("delete resource", "resource name:", pod.GetName(), "get GVK", pod.GroupVersionKind())
			return err
		}
	}

	// delete all deployment
	deploylist := appsv1.DeploymentList{}
	if err := r.List(ctx, &deploylist, client.InNamespace(namespace)); client.IgnoreNotFound(err) != nil {
		return err
	}
	for _, deploy := range deploylist.Items {
		r.Logger.Info("delete resource", "resource name:", deploy.GetName(), "get GVK", deploy.GroupVersionKind())
		if err := r.Delete(ctx, &deploy); err != nil {
			return err
		}
	}

	// delete all daemonset
	daemonsetlist := appsv1.DaemonSetList{}
	if err := r.List(ctx, &daemonsetlist, client.InNamespace(namespace)); client.IgnoreNotFound(err) != nil {
		return err
	}
	for _, daemonset := range daemonsetlist.Items {
		r.Logger.Info("delete resource", "resource name:", daemonset.GetName(), "get GVK", daemonset.GroupVersionKind())
		if err := r.Delete(ctx, &daemonset); err != nil {
			return err
		}
	}

	// delete all replicaset
	replicalist := appsv1.ReplicaSetList{}
	if err := r.List(ctx, &replicalist, client.InNamespace(namespace)); client.IgnoreNotFound(err) != nil {
		return err
	}
	for _, replica := range replicalist.Items {
		r.Logger.Info("delete resource", "resource name:", replica.GetName(), "get GVK", replica.GroupVersionKind())
		if err := r.Delete(ctx, &replica); err != nil {
			return err
		}
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DebtReconciler) SetupWithManager(mgr ctrl.Manager) error {
	const controllerName = "DebtController"
	r.Logger = ctrl.Log.WithName(controllerName)
	var smallBlockWaitSecondString, mediumBlockWaitSecondString string
	if smallBlockWaitSecondString = os.Getenv("SMALL_BLOCK_WAIT_SECOND"); smallBlockWaitSecondString != "" {
		smallBlockWaitSecond, err := strconv.Atoi(smallBlockWaitSecondString)
		if err != nil {
			r.Logger.Error(err, "SMALL_BLOCK_WAIT_SECOND is not a number")
		}
		DebtConfig["Small"] = int64(smallBlockWaitSecond)
	}

	if mediumBlockWaitSecondString = os.Getenv("MEDIUM_BLOCK_WAIT_SECOND"); mediumBlockWaitSecondString != "" {
		if mediumBlockWaitSecond, err := strconv.Atoi(mediumBlockWaitSecondString); err != nil {
			r.Logger.Error(err, "MEDIUM_BLOCK_WAIT_SECOND is not a number")
		} else {
			DebtConfig["Medium"] = int64(mediumBlockWaitSecond)
		}

	}
	r.Logger.Info("DebtConfig", "DebtConfig", DebtConfig)
	r.accountSystemNamespace = os.Getenv("ACCOUNT_SYSTEM_NAMESPACE")
	if r.accountSystemNamespace == "" {
		r.accountSystemNamespace = "account-system"
	}
	return ctrl.NewControllerManagedBy(mgr).
		// update status should not enter reconcile
		For(&accountv1.Debt{}, builder.WithPredicates(predicate.Or(predicate.GenerationChangedPredicate{}))).
		// Uncomment the following line adding a pointer to an instance of the controlled resource as an argument
		Watches(&source.Kind{Type: &accountv1.Account{}}, &handler.EnqueueRequestForObject{}).
		Complete(r)
}
