package router

import (
	"anime_com/controllers"
	"github.com/kataras/iris/v12"
)

func InitRouter(app *iris.Application) {
	animeCom := app.Party("/anime_com")
	{
		animeCom.Get("/find", controllers.FindAnimeCom)
		animeCom.Post("/del", controllers.DelAnimeCom)
		animeCom.Post("/add", controllers.AddAnimeCom)
		animeCom.Post("/update", controllers.UpdateAnimeCom)
	}

}
