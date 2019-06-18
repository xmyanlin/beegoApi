package middleware

import "github.com/astaxie/beego/context"

func Token(ctx *context.Context){
	ctx.Output.JSON("参数不对", false, false )
}
