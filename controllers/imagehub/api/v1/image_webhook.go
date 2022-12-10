/*
Copyright 2022.

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

package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var imagelog = logf.Log.WithName("image-resource")

//+kubebuilder:webhook:path=/mutate-imagehub-sealos-io-v1-image,mutating=true,failurePolicy=fail,sideEffects=None,groups=imagehub.sealos.io,resources=images,verbs=create;update,versions=v1,name=mimage.kb.io,admissionReviewVersions=v1

// ImageMutater add lables to Images
type ImageMutater struct {
	Client  client.Client
	decoder *admission.Decoder
}

func (m *ImageMutater) Handle(ctx context.Context, req admission.Request) admission.Response {
	i := &Image{}
	err := m.decoder.Decode(req, i)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	imagelog.Info("default", "name", i.Name)
	i.ObjectMeta = initAnnotationAndLabels(i.ObjectMeta)
	i.ObjectMeta.Labels[SealosOrgLable] = i.Spec.Name.GetOrg()
	i.ObjectMeta.Labels[SealosRepoLabel] = i.Spec.Name.GetRepo()
	i.ObjectMeta.Labels[SealosTagLabel] = i.Spec.Name.GetTag()

	marshaledPod, err := json.Marshal(i)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}
	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-imagehub-sealos-io-v1-image,mutating=false,failurePolicy=fail,sideEffects=None,groups=imagehub.sealos.io,resources=images,verbs=create;update,versions=v1,name=vimage.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Image{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (i *Image) ValidateCreate() error {
	imagelog.Info("validate create", "name", i.Name)
	if !i.checkSpecName() {
		return fmt.Errorf("image name illegal")
	}
	if !i.checkLables() {
		return fmt.Errorf("image lables illegal")
	}
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (i *Image) ValidateUpdate(old runtime.Object) error {
	imagelog.Info("validate update", "name", i.Name)
	if !i.checkSpecName() {
		return fmt.Errorf("image name illegal")
	}
	if !i.checkLables() {
		return fmt.Errorf("image lables illegal")
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (i *Image) ValidateDelete() error {
	imagelog.Info("validate delete", "name", i.Name)
	return nil
}
