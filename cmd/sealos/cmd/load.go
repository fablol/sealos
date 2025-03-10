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

	"github.com/spf13/cobra"

	"github.com/labring/sealos/pkg/image/types"
	"github.com/labring/sealos/pkg/utils/strings"

	"github.com/labring/sealos/pkg/image"
)

func newLoadCmd() *cobra.Command {
	var archiveName string
	var loadCmd = &cobra.Command{
		Use:     "load",
		Short:   "load cloud image",
		Example: `sealos load -i kubernetes.tar`,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			registrySvc, err := image.NewImageService()
			if err != nil {
				return err
			}
			return registrySvc.Load(archiveName)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !strings.In(types.DefaultTransport, []string{types.OCIArchive, types.DockerArchive}) {
				return fmt.Errorf("transport parameters must be %s or %s", types.OCIArchive, types.DockerArchive)
			}
			return nil
		},
	}
	loadCmd.Flags().StringVarP(&archiveName, "input", "i", "", "load image from tar archive file")
	loadCmd.Flags().StringVarP(&types.DefaultTransport, "transport", "t", types.OCIArchive, fmt.Sprintf("load image transport from tar archive file.(optional value: %s, %s)", types.OCIArchive, types.DockerArchive))
	return loadCmd
}

func init() {
	rootCmd.AddCommand(newLoadCmd())
}
