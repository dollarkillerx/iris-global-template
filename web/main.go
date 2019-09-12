package main

import (
	"github.com/kataras/iris"
	"iris-global-template/config"
	"iris-global-template/web/register"
)

func main() {
	app := register.Iris()

	app.Run(iris.Addr(config.MyConfig.App.Host), iris.WithCharset("UTF-8"))
}
