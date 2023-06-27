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

package manager

import (
	"encoding/json"
	"reflect"

	corev1 "k8s.io/api/core/v1"
)

type Config struct {
	CollectorURL      string `json:"CollectorURL"`
	NotificationURL   string `json:"NotificationURL"`
	RegisterURL       string `json:"RegisterURL"`
	CloudSyncURL      string `json:"CloudSyncURL"`
	LicenseMonitorURL string `json:"LicenseMonitorURL"`
	// Add other fields here to support future expansion needs.
}

type SyncResponse struct {
	Config `json:",inline"`
}

type SyncRequest struct {
	UID string `json:"uid"`
}

func IsConfigMapChanged(resp SyncResponse, cm *corev1.ConfigMap) bool {
	var changed bool = false
	var configMapJson map[string]string
	if err := json.Unmarshal([]byte(cm.Data["config.json"]), &configMapJson); err != nil {
		panic(err)
	}
	newConfigMapValue := reflect.ValueOf(resp)
	newConfigMapType := newConfigMapValue.Type()

	for i := 0; i < newConfigMapValue.NumField(); i++ {
		fieldName := newConfigMapType.Field(i).Name
		fieldValue := newConfigMapValue.Field(i).String()

		cmKey := fieldName

		if cmValue, ok := cm.Data[cmKey]; !ok || cmValue != fieldValue {
			configMapJson[fieldName] = fieldValue
			changed = true
		}
	}
	if changed {
		updatedJson, err := json.Marshal(configMapJson)
		if err != nil {
			panic(err)
		}
		cm.Data["config.json"] = string(updatedJson)
	}

	return changed
}
