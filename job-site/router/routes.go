package router

import (
	"fmt"
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
	router.Methods(http.MethodGet).Path("/countries").Handler(kithttp.NewServer(handlers.GetAllCountries(service), handlers.DecodeRequest, handlers.EncodeResponse))
	router.Methods(http.MethodGet).Path("/jobtype").Handler(kithttp.NewServer(handlers.GetAllJobType(service), handlers.DecodeRequest, handlers.EncodeResponse))
	router.Methods(http.MethodGet).Path("/jobs").Handler(kithttp.NewServer(handlers.GetAllJobs(service), handlers.DecodeRequest, handlers.EncodeResponse))
	router.Methods(http.MethodGet).Path("/job/{id}").Handler(kithttp.NewServer(handlers.GetJob(service), handlers.DecodeGetID, handlers.EncodeResponse))

	profile := router.PathPrefix("/profile").Subrouter()
	profile.Methods(http.MethodGet).Handler(kithttp.NewServer(handlers.GetProfile(service), handlers.Authorization.Authorization(handlers.DecodeGetProfileRequest), handlers.EncodeResponse))
	profile.Methods(http.MethodPatch).Handler(kithttp.NewServer(handlers.UpdateProfile(service), handlers.Authorization.Authorization(handlers.DecodeUpdateProfileRequest), handlers.EncodeResponse))

	admin := router.PathPrefix("/admin/job").Subrouter()
	admin.Methods(http.MethodPost).Handler(kithttp.NewServer(handlers.PostJob(service), handlers.Authorization.Authorization(handlers.DecodePostJobRequest), handlers.EncodeResponse))
	admin.Methods(http.MethodPatch).Path("/{id}").Handler(kithttp.NewServer(handlers.UpdateJob(service), handlers.Authorization.Authorization(handlers.DecodeUpdateJobRequest), handlers.EncodeResponse))
	admin.Methods(http.MethodDelete).Path("/{id}").Handler(kithttp.NewServer(handlers.DeleteJob(service), handlers.Authorization.Authorization(handlers.DecodeDeleteJobRequest), handlers.EncodeResponse))

	comments := router.PathPrefix("/comment").Subrouter()
	comments.Methods(http.MethodPost).Path("/{id}").Handler(kithttp.NewServer(handlers.PostComments(service), handlers.Authorization.Authorization(handlers.DecodePostCommentsRequest), handlers.EncodeResponse))
	comments.Methods(http.MethodGet).Path("/{id}").Handler(kithttp.NewServer(handlers.ReadCommentByID(service), handlers.DecodeGetID, handlers.EncodeResponse))
	comments.Methods(http.MethodGet).Path("/post/{id}").Handler(kithttp.NewServer(handlers.ReadCommentByPost(service), handlers.DecodeGetID, handlers.EncodeResponse))
	comments.Methods(http.MethodPatch).Path("/{id}").Handler(kithttp.NewServer(handlers.UpdateCommentByID(service), handlers.Authorization.Authorization(handlers.DecodeUpdateCommentByIDRequest), handlers.EncodeResponse))
	comments.Methods(http.MethodDelete).Path("/{id}").Handler(kithttp.NewServer(handlers.DeleteComment(service), handlers.Authorization.Authorization(handlers.DecodeDeleteCommentRequest), handlers.EncodeResponse))

	fmt.Println("Starting the server......")
	if err := http.ListenAndServe(":8000", router); err != nil {
		return err
	}

	return nil
}
