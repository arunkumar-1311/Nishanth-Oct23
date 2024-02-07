package router

import (
	"job-post/adaptor"
	"job-post/handler"
	"job-post/middleware"
	"job-post/service"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func Router(db adaptor.Database) error {

	var handlers handler.Endpoints
	handlers.DB = db
	handlers.Authorization = middleware.AcquireMiddleware()
	var service service.Service
	router := mux.NewRouter().StrictSlash(true)

	router.Methods(http.MethodPost).Path("/signup").Handler(kithttp.NewServer(handlers.Register(service), handlers.DecodeRegisterRequest, handlers.EncodeResponse))
	router.Methods(http.MethodPost).Path("/login").Handler(kithttp.NewServer(handlers.Login(service), handlers.DecodeLoginRequest, handlers.EncodeResponse))

	profile := router.PathPrefix("/profile").Subrouter()
	profile.Methods(http.MethodGet).Handler(kithttp.NewServer(handlers.GetProfile(service), handlers.Authorization.Authorization(handlers.DecodeGetProfileRequest), handlers.EncodeResponse))
	profile.Methods(http.MethodPatch).Handler(kithttp.NewServer(handlers.UpdateProfile(service), handlers.Authorization.Authorization(handlers.DecodeUpdateProfileRequest), handlers.EncodeResponse))

	admin := router.Path("/admin/job").Subrouter()
	admin.Methods(http.MethodPost).Handler(kithttp.NewServer(handlers.PostJob(service), handlers.Authorization.Authorization(handlers.DecodePostJobRequest), handlers.EncodeResponse))
	if err := http.ListenAndServe(":8000", router); err != nil {
		return err
	}
	return nil
}
