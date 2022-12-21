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
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RepositorySpec defines the desired state of Repository, Org automatically generated by Name.
type RepositorySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	//+kubebuilder:validation:Required
	Name RepoName `json:"name"` // e.g: "libring/mysql"
}

type RepoName string

// IsLegal check name is legal
// name.eg: labring/mysql:v8.0.31
func (n RepoName) IsLegal() bool {
	return len(strings.Split(string(n), "/")) == 2
}

func (n RepoName) GetOrg() string {
	str := strings.FieldsFunc(string(n), func(r rune) bool {
		return r == '/' || r == ':'
	})
	return str[0]
}

func (n RepoName) GetRepo() string {
	str := strings.FieldsFunc(string(n), func(r rune) bool {
		return r == '/' || r == ':'
	})
	return str[1]
}

func (n RepoName) ToMetaName() string {
	return n.GetOrg() + "." + n.GetRepo()
}

type RepoInfo RepositoryStatus

type TagList []TagData

// RepositoryStatus defines the observed state of Repository
type RepositoryStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Name      RepoName `json:"name,omitempty"` // e.g: "libring/mysql"
	Tags      TagList  `json:"tags,omitempty"`
	LatestTag *TagData `json:"latestTag,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,shortName=repo

// Repository is the Schema for the repositories API
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RepositorySpec   `json:"spec,omitempty"`
	Status RepositoryStatus `json:"status,omitempty"`
}

func (r *Repository) checkSpecName() bool {
	return r.Spec.Name.IsLegal()
}
func (r *Repository) checkLabels() bool {
	return r.Labels[SealosOrgLable] == r.Spec.Name.GetOrg() &&
		r.Labels[SealosRepoLabel] == r.Spec.Name.GetRepo()
}
func (r *Repository) getSpecName() string {
	return string(r.Spec.Name)
}
func (r *Repository) getOrgName() string {
	return r.Spec.Name.GetOrg()
}
func (r *Repository) getName() string {
	return r.Name
}

func (r *Repository) genKeywordsLabels(img *Image) {
	mp := make(map[string]string)
	for _, keyword := range img.Spec.DetailInfo.Keywords {
		label := fmt.Sprintf("%s%s", KeywordsLabelPrefix, keyword)
		mp[label] = ""
	}
}

//+kubebuilder:object:root=true

// RepositoryList contains a list of Repository
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}
