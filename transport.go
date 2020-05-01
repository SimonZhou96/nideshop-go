package nideshop

import (
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

	topicListEndpoint := MakeTopicListEndpoint(s)
	r.Methods("GET").Path("/api/topic/list").Handler(httptransport.NewServer(
		topicListEndpoint,
		DecodeTopicListRequest,
		EncodeTopicListResponse,
	))

	catalogIndexEndpoint := MakeCatalogIndexEndpoint(s)
	r.Methods("GET").Path("/api/catalog/index").Handler(httptransport.NewServer(
		catalogIndexEndpoint,
		DecodeCatalogIndexRequest,
		EncodeCatalogIndexResponse,
	))

	catalogCurrentEndpoint := MakeCatalogCurrentEndpoint(s)
	r.Methods("GET").Path("/api/catalog/current").Handler(httptransport.NewServer(
		catalogCurrentEndpoint,
		DecodeCatalogCurrentRequest,
		EncodeCatalogCurrentResponse,
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

func DecodeTopicListRequest(_ context.Context, r*http.Request) (interface{}, error) {
	var req topicListRequest
	req.page, _ = strconv.Atoi(r.URL.Query()["page"][0])
	req.size, _ = strconv.Atoi(r.URL.Query()["size"][0])

	return req, nil
}

func EncodeTopicListResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	newRes := response.(topicListResponse)
	return json.NewEncoder(w).Encode(newRes.topicList)
}

func DecodeCatalogIndexRequest(_ context.Context, r*http.Request)  (interface{}, error) {
	var req catalogIndexRequest
	return req, nil
}

func EncodeCatalogIndexResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	newRes := response.(catalogIndexResponse)
	return json.NewEncoder(w).Encode(newRes.catelogIndex)
}

func DecodeCatalogCurrentRequest(_ context.Context, r*http.Request)  (interface{}, error) {
	var req catalogCurrentRequest
	req.parent_id, _ = strconv.Atoi(r.URL.Query()["id"][0])

	return req, nil
}
func EncodeCatalogCurrentResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	newRes := response.(catalogCurrentResponse)
	return json.NewEncoder(w).Encode(newRes.catalogCurrent)
}