package nideshop

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"log"
)

// Endpoints are exposed
type Endpoints struct {
	GetEndPoint endpoint.Endpoint
	LoadMainPageEndpoint endpoint.Endpoint
}

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
			log.Fatal("error at load main page end point")
		}
		return loadMainPageDataResponse{indexMainPageData:itf}, nil
	}
}

func MakeGoodCountEndpoint(srv Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		itf, err := srv.LoadGoodCount(ctx)
		if err != nil {
			log.Fatal("error at good count end point")
		}
		return goodCountResponse{goodCount:itf}, nil
	}
}

func MakeGoodCategoryEndpoint(srv Service) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

	}
}
//// Load main page endpoint mapping
//func (e Endpoints) LoadMainPage(ctx context.Context) error {
//	req := LoadMainPageDataRequest{}
//	resp, err := e.LoadMainPageEndpoint(ctx, req)
//	if err != nil {
//		return err
//	}
//	loadMainPageResp := resp.(LoadMainPageDataResponse)
//	if loadMainPageResp.status == "okk" {
//		return nil
//	}
//	return errors.New(loadMainPageResp.status)
//}