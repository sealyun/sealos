module github.com/fanux/sealos

go 1.15

require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.985
	github.com/containers/image/v5 v5.16.0
	github.com/containers/storage v1.36.0 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/distribution/distribution/v3 v3.0.0-20211125133600-cc4627fc6e5f
	github.com/docker/cli v20.10.12+incompatible
	github.com/docker/docker v20.10.8+incompatible
	github.com/emirpasic/gods v1.12.0
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/huaweicloud/huaweicloud-sdk-go-v3 v0.0.72
	github.com/imdario/mergo v0.3.12
	github.com/mitchellh/go-homedir v1.1.0
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.16.0
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.0.3-0.20211202193544-a5463b7f9c84
	github.com/pelletier/go-toml v1.9.3
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.13.0
	github.com/schollz/progressbar/v3 v3.8.5
	github.com/sealyun/lvscare v1.1.3-beta.2
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/vishvananda/netlink v1.1.1-0.20201029203352-d40f9887b852
	go.etcd.io/etcd/client/v3 v3.5.1
	go.etcd.io/etcd/etcdutl/v3 v3.5.1
	go.uber.org/zap v1.19.0
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	k8s.io/api v0.21.0
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
	k8s.io/cluster-bootstrap v0.21.0
	k8s.io/kube-proxy v0.21.0
	k8s.io/kubelet v0.21.0
	k8s.io/utils v0.0.0-20211116205334-6203023598ed
	sigs.k8s.io/yaml v1.2.0
)

replace github.com/vishvananda/netlink => github.com/vishvananda/netlink v1.1.0
