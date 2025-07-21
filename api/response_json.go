/*
 * @Author: JimZhang
 * @Date: 2025-07-19 13:28:29
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-21 20:25:00
 * @FilePath: /server/api/response_json.go
 * @Description:
 *
 */
package api

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type ResponseJson struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
	Total  int64  `json:"total,omitempty"`
}

func (m ResponseJson) IsEmpty() bool {
	return reflect.DeepEqual(m, ResponseJson{})
}

func HttpResponse(ctx *gin.Context, status int, resp ResponseJson) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
		return
	}
	ctx.AbortWithStatusJSON(status, resp)
}

func buildStatus(resp ResponseJson, nDefaultStatus int) int {

	if resp.Status != 0 {
		return resp.Status
	}
	return nDefaultStatus
}

func OK(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(
		ctx,
		buildStatus(resp, http.StatusOK),
		resp,
	)
}

func Fail(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(
		ctx,
		buildStatus(resp, http.StatusBadRequest),
		resp,
	)
}

func ServerFail(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(
		ctx,
		buildStatus(resp, http.StatusInternalServerError),
		resp,
	)
}
