package main

import (
	"iris-global-template/config"
	"iris-global-template/web/register"
	"github.com/kataras/iris"
)

func main() {
	app := register.Iris()

	app.Run(iris.Addr(config.MyConfig.App.Host),iris.WithCharset("UTF-8"))
}
