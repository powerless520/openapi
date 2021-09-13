/**
* @Author: lik
* @Date: 2021/3/5 11:41
* @Version 1.0
 */
package main

import (
	"DTCloudAPI/global/variable"
	_ "DTCloudAPI/initialization"
	"DTCloudAPI/routers"
)

func main() {
	router := routers.InitApiRouter()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Api.Port"))
}
