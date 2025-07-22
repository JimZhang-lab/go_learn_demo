package middleware

import (
	"net/http"
	"server/api"
	"server/global"
	"server/global/constants"
	"server/model"
	"server/service"
	"server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_INVALID_TOKEN     = 10401 // 无效的token
	ERR_CODE_TOKEN_PARSE       = 10402 // token 解析失败
	ERR_CODE_TOKEN_NOT_MATCHED = 10403 // token 访问者登录时 Token 不一致
	ERR_CODE_TOKEN_EXPIRED     = 10404 // token 过期
	ERR_CODE_TOKEN_NOT_RENEW   = 10405 // token 续期失败
	TOKEN_NAME                 = "Authorization"
	TOKEN_PRIFIX               = "Bearer: "
	RENEW_TOKEN_DURATION       = 10 * 60 * time.Second
)

func tokenErr(c *gin.Context, code int) {
	api.Fail(c, api.ResponseJson{
		Status: http.StatusUnauthorized,
		Code:   code,
		Msg:    "Invalid token",
	})
}

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader(TOKEN_NAME)
		// token 不存在
		if token == "" || !strings.HasPrefix(token, TOKEN_PRIFIX) {
			tokenErr(c, ERR_CODE_INVALID_TOKEN)
			return
		}

		// token 无法解析
		token = token[len(TOKEN_PRIFIX):]
		claims, err := utils.ParseToken(token)
		nUserId := claims.ID
		if err != nil || nUserId == 0 {
			tokenErr(c, ERR_CODE_TOKEN_PARSE)
			return
		}

		// Token 与访问者登录对应的 token 不一致，直接返回
		stUserId := strconv.Itoa(int(nUserId))
		stRedisUserIdKey := strings.Replace(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", stUserId, -1)
		stRedisToken, err := global.RedisClient.Get(stRedisUserIdKey)
		if err != nil || stRedisToken != token {
			tokenErr(c, ERR_CODE_TOKEN_NOT_MATCHED)
			return
		}
		// Token 已过期，直接返回
		expireDuration, err := global.RedisClient.GetExpireDuration(stRedisUserIdKey)
		if err != nil || expireDuration <= 0 {
			tokenErr(c, ERR_CODE_TOKEN_EXPIRED)
			return
		}
		// 小于 10 分钟续签
		if expireDuration.Seconds() < RENEW_TOKEN_DURATION.Seconds() {
			stNewToken, err := service.GenerateAndCacheUserTokenToRedis(int64(nUserId), claims.Name)
			if err != nil {
				tokenErr(c, ERR_CODE_TOKEN_NOT_RENEW)
				return
			}
			c.Header("token", stNewToken)
		}

		// iUser, err := dao.NewUserDao().GetUserById(int64(nUserId))
		// if err != nil {
		// 	tokenErr(c)
		// 	return
		// }
		// c.Set(constants.LOGIN_USER, iUser)
		c.Set(constants.LOGIN_USER, model.LoginUser{
			Id:       int64(nUserId),
			Username: claims.Name,
		})
		c.Next()
	}
}
