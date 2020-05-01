package nideshop

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	_ "github.com/go-kit/kit/log"
	log2 "log"
)

/**
Make...() will receive the service as argument, then use a type assertion to "force"
the request type to a specific one and use it to call the service method for it, it will
be used in the main.go file
**/

// MakeGetLoadMainPageDataEndpoint returns the response from our service "get"
func MakeLoadMainPageDataEndpoint(srv Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		itf, err := srv.LoadMainPageData(ctx)
		if err != nil {
			log2.Fatal(err)
		}
		return loadMainPageDataResponse{indexMainPageData:itf}, nil
	}
}

func MakeGoodCountEndpoint(srv Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		itf, err := srv.LoadGoodCount(ctx)
		if err != nil {
			log2.Fatal(err)
		}
		return goodCountResponse{goodCount:itf}, nil
	}
}


func MakeTopicListEndpoint(srv Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		topicsList, err := srv.LoadTopicLists(request.(topicListRequest).page,request.(topicListRequest).size)
		if err != nil {
			log2.Fatal(err)
		}
		return topicListResponse{topicList:topicsList}, nil
	}
}

func MakeCatalogIndexEndpoint(srv Service) endpoint.Endpoint {
	
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		catelogIndexStruct, err := srv.LoadCatalogIndex(ctx)
		if err != nil {
			log2.Fatal(err)
		}
		return catalogIndexResponse{catelogIndex:catelogIndexStruct},nil
	}
}
func MakeCatalogCurrentEndpoint(srv Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		parent_id := request.(catalogCurrentRequest).parent_id
		catalogCurrentStruct, err := srv.LoadCatalogCurrent(ctx, parent_id)
		if err != nil {
			log2.Fatal(err)
		}
		return catalogCurrentResponse{catalogCurrentStruct}, nil
	}
}
