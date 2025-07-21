/*
 * @Author: JimZhang
 * @Date: 2025-07-20 17:12:02
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-20 17:50:04
 * @FilePath: /server/service/host_service.go
 * @Description:
 *
 */
package service

import (
	"context"
	"fmt"
	"server/service/dto"

	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/spf13/viper"
)

var hostService *HostService

type HostService struct {
	BaseService
}

func NewHostService() *HostService {
	if hostService == nil {
		hostService = &HostService{}
	}
	return hostService
}
func (m *HostService) Shutdown(iShutdownHostDTO dto.ShutdownHostDTO) error {
	var errResult error

	stHostIP := iShutdownHostDTO.HostIP
	fmt.Println("shutdown host:", stHostIP)

	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "ssh",
		User:       viper.GetString("ansible.user.name"),
	}
	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  fmt.Sprintf("%s,", stHostIP),
		ModuleName: "shutdown",
		Args:       viper.GetString("ansible.shutdownHost.args"),
		ExtraVars: map[string]any{
			"ansible_password": viper.GetString("ansible.user.password"),
		},
	}

	adhoc := &adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
		StdoutCallback:    "oneline",
	}
	errResult = adhoc.Run(context.TODO())
	return errResult
}
