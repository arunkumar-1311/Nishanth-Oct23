package service

import (
	"encoding/json"
	"online-purchase/models"

	"github.com/astaxie/beego/context"
)

// Helps to set header and send response
func SendResponse(c *context.Context, status int, err string, message, method string, data ...interface{}) error {
	
	c.ResponseWriter.WriteHeader(status)
	resp := models.Message{Data: data, Error: err, Code: status, Message: message}
	if err := json.NewEncoder(c.ResponseWriter).Encode(&resp); err != nil {
		return err
	}
	return nil
}
