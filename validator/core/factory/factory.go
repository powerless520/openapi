/**
* @Author: lik
* @Date: 2021/3/7 16:19
* @Version 1.0
 */
package factory

import (
	"DTCloudAPI/global/errno"
	"DTCloudAPI/global/variable"
	"DTCloudAPI/validator/core/container"
	"DTCloudAPI/validator/core/interf"
	"github.com/gin-gonic/gin"
)

// 表单参数验证器工厂（请勿修改）
func Create(key string) func(context *gin.Context) {

	if value := container.CreateContainersFactory().Get(key); value != nil {
		if val, isOk := value.(interf.ValidatorInterface); isOk {
			return val.CheckParams
		}
	}
	variable.ZapLog.Error(errno.ErrorsValidatorNotExists + ", 验证器模块：" + key)
	return nil
}
