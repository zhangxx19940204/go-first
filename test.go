package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"net"
)
func main() {
	app := iris.New()
	app.Use(recover2.New())
	app.Use(logger.New())

	app.Get("/", func(ctx iris.Context){
		ctx.HTML("hello ")
	})

	// 子路由器
	app.PartyFunc("/cpanel", func(child iris.Party) {
		child.Use(myAuthMiddlewareHandler)
		child.Get("/{id:int}" ,func(ctx iris.Context) {
			id:=ctx.Params().Get("id")
			ctx.WriteString(id)
			ctx.HTML("son456")
		})
	})

	// 	OR
	//分组路由器
	cpanel := app.Party("/cpanel",myAuthMiddlewareHandler)
	cpanel.Get("/{id:int}",func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		ctx.WriteString(id)
		ctx.HTML("<h1>son789</h1>")
	})


	//输出html
	//请求方式 GET
	//访问地址:http://localhost:8089/welcome
	app.Handle("GET","/welcome", func(ctx iris.Context) {
		ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
	})

	//输出字符串
	//类似于app.Handle("GET","/ping",[.....])
	//请求方式GET
	//请求地址： http://localhost:8089/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message":"dello iris !"})
	})


	// app.Run(iris.Addr(":8089"))

	// app.Run(iris.Server(&http.Server{Addr:"8089"})) // net/http服务器，iris为其提供服务

	// // 使用自定义net.Listener()
	l, err := net.Listen("tcp4",":8090")
	if err != nil {
		panic(err)
	}
	//添加相应配置
	// app.Configure(iris.WithConfiguration(iris.Configuration{DisableStartupLog:false}))
	//通过绑定路由的方法进行绑定配置
	// app.Run(iris.Listener(l),iris.WithConfiguration(iris.Configuration{
	// 	DisableInterruptHandler:false,
	// 	DisablePathCorrection:false,
	// }))
	// app.Run(iris.Listener(l),iris.WithConfiguration(iris.YAML("./configs/iris.yml")))
	app.Run(iris.Listener(l),iris.WithConfiguration(iris.TOML("./configs/iris.tml")))


}

func myAuthMiddlewareHandler(ctx iris.Context){
	ctx.WriteString("Authentication failed")
	ctx.Next()//继续执行后续的handler(使用这个 PartyFunc 需要加这句话)
}