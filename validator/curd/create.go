/**
* @Author: lik
* @Date: 2021/3/9 11:17
* @Version 1.0
 */
package curd

import (
	"DTCloudAPI/global/constant"
	"DTCloudAPI/global/token"
	"DTCloudAPI/handler/curd"
	"DTCloudAPI/util/json_params"
	"DTCloudAPI/util/response"
	"github.com/gin-gonic/gin"
)

type PublicCreate struct {
}

func (c PublicCreate) CheckParams(context *gin.Context) {
	dat := make(map[string]interface{})
	err := json_params.QueryParams(context.Request.URL.Query(), &dat)

	if err != nil {
		response.Fail(context, constant.CurdUpdateFailCode, constant.ParamsFailMsg, "")
		return
	}

	accessToken := json_params.SliceValues(dat["access_token"].([]interface{}))

	userTokenFactory := token.CreateUserFactory()
	customClaims, _ := userTokenFactory.UserJwt.ParseToken(accessToken)

	if customClaims == nil {
		response.Fail(context, constant.CurdUpdateFailCode, constant.ErrorsTokenInvalid, "")
		return
	}

	bol := (&curd.CreateData{}).PublicCreate(dat, customClaims)
	if bol{
		response.Success(context, constant.CurdStatusOkMsg, nil)
		return
	}
	response.Fail(context, constant.CurdCreatFailCode, constant.CurdCreatFailMsg, nil)
	return

}
