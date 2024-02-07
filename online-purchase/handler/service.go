package handler

import (
	"online-purchase/adaptor"

	"github.com/astaxie/beego"
)

type Handlers struct {
	adaptor.Database
	beego.Controller
}
