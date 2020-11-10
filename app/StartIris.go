package app

import (
	"github.com/kataras/iris/v12"
)

func StartIris() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", showForm)
	app.Post("/", handleForm)

	app.Listen(":8080")
}

func showForm(ctx iris.Context) {
	ctx.ViewData("message", "Hello World!")
	ctx.View("hello.html")

}

type formExample struct {
	Colors []string `form:"colors[]"`
}

func handleForm(ctx iris.Context) {
	var form formExample

	err := ctx.ReadForm(&form)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	ctx.JSON(iris.Map{"Colors": form.Colors})
}
