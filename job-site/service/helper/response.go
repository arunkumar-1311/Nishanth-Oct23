package helper

import (
	"context"
	"job-post/models"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

// Helps to set header and send response
func SendResponse(ctx context.Context, writer http.ResponseWriter, response interface{}) error {

	writer.Header().Add("content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	resp, ok := response.(models.ResponseMessage)

	if !ok {
		writer.WriteHeader(http.StatusInternalServerError)
		err := kithttp.EncodeJSONResponse(ctx, writer, models.ResponseMessage{Data: "", Error: "Can't assert the response", Code: http.StatusInternalServerError, Message: "Please try again later"})
		return err
	}

	if resp.Token != "" {
		writer.Header().Add("Authorization", resp.Token)
	}
	writer.WriteHeader(resp.Code)
	return kithttp.EncodeJSONResponse(ctx, writer, resp)
}
