package buildah

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/containers/common/pkg/config"
	"github.com/containers/storage/pkg/homedir"
	"github.com/containers/storage/pkg/unshare"
	"github.com/containers/storage/types"

	"github.com/labring/sealos/pkg/utils/file"
	"github.com/labring/sealos/pkg/utils/logger"
)

var (
	DefaultConfigFile                  string
	DefaultSignaturePolicyPath         = config.DefaultSignaturePolicyPath
	DefaultRootlessSignaturePolicyPath = "containers/policy.json"
	DefaultGraphRoot                   = "/var/lib/containers/storage"
	DefaultRegistriesFilePath          = "/etc/containers/registries.conf"
	DefaultRootlessRegistriesFilePath  = "containers/registries.conf"
)

func init() {
	var err error
	DefaultConfigFile, err = types.DefaultConfigFile(unshare.IsRootless())
	if err != nil {
		logger.Fatal(err)
	}
	if unshare.IsRootless() {
		configHome, err := homedir.GetConfigHome()
		if err != nil {
			logger.Fatal(err)
		}
		DefaultSignaturePolicyPath = filepath.Join(configHome, DefaultRootlessSignaturePolicyPath)
		DefaultRegistriesFilePath = filepath.Join(configHome, DefaultRootlessRegistriesFilePath)
	}
}

const defaultPolicy = `
{
    "default": [
        {
            "type": "insecureAcceptAnything"
        }
    ],
    "transports":
        {
            "docker-daemon":
                {
                    "": [{"type":"insecureAcceptAnything"}]
                }
        }
}
`

const defaultRegistries = `unqualified-search-registries = ["docker.io"]

[[registry]]
prefix = "docker.io/labring"
location = "docker.io/labring"
`

func SetupContainerPolicy() error {
	return writeFileIfNotExists(DefaultSignaturePolicyPath, []byte(defaultPolicy))
}

func SetupRegistriesFile() error {
	return writeFileIfNotExists(DefaultRegistriesFilePath, []byte(defaultRegistries))
}

func writeFileIfNotExists(filename string, data []byte) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		err = file.WriteFile(filename, data)
	}
	return err
}

func MaybeReexecUsingUserNamespace() error {
	if !unshare.IsRootless() || strings.ToLower(os.Getenv("DISABLE_AUTO_ROOTLESS")) == "true" {
		return nil
	}
	if _, present := os.LookupEnv("BUILDAH_ISOLATION"); !present {
		if err := os.Setenv("BUILDAH_ISOLATION", "rootless"); err != nil {
			return fmt.Errorf("error setting BUILDAH_ISOLATION=rootless in environment: %v", err)
		}
	}

	// force reexec using the configured ID mappings
	unshare.MaybeReexecUsingUserNamespace(true)
	return nil
}

type Setter func() error

var defaultSetters = []Setter{
	MaybeReexecUsingUserNamespace,
	SetupContainerPolicy,
	SetupRegistriesFile,
}

func TrySetupWithDefaults(setters ...Setter) error {
	if len(setters) == 0 {
		setters = defaultSetters
	}
	for i := range setters {
		if err := setters[i](); err != nil {
			return err
		}
	}
	return nil
}
