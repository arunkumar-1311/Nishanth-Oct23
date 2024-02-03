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
	beego.ErrorHandler("404", handlers.PageNotFound)
	beego.Get("/brands", handlers.GetBrands)
	beego.Get("/rams", handlers.GetAllRAMs)
	beego.Get("/brand/:id", handlers.GetBrandByID)
	beego.Get("/ram/:id", handlers.GetRamByID)
	beego.Get("/orders", handlers.GetAllOrders)

	user := beego.NewNamespace("/user",
		beego.NSPost("/signup", handlers.Register),
		beego.NSPost("/login", handlers.Login))

	admin := beego.NewNamespace("/admin",
		beego.NSBefore(middleware.Authorization),
		beego.NSPost("/brand", handlers.CreateBrand),
		beego.NSPatch("/brand/:id", handlers.UpdateBrand),
		beego.NSDelete("/brand/:id", handlers.DeleteBrand),
		beego.NSPost("/ram", handlers.CreateRAM),
		beego.NSPatch("/ram/:id", handlers.UpdateRAM),
		beego.NSDelete("/ram/:id", handlers.DeleteRAM),
		beego.NSGet("/inactive/orders", handlers.GetInactiveOrders),
		beego.NSGet("/orderstatus", handlers.GetAllOrderStatus),
		beego.NSPatch("/order/:id", handlers.UpdateStatus))

	order := beego.NewNamespace("/order",
		beego.NSBefore(middleware.Authorization),
		beego.NSPost("", handlers.CreateOrder),
		beego.NSGet("/:id", handlers.GetOrderByID),
		beego.NSDelete("/cancel/:id", handlers.CancelOrder))

	beego.AddNamespace(user, admin, order)
	fmt.Println("Starting the Server.....")
	beego.Run("localhost:8000")

}
