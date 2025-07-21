/*
 * @Author: JimZhang
 * @Date: 2025-07-20 17:13:16
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-20 17:25:34
 * @FilePath: /server/service/dto/host_dto.go
 * @Description:
 *
 */
package dto

type ShutdownHostDTO struct {
	HostIP string `json:"host_ip" binding:"required" message:"主机IP不能为空"`
}
