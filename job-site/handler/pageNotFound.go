package handler

import (
	"context"
	"fmt"
	"job-post/logger"
	"job-post/models"
	"job-post/service/helper"
	"net/http"

	"github.com/go-kit/log/level"
)

func (PageNotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := models.ResponseMessage{
		Data:    "",
		Error:   fmt.Sprintf("%v invalid path", r.URL.Path),
		Code:    404,
		Message: "Invalid Path",
	}

	if err := helper.SendResponse(context.TODO(), w, response); err != nil {
		level.Error(logger.GokitLogger(err)).Log()
		return
	}
}
