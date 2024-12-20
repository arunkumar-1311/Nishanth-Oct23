package router

import (
	"fmt"
	"job-post/adaptor"
	"job-post/handler"
	"job-post/service"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func Router(db adaptor.Database) error {

	var handlers handler.Endpoints
	handlers.DB, handlers.MW.DB = db, db
	api := handler.AcqurieAPI(handlers)

	var service service.Service
	router := mux.NewRouter().StrictSlash(true)

	router.NotFoundHandler = handler.PageNotFound{}
	router.Methods(http.MethodPost).Path("/{role:(?:admin|user)}/signup").Handler(kithttp.NewServer(api.Register(service), api.DecodeRegisterRequest, api.EncodeResponse))
	router.Methods(http.MethodPost).Path("/login").Handler(kithttp.NewServer(api.Login(service), api.DecodeLoginRequest, api.EncodeResponse))
	router.Methods(http.MethodGet).Path("/countries").Handler(kithttp.NewServer(api.GetAllCountries(service), api.DecodeRequest, api.EncodeResponse))
	router.Methods(http.MethodGet).Path("/jobtype").Handler(kithttp.NewServer(api.GetAllJobType(service), api.DecodeRequest, api.EncodeResponse))
	router.Methods(http.MethodGet).Path("/jobs").Handler(kithttp.NewServer(api.GetAllJobs(service), api.DecodeGetAllJobsRequest, api.EncodeResponse))
	router.Methods(http.MethodGet).Path("/job/{id}").Handler(kithttp.NewServer(api.GetJob(service), api.DecodeGetID, api.EncodeResponse))

	profile := router.PathPrefix("/profile").Subrouter()
	profile.Methods(http.MethodPatch).Handler(kithttp.NewServer(api.UpdateProfile(service), handlers.MW.Authorization(api.DecodeUpdateProfileRequest), api.EncodeResponse))
	profile.Methods(http.MethodDelete).Handler(kithttp.NewServer(api.DeleteProfile(service), handlers.MW.Authorization(api.DecodeClaims), api.EncodeResponse))
	profile.Methods(http.MethodGet).Path("/logout").Handler(kithttp.NewServer(api.LogOut(service), handlers.MW.Authorization(api.DecodeClaims), api.EncodeResponse))
	profile.Methods(http.MethodGet).Handler(kithttp.NewServer(api.GetProfile(service), handlers.MW.Authorization(api.DecodeClaims), api.EncodeResponse))

	admin := router.PathPrefix("/admin/job").Subrouter()
	admin.Methods(http.MethodPost).Handler(kithttp.NewServer(api.PostJob(service), handlers.MW.Authorization(api.DecodePostJobRequest), api.EncodeResponse))
	admin.Methods(http.MethodPatch).Path("/{id}").Handler(kithttp.NewServer(api.UpdateJob(service), handlers.MW.Authorization(api.DecodeUpdateJobRequest), api.EncodeResponse))
	admin.Methods(http.MethodDelete).Path("/{id}").Handler(kithttp.NewServer(api.DeleteJob(service), handlers.MW.Authorization(api.DecodeDeleteJobRequest), api.EncodeResponse))

	comments := router.PathPrefix("/comment").Subrouter()
	comments.Methods(http.MethodPost).Path("/{id}").Handler(kithttp.NewServer(api.PostComments(service), handlers.MW.Authorization(api.DecodePostCommentsRequest), api.EncodeResponse))
	comments.Methods(http.MethodGet).Path("/{id}").Handler(kithttp.NewServer(api.ReadCommentByID(service), api.DecodeGetID, api.EncodeResponse))
	comments.Methods(http.MethodGet).Path("/post/{id}").Handler(kithttp.NewServer(api.ReadCommentByPost(service), api.DecodeGetID, api.EncodeResponse))
	comments.Methods(http.MethodPatch).Path("/{id}").Handler(kithttp.NewServer(api.UpdateCommentByID(service), handlers.MW.Authorization(api.DecodeUpdateCommentByIDRequest), api.EncodeResponse))
	comments.Methods(http.MethodDelete).Path("/{id}").Handler(kithttp.NewServer(api.DeleteComment(service), handlers.MW.Authorization(api.DecodeDeleteCommentRequest), api.EncodeResponse))

	fmt.Println("Starting the server......")
	if err := http.ListenAndServe(":8000", router); err != nil {
		return err
	}

	return nil
}
