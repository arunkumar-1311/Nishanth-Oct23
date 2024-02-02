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
	beego.Get("/rams", handlers.GetAllRAMs)
	beego.Get("/brand/:id", handlers.GetBrandByID)
	beego.Get("/ram/:id", handlers.GetRamByID)

	user := beego.NewNamespace("/user",
		beego.NSPost("/new", handlers.Register),
		beego.NSPost("", handlers.Login))

	admin := beego.NewNamespace("/admin",
		beego.NSBefore(middleware.Authorization),
		beego.NSPost("/brand", handlers.CreateBrand),
		beego.NSPatch("/brand/:id", handlers.UpdateBrand),
		beego.NSDelete("/brand/:id", handlers.DeleteBrand),
		beego.NSPost("/ram", handlers.CreateRAM),
		beego.NSPatch("/ram/:id", handlers.UpdateRAM),
		beego.NSDelete("/ram/:id", handlers.DeleteRAM),
		beego.NSGet("/orders", handlers.GetAllOrders),
		beego.NSGet("/orderstatus", handlers.GetAllOrderStatus))

	order := beego.NewNamespace("/order",
		beego.NSBefore(middleware.Authorization),
		beego.NSPost("", handlers.CreateOrder),
		beego.NSGet("/:id", handlers.GetOrderByID),
		beego.NSPatch("/:id", handlers.CancelOrder))

	beego.AddNamespace(user, admin, order)
	fmt.Println("Starting the Server.....")
	beego.Run("localhost:8000")

}
