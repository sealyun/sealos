/*
Copyright 2023 yxxchange@163.com.

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
package cloudtool

import (
	"context"
	"net/http"
	"strings"
	"time"

	ntf "github.com/labring/sealos/controllers/common/notification/api/v1"
	"github.com/labring/sealos/pkg/utils/logger"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	cl "sigs.k8s.io/controller-runtime/pkg/client"
)

type Cloud interface {
	setCloudArgs(method string, url string) interface{}
	createRequest(interface{}) (*http.Request, error)
	getResponse(*http.Request) (*http.Response, error)
	readResponse(*http.Response) ([]byte, error)
}

type HandlerCR interface {
	produceCR(map[string][]string, []byte) []ntf.Notification
}

func CloudPull(cloud Cloud, method string, url string) ([]byte, error) {
	content := cloud.setCloudArgs(method, url)
	var req *http.Request
	var resp *http.Response
	var err error
	if req, err = cloud.createRequest(content); err != nil {
		return nil, err
	}
	if resp, err = cloud.getResponse(req); err != nil {
		return nil, err
	}
	return cloud.readResponse(resp)
}

func CloudCreateCR(h HandlerCR, client cl.Client, resp []byte) error {
	start := time.Now()
	namespaceGroup, err := getNamespaceGroup(client)
	duration := time.Since(start)
	logger.Info("duration of getnamespaceGroup:", duration)
	if err != nil {
		return err
	}
	CRs := h.produceCR(namespaceGroup, resp)
	for _, res := range CRs {
		var tmp ntf.Notification
		key := types.NamespacedName{Namespace: res.Namespace, Name: res.Name}
		if err = client.Get(context.Background(), key, &tmp); err == nil {
			res.ResourceVersion = tmp.ResourceVersion
			if err = client.Update(context.Background(), &res); err != nil {
				logger.Info("failed to create ", "user_id: ", res.Namespace, "notification id: ", res.Name, "Error: ", err)
			}
		} else {
			if cl.IgnoreNotFound(err) == nil {
				if err = client.Create(context.Background(), &res); err != nil {
					logger.Info("failed to create ", "user_id: ", res.Namespace, "notification id: ", res.Name, "Error: ", err)
				}
			} else {
				logger.Error("failed to create ", "user_id: ", res.Namespace, "notification id: ", res.Name, "Error: ", err)
			}
		}
	}
	return nil
}

func getNamespaceGroup(client cl.Client) (map[string][]string, error) {
	namespaceList := &corev1.NamespaceList{}
	if err := client.List(context.Background(), namespaceList); err != nil {
		logger.Error("failed to get namespace resource ", err)
		return map[string][]string{}, err
	}
	//divide namespace to diff groups
	var namespaceGroup = map[string][]string{
		"ns-":   {},
		"adm-":  {},
		"root-": {},
	}
	for _, namespace := range namespaceList.Items {
		for prefix := range namespaceGroup {
			if strings.HasPrefix(namespace.Name, prefix) {
				namespaceGroup[prefix] = append(namespaceGroup[prefix], namespace.Name)
				break
			}
		}
	}
	return namespaceGroup, nil
}
