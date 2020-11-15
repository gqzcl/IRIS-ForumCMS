package app

import (
	"github.com/kataras/iris/v12"
)

//StartIris is a start function
func StartIris() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", showForm)
	app.Listen(":8080")
}

func showForm(ctx iris.Context) {
	ctx.ViewData("message", "Hello World!")
	ctx.View("hello.html")

}
