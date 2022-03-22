/*
Copyright 2022 cuisongliu@qq.com.

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

package store

import (
	"errors"
	"github.com/fanux/sealos/pkg/types/v1beta1"
	"github.com/fanux/sealos/pkg/utils/hash"
)

type Store interface {
	Save(p *v1beta1.Resource) error
}

type store struct {
}

func (s *store) Save(p *v1beta1.Resource) error {
	if p.Spec.Path == "" {
		return errors.New("package path not allow empty")
	}
	p.Name = hash.ToString(p.Spec.Path)
	if p.Spec.Type.IsOCI() {
		return s.oci(p)
	}
	if p.Spec.Type.IsTarGz() {
		return s.tarGz(p)
	}
	return s.dir(p,p.Spec.Path,false)
}

func NewStore() Store {
	return &store{}
}
