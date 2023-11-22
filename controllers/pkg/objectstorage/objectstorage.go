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

package objectstorage

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func GetUserObjectStorageSize(client *minio.Client, username string) (int64, int64, error) {
	buckets, err := client.ListBuckets(context.Background())
	if err != nil {
		return 0, 0, fmt.Errorf("failed to list object storage buckets")
	}

	var expectBuckets []string
	for _, bucket := range buckets {
		if strings.HasPrefix(bucket.Name, username) {
			expectBuckets = append(expectBuckets, bucket.Name)
		}
	}

	var totalSize int64
	var objectsCount int64
	for _, bucketName := range expectBuckets {
		objects := client.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
			Recursive: true,
		})
		for object := range objects {
			totalSize += object.Size
			objectsCount++
		}
	}

	return totalSize, objectsCount, nil
}

func GetUserObjectStorageFlow(client *minio.Client, host, username string) (int64, error) {
	buckets, err := client.ListBuckets(context.Background())
	if err != nil {
		return 0, fmt.Errorf("failed to list object storage buckets")
	}

	var expectBuckets []string
	for _, bucket := range buckets {
		if strings.HasPrefix(bucket.Name, username) {
			expectBuckets = append(expectBuckets, bucket.Name)
		}
	}

	var totalFlow int64
	for _, bucketName := range expectBuckets {
		flow, err := QueryPrometheus(host, bucketName)
		if err != nil {
			return 0, fmt.Errorf("failed to query prometheus, bucket: %v, err: %v", bucketName, err)
		}
		totalFlow += flow
	}

	return totalFlow, nil
}

func QueryPrometheus(host, bucketName string) (int64, error) {
	client, err := api.NewClient(api.Config{
		Address: host,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to new prometheus client, host: %v, err: %v", host, err)
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rcvdQuery := "sum(minio_bucket_traffic_received_bytes{bucket=\"" + bucketName + "\"})"
	rcvdResult, rcvdWarnings, err := v1api.Query(ctx, rcvdQuery, time.Now(), v1.WithTimeout(5*time.Second))
	if err != nil {
		return 0, fmt.Errorf("failed to query prometheus, query: %v, err: %v", rcvdQuery, err)
	}

	if len(rcvdWarnings) > 0 {
		return 0, fmt.Errorf("there are warnings: %v", rcvdWarnings)
	}

	sentQuery := "sum(minio_bucket_traffic_sent_bytes{bucket=\"" + bucketName + "\"})"
	sentResult, sentWarnings, err := v1api.Query(ctx, sentQuery, time.Now(), v1.WithTimeout(5*time.Second))
	if err != nil {
		return 0, fmt.Errorf("failed to query prometheus, query: %v, err: %v", sentQuery, err)
	}

	if len(sentWarnings) > 0 {
		return 0, fmt.Errorf("there are warnings: %v", sentWarnings)
	}

	re := regexp.MustCompile(`\d+`)
	rcvdStr := re.FindString(rcvdResult.String())
	sentStr := re.FindString(sentResult.String())

	rcvdBytes, err := strconv.ParseInt(rcvdStr, 10, 64)
	sentBytes, err := strconv.ParseInt(sentStr, 10, 64)

	fmt.Printf("received bytes: %d, send bytes: %d\n", rcvdBytes, sentBytes)

	return rcvdBytes + sentBytes, nil
}
