package router

import (
	"fmt"
	"online-purchase/adaptor"
	"online-purchase/handler"
	"online-purchase/middleware"

	"github.com/astaxie/beego"
)

func Routes(db adaptor.Database) {

	var handlers handler.Handlers
	handlers.Database = db
	beego.Get("/brands", handlers.GetBrands)

	user := beego.NewNamespace("/user",
		beego.NSPost("/new", handlers.Register),
		beego.NSPost("", handlers.Login))

	admin := beego.NewNamespace("/admin",
		beego.NSBefore(middleware.Authorization),
		beego.NSPost("/brand", handlers.CreateBrand),
		beego.NSPatch("/brand/:id", handlers.UpdateBrand),
		beego.NSDelete("/brand/:id", handlers.DeleteBrand))
	beego.AddNamespace(user, admin)
	fmt.Println("Starting the Server.....")
	beego.Run("localhost:8000")

}
