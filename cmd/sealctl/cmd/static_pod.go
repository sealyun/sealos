// Copyright © 2021 sealos.
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

package cmd

import (
	"fmt"
	"github.com/fanux/sealos/cmd/sealctl/boot"
	"github.com/fanux/sealos/pkg/ipvs"
	"github.com/fanux/sealos/pkg/types/contants"
	"github.com/fanux/sealos/pkg/utils/logger"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
)


func NewStaticPodCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "static-pod",
		Short: "generator static pod",
		//Run: func(cmd *cobra.Command, args []string) {
		//
		//},
	}
	// check route for host
	cmd.AddCommand(NewLvscareCmd())
	cmd.PersistentFlags().StringVar(&flag.StaticPod.staticPodPath, "path", "/etc/kubernetes/manifests", "default kubernetes static pod path")
	return cmd
}

func NewLvscareCmd() *cobra.Command {
	var vip, image, name string
	var masters []string
	var printBool bool
	var cmd = &cobra.Command{
		Use:   "lvscare",
		Short: "generator lvscare static pod file",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(masters) == 0 {
				return fmt.Errorf("master not allow empty")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fileName := fmt.Sprintf("%s.yaml", name)
			yaml, err := ipvs.LvsStaticPodYaml(vip, masters, image, name)
			if err != nil {
				logger.Error(err)
				os.Exit(1)
			}
			if printBool {
				println(yaml)
				return
			}
			logger.Debug("lvscare static pod yaml is %s", yaml)
			if err = boot.InitRootDirectory([]string{flag.StaticPod.staticPodPath}); err != nil {
				logger.Error("init dir is error: %v", err)
				os.Exit(1)
			}
			err = ioutil.WriteFile(path.Join(flag.StaticPod.staticPodPath, fileName), []byte(yaml), 0755)
			if err != nil {
				logger.Error(err)
				os.Exit(1)
			}
			logger.Info("generator lvscare static pod is success")
		},
	}
	// manually to set host via gateway
	cmd.Flags().StringVar(&vip, "vip", "10.103.97.2", "default vip IP")
	cmd.Flags().StringVar(&name, "name", contants.LvsCareStaticPodName, "generator lvscare static pod name")
	cmd.Flags().StringVar(&image, "image", contants.DefaultLvsCareImage, "generator lvscare static pod image")
	cmd.Flags().StringSliceVar(&masters, "masters", nil, "generator masters addrs")
	cmd.Flags().BoolVar(&printBool, "print", false, "is print yaml")
	return cmd
}

func init() {
	rootCmd.AddCommand(NewStaticPodCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostnameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostnameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
