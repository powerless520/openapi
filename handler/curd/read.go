/**
* @Author: lik
* @Date: 2021/3/10 9:30
* @Version 1.0
 */
package curd

import (
	"DTCloudAPI/global/constant"
	"DTCloudAPI/service/curd"
	"github.com/gin-gonic/gin"
	"strings"
)

type ReadData struct {
}

func (c *ReadData) PublicRead(context *gin.Context){
	fields := context.GetString(constant.ValidatorPrefix + "fields")
	ids := context.GetString(constant.ValidatorPrefix + "ids")
	page := context.GetInt64(constant.ValidatorPrefix + "page")
	limit := context.GetInt64(constant.ValidatorPrefix + "limit")
	order := context.GetString(constant.ValidatorPrefix + "order")
	table := strings.Replace(context.GetString(constant.ValidatorPrefix + "model"), ".", "_", -1)


	curd.CreateUserCurdFactory().PublicRead(fields,ids,page,limit,order,table)



}
