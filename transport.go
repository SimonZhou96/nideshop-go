package nideshop

import (
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)
/**
Our service will be exposed using HTTP.
We are now going to model the accepted HTTP requests and responses.
**/
type User struct{
	Id      int
	Name 	string
	Email	string
	Phone 	string
}

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func MakeHTTPHandler(s Service) http.Handler {
	r := mux.NewRouter()
	loadMainPageEndpoint := MakeLoadMainPageDataEndpoint(s)
	r.Methods("GET").Path("/api/index/index").Handler(httptransport.NewServer(
		loadMainPageEndpoint,
		DecodeLoadMainPageDataRequest,
		EncodeLoadMainPageResponse,
	))

	goodCountEndpoint := MakeGoodCountEndpoint(s)
	r.Methods("GET").Path("/api/goods/count").Handler(httptransport.NewServer(
		goodCountEndpoint,
		DecodeGoodCountRequest,
		EncodeGoodCountResponse,
	))


	return r
}
func DecodeLoadMainPageDataRequest(_ context.Context, r*http.Request) (interface{}, error) {
	var req loadMainPageDataRequest
	return req, nil
}

func EncodeLoadMainPageResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	newRes := response.(loadMainPageDataResponse)
	json.NewEncoder(w).Encode(newRes.indexMainPageData)
	return nil
}

func DecodeGoodCountRequest(_ context.Context, r*http.Request) (interface{}, error) {
	var req goodCountRequest
	return req, nil
}

func EncodeGoodCountResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	newRes := response.(goodCountResponse)
	return json.NewEncoder(w).Encode(newRes.goodCount)
}