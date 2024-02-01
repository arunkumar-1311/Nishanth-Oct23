package router

import (
	"fmt"
	"online-purchase/adaptor"
	"online-purchase/handler"

	"github.com/astaxie/beego"
)

func Routes(db adaptor.Database) {

	var handlers handler.Handlers
	handlers.Database = db

	user := beego.NewNamespace("/user",
		beego.NSPost("/new", handlers.Register),
		beego.NSPost("", handlers.Login))

	beego.AddNamespace(user)
	fmt.Println("Starting the Server.....")
	beego.Run("localhost:8000")

}
