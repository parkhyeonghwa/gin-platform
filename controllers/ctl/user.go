package ctl

import (
	"github.com/kataras/iris"
)

func Test1(ctx *iris.Context) {
	param1 := ctx.Param("param1")

	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.WriteHTML(200, param1)
}
