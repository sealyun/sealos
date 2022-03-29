package processor

import (
	"fmt"

	"github.com/fanux/sealos/pkg/clusterfile"
	"github.com/fanux/sealos/pkg/config"
	"github.com/fanux/sealos/pkg/filesystem"
	"github.com/fanux/sealos/pkg/guest"
	"github.com/fanux/sealos/pkg/image"
	"github.com/fanux/sealos/pkg/runtime"
	v2 "github.com/fanux/sealos/pkg/types/v1beta1"
	"github.com/fanux/sealos/pkg/utils/contants"
	"github.com/fanux/sealos/pkg/utils/yaml"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type CreateProcessor struct {
	ClusterFile     clusterfile.Interface
	ImageManager    image.Service
	ClusterManager  image.ClusterService
	RegistryManager image.RegistryService
	Runtime         runtime.Interface
	Guest           guest.Interface
	Config          config.Interface
	img             *v1.Image
	cManifest       *image.ClusterManifest
}

func (c *CreateProcessor) Execute(cluster *v2.Cluster) error {
	err := yaml.MarshalYamlToFile(contants.Clusterfile(cluster.GetClusterName()), cluster)
	if err != nil {
		return err
	}
	pipLine, err := c.GetPipeLine()
	if err != nil {
		return err
	}

	for _, f := range pipLine {
		if err = f(cluster); err != nil {
			return err
		}
	}

	return nil
}
func (c *CreateProcessor) GetPipeLine() ([]func(cluster *v2.Cluster) error, error) {
	var todoList []func(cluster *v2.Cluster) error
	todoList = append(todoList,
		//c.GetPhasePluginFunc(plugin.PhaseOriginally),
		c.CreateCluster,
		c.RunConfig,
		c.MountRootfs,
		//c.GetPhasePluginFunc(plugin.PhasePreInit),
		c.Init,
		c.Join,
		//c.GetPhasePluginFunc(plugin.PhasePreGuest),
		c.RunGuest,
		c.DeleteCluster,
		//c.GetPhasePluginFunc(plugin.PhasePostInstall),
	)
	return todoList, nil
}

func (c *CreateProcessor) CreateCluster(cluster *v2.Cluster) error {
	err := c.RegistryManager.Pull(cluster.Spec.Image)
	if err != nil {
		return err
	}
	img, err := c.ImageManager.Inspect(cluster.Spec.Image)
	if err != nil {
		return err
	}
	c.img = img
	runTime, err := runtime.NewDefaultRuntime(cluster, c.ClusterFile.GetKubeadmConfig(), img)
	if err != nil {
		return fmt.Errorf("failed to init runtime, %v", err)
	}
	c.Runtime = runTime
	c.cManifest, err = c.ClusterManager.Create(cluster.Name, cluster.Spec.Image)
	return err
}

func (c *CreateProcessor) RunConfig(cluster *v2.Cluster) error {
	c.Config = config.NewConfiguration(c.cManifest.MountPoint, c.ClusterFile.GetConfigs())
	return c.Config.Dump(contants.Clusterfile(cluster.GetClusterName()))
}

func (c *CreateProcessor) MountRootfs(cluster *v2.Cluster) error {
	hosts := append(cluster.GetMasterIPList(), cluster.GetNodeIPList()...)
	fs, err := filesystem.NewRootfsMounter(c.cManifest, c.img)
	if err != nil {
		return err
	}

	return fs.MountRootfs(cluster, hosts)
}

func (c *CreateProcessor) Init(cluster *v2.Cluster) error {
	return c.Runtime.Init()
}

func (c *CreateProcessor) Join(cluster *v2.Cluster) error {
	err := c.Runtime.JoinMasters(cluster.GetMasterIPList()[1:])
	if err != nil {
		return err
	}
	err = c.Runtime.JoinNodes(cluster.GetNodeIPList())
	if err != nil {
		return err
	}

	return yaml.MarshalYamlToFile(contants.Clusterfile(cluster.GetClusterName()), cluster)
}

func (c *CreateProcessor) RunGuest(cluster *v2.Cluster) error {
	return c.Guest.Apply(cluster)
}
func (c *CreateProcessor) DeleteCluster(cluster *v2.Cluster) error {
	return c.ClusterManager.Delete(cluster.Name)
}

func NewCreateProcessor(clusterFile clusterfile.Interface) (Interface, error) {
	imgSvc, err := image.NewImageService()
	if err != nil {
		return nil, err
	}

	clusterSvc, err := image.NewDefaultClusterService()
	if err != nil {
		return nil, err
	}

	registrySvc, err := image.NewDefaultRegistryService()
	if err != nil {
		return nil, err
	}

	gs, err := guest.NewGuestManager()
	if err != nil {
		return nil, err
	}

	return &CreateProcessor{
		ClusterFile:     clusterFile,
		ImageManager:    imgSvc,
		ClusterManager:  clusterSvc,
		RegistryManager: registrySvc,
		Guest:           gs,
	}, nil
}
