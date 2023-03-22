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

package server

import (
	"context"

	"github.com/labring/image-cri-shim/pkg/types"

	"github.com/labring/sealos/pkg/registry/name"
	"github.com/labring/sealos/pkg/utils/registry"

	types2 "github.com/docker/docker/api/types"

	api "k8s.io/cri-api/pkg/apis/runtime/v1"

	"github.com/labring/sealos/pkg/utils/logger"
)

type v1ImageService struct {
	imageClient       api.ImageServiceClient
	CRIConfigs        map[string]types2.AuthConfig
	OfflineCRIConfigs map[string]types2.AuthConfig
}

func (s *v1ImageService) ListImages(ctx context.Context,
	req *api.ListImagesRequest) (*api.ListImagesResponse, error) {
	logger.Debug("ListImages: %+v", req)
	rsp, err := s.imageClient.ListImages(ctx, req)

	if err != nil {
		return nil, err
	}

	return rsp, err
}

func (s *v1ImageService) ImageStatus(ctx context.Context,
	req *api.ImageStatusRequest) (*api.ImageStatusResponse, error) {
	logger.Debug("ImageStatus: %+v", req)
	if req.Image != nil {
		req.Image.Image, _, _ = replaceImage(req.Image.Image, "ImageStatus", s.OfflineCRIConfigs)
	}
	rsp, err := s.imageClient.ImageStatus(ctx, req)

	if err != nil {
		return nil, err
	}

	return rsp, err
}

func (s *v1ImageService) PullImage(ctx context.Context,
	req *api.PullImageRequest) (*api.PullImageResponse, error) {
	logger.Debug("PullImage begin: %+v", req)
	if req.Image != nil {
		imageName, ok, auth := replaceImage(req.Image.Image, "PullImage", s.OfflineCRIConfigs)
		if req.Auth == nil {
			if ok {
				req.Auth = types.ToV1AuthConfig(auth)
			} else {
				ref, _ := name.ParseReference(imageName)
				for domain, v := range s.CRIConfigs {
					if registry.NormalizeRegistry(domain) == ref.Context().RegistryStr() {
						req.Auth = types.ToV1AuthConfig(&v)
						break
					}
				}
			}
		}
		req.Image.Image = imageName
	}
	logger.Debug("PullImage after: %+v", req)
	rsp, err := s.imageClient.PullImage(ctx, req)
	if err != nil {
		return nil, err
	}

	return rsp, err
}

func (s *v1ImageService) RemoveImage(ctx context.Context,
	req *api.RemoveImageRequest) (*api.RemoveImageResponse, error) {
	logger.Debug("RemoveImage: %+v", req)
	if req.Image != nil {
		req.Image.Image, _, _ = replaceImage(req.Image.Image, "RemoveImage", s.OfflineCRIConfigs)
	}
	rsp, err := s.imageClient.RemoveImage(ctx, req)

	if err != nil {
		return nil, err
	}

	return rsp, err
}

func (s *v1ImageService) ImageFsInfo(ctx context.Context,
	req *api.ImageFsInfoRequest) (*api.ImageFsInfoResponse, error) {
	logger.Debug("ImageFsInfo: %+v", req)
	rsp, err := s.imageClient.ImageFsInfo(ctx, req)

	if err != nil {
		return nil, err
	}

	return rsp, err
}
