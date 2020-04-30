package nideshop

import (
	"context"
	"log"
)

/**to retreive the data from database of the main page, design the interface**/

type service struct {
	repository Repository
}

func (m service) LoadGoodCategory(ctx context.Context) (goodCategory, error) {
	goodsCategory, err := m.repository.QueryNideshopCategory()
	if err != nil {
		log.Fatal("error at service load good count!")
	}
	return goodsCategory, nil
}

func (m service) LoadGoodCount(ctx context.Context) (goodCount, error) {
	goodCount, err := m.repository.QueryNideshopGoodsCount()
	if err != nil {
		log.Fatal("error at service load good count!")
	}
	return goodCount, nil
}

func (m service) LoadMainPageData(ctx context.Context) (indexPageJSON, error) {
	data, err := m.repository.DBLoadMainPageData(ctx)
	if err != nil {
		log.Fatal("error at service load main page!")
	}
	return data, nil
}

func NewLoadMainPageService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}