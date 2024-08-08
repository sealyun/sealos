// Copyright © 2023 sealos.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resources

import (
	"strings"

	"github.com/labring/sealos/controllers/pkg/utils/label"

	corev1 "k8s.io/api/core/v1"

	sealos_networkmanager "github.com/dinoallo/sealos-networkmanager-protoapi"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
AppType (sort by label) :
	Database： app.kubernetes.io/instance=gitea  app.kubernetes.io/managed-by=kubeblocks apps.kubeblocks.io/component-name
	AppLaunchpad：app: xxx
	Terminal： TerminalID: xxx
	Cronjob：job-name: xxx
	Other: in addition to the above all labels
*/

const (
	Pod = "Pod"
	PVC = "PVC"
)

const (
	DBPodLabelInstanceKey       = "app.kubernetes.io/instance"
	DBPodLabelManagedByKey      = "app.kubernetes.io/managed-by"
	DBPodLabelManagedByValue    = "kubeblocks"
	DBPodLabelComponentNameKey  = "apps.kubeblocks.io/component-name"
	TerminalIDLabelKey          = "TerminalID"
	AppLabelKey                 = "app"
	AppDeployLabelKey           = "cloud.sealos.io/app-deploy-manager"
	AppStoreDeployLabelKey      = "cloud.sealos.io/deploy-on-sealos"
	JobNameLabelKey             = "job-name"
	ACMEChallengeKey            = "acme.cert-manager.io/http01-solver"
	KubeBlocksBackUpName        = "kubeblocks-backup-data"
	dataProtectionBackupRepoKey = "dataprotection.kubeblocks.io/backup-repo-name"
	InstanceLabelKey            = "app.kubernetes.io/instance"
)

type ResourceNamed struct {
	_name string
	// db or app or terminal or job or other
	_type      string
	parentType string
	parentName string
	labels     map[string]string
}

func NewResourceNamed(cr client.Object) *ResourceNamed {
	labels := cr.GetLabels()
	p := &ResourceNamed{labels: labels}
	switch {
	case labels[DBPodLabelComponentNameKey] != "":
		p._type = DB
		p._name = labels[DBPodLabelInstanceKey]
	case labels[TerminalIDLabelKey] != "" || (labels[label.AppManagedBy] == label.DefaultManagedBy && labels[label.AppPartOf] == "terminal"):
		p._type = TERMINAL
		p._name = ""
	case labels[AppLabelKey] != "":
		p._type = APP
		p._name = labels[AppLabelKey]
	case labels[AppDeployLabelKey] != "":
		p._type = APP
		p._name = labels[AppDeployLabelKey]
	case labels[JobNameLabelKey] != "":
		p._type = JOB
		p._name = strings.SplitN(labels[JobNameLabelKey], "-", 2)[0]
	case labels[ACMEChallengeKey] != "":
		p._type = APP
		p._name = getACMEResolverName(cr)
	case cr.GetName() == KubeBlocksBackUpName || labels[dataProtectionBackupRepoKey] != "":
		p._type = DBBackup
		p._name = KubeBlocksBackUpName
		if labels[InstanceLabelKey] != "" {
			p._name = labels[InstanceLabelKey]
		}
	default:
		p._type = OTHER
		p._name = ""
	}
	return p
}

func (r *ResourceNamed) SetInstanceParent(instances map[string]struct{}) {
	for ins := range instances {
		if strings.HasPrefix(r._name, ins) {
			r.parentType = AppStore
			r.parentName = ins
		}
	}
}

func NewObjStorageResourceNamed(bucket string) *ResourceNamed {
	return &ResourceNamed{
		_type: ObjectStorage,
		_name: bucket,
	}
}

const (
	acmesolver                          = "acmesolver"
	acmesolverContainerArgsDomainPrefix = "--domain="
)

func getACMEResolverName(obj client.Object) string {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		return ""
	}
	for _, container := range pod.Spec.Containers {
		if container.Name != acmesolver {
			continue
		}
		for _, arg := range container.Args {
			if strings.HasPrefix(arg, acmesolverContainerArgsDomainPrefix) {
				return acmesolver + "-" + strings.TrimPrefix(arg, acmesolverContainerArgsDomainPrefix)
			}
		}
	}
	return pod.Name
}

func (r *ResourceNamed) Type() uint8 {
	return AppType[r._type]
}

func (r *ResourceNamed) ParentType() uint8 {
	return AppType[r.parentType]
}

func (r *ResourceNamed) ParentName() string {
	return r.parentName
}

func (r *ResourceNamed) Labels() map[string]string {
	label := make(map[string]string)
	switch r.Type() {
	case db:
		label[DBPodLabelComponentNameKey] = r.labels[DBPodLabelComponentNameKey]
		label[DBPodLabelInstanceKey] = r.labels[DBPodLabelInstanceKey]
	case terminal:
		label[TerminalIDLabelKey] = r.labels[TerminalIDLabelKey]
	case app:
		label[AppLabelKey] = r.labels[AppLabelKey]
	case job:
		label[JobNameLabelKey] = r.labels[JobNameLabelKey]
		//case other:
		//default:
	}
	return label
}

var notExistLabels = func() map[uint8][]*sealos_networkmanager.Label {
	labels := make(map[uint8][]*sealos_networkmanager.Label)
	for k := range AppTypeReverse {
		labels[k] = getNotExistLabels(k)
	}
	return labels
}()

func (r *ResourceNamed) GetNotExistLabels() []*sealos_networkmanager.Label {
	return notExistLabels[r.Type()]
}

func getNotExistLabels(tp uint8) []*sealos_networkmanager.Label {
	var labels []*sealos_networkmanager.Label
	for appType := range AppTypeReverse {
		if tp == appType {
			continue
		}
		switch appType {
		case db:
			labels = append(labels, &sealos_networkmanager.Label{
				Key: DBPodLabelComponentNameKey,
			}, &sealos_networkmanager.Label{
				Key: DBPodLabelManagedByKey,
			})
		case app:
			labels = append(labels, &sealos_networkmanager.Label{
				Key: AppLabelKey,
			})
		case terminal:
			labels = append(labels, &sealos_networkmanager.Label{
				Key: TerminalIDLabelKey,
			})
		case job:
			labels = append(labels, &sealos_networkmanager.Label{
				Key: JobNameLabelKey,
			})
		}
	}
	return labels
}

func (r *ResourceNamed) GetInLabels() []*sealos_networkmanager.Label {
	var labelsEqual []*sealos_networkmanager.Label
	switch r.Type() {
	case db:
		labelsEqual = append(labelsEqual, &sealos_networkmanager.Label{
			Key:   DBPodLabelComponentNameKey,
			Value: []string{r.labels[DBPodLabelComponentNameKey]},
		})
	case terminal:
		labelsEqual = append(labelsEqual, &sealos_networkmanager.Label{
			Key:   TerminalIDLabelKey,
			Value: []string{r.labels[TerminalIDLabelKey]},
		})
	case app:
		labelsEqual = append(labelsEqual, &sealos_networkmanager.Label{
			Key:   AppLabelKey,
			Value: []string{r.labels[AppLabelKey]},
		})
	case job:
		labelsEqual = append(labelsEqual, &sealos_networkmanager.Label{
			Key:   JobNameLabelKey,
			Value: []string{r.labels[JobNameLabelKey]},
		})
	}
	return labelsEqual
}

func (r *ResourceNamed) TypeString() string {
	return r._type
}

func (r *ResourceNamed) Name() string {
	return r._name
}

func (r *ResourceNamed) String() string {
	return r._type + "/" + r._name
}
