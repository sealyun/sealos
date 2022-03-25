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

package image

import (
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// defaultImageService is the default service, which is used for image pull/push
type defaultImageService struct {
}

func (d *defaultImageService) Rename(src, dst string) error {
	panic("implement me")
}

func (d *defaultImageService) Remove(images ...string) error {
	panic("implement me")
}

func (d *defaultImageService) Inspect(image string) (*v1.Image, error) {
	panic("implement me")
}

func (d *defaultImageService) Build(options BuildOptions, contextDir, imageName string) error {
	panic("implement me")
}

func (d *defaultImageService) Prune() error {
	panic("implement me")
}

func (d *defaultImageService) ListImages(opt ListOptions) ([]v1.Image, error) {
	panic("implement me")
}

func NewImageService() (Service, error) {
	return &defaultImageService{}, nil
}
