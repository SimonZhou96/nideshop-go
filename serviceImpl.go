package nideshop

import (
	"context"
	"log"
)

/**to retreive the data from database of the main page, design the interface**/

type service struct {
	repository Repository
}

func (m service) LoadCatalogCurrent(ctx context.Context, parent_id int) (catalogCurrent, error) {
	currentCategoryData,err := m.repository.QueryNideshopCatalogCurrent(parent_id)
	if err != nil {
		return catalogCurrent{}, err
	}
	return catalogCurrent{
		data:   currentCategoryData,
		Errmsg: "",
		Errno:  0,
	},nil
}

func (m service) LoadCatalogIndex(ctx context.Context) (catalogIndex, error) {
	data, err := m.repository.QueryNideshopCatalogIndex()
	if err != nil {
		return catalogIndex{}, err
	}

	return catalogIndex{
		data:   data,
		Errmsg: "",
		Errno:  0,
	}, nil
}


func (m service) LoadTopicLists(page int, size int) (topicList, error) {
	topicsList, errmsg, errno, topicCount := m.repository.QueryNideshopTopicList(page, size)

	var data = data{
		TopicList:        topicsList,
		TopicCount:       topicCount,
		TopicCurrentPage: page,
	}

	return topicList{
		data:   data,
		Errmsg: errmsg,
		Errno:  errno,
	}, nil
}

func (m service) LoadGoodCategory(ctx context.Context) (goodCategory, error) {
	goodsCategory, err := m.repository.QueryNideshopCategory()
	if err != nil {
		log.Fatal("error at service load good count!")
	}
	var data = data{
		GoodsCount:        0,
		NideshopAd:       nil,
		NideshopChannel:  nil,
		NideshopBrand:    nil,
		NideshopTopic:    nil,
		NideshopCategory: nil,
		NideshopHotGoods: nil,
		NideshopNewGoods: nil,
		BrotherCategory:  goodsCategory,
		TopicList:        nil,
	}
	return goodCategory{data: data}, nil
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