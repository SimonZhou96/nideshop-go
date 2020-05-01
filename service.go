package nideshop

import (
	"context"
	"database/sql"
	"github.com/go-kit/kit/log"
	"nideshop-go/models"
)

type indexPageJSON struct {
	data  `json:"data"`
	Errno int `json:"errno"`
}
type goodCount struct {
	data `json:"data"`
}

type goodCategory struct {
	data `json:"data"`
}

type catalogIndex struct {
	data `json:"data"`
	Errmsg string `json:"errmsg"`
	Errno int `json:"errno"`
}
type topicList struct {
	data `json:"data"`
	Errmsg string `json:"errmsg"`
	Errno int `json:"errno"`
}
type repo struct {
	db *sql.DB
	logger log.Logger
}

type catalogCurrent struct {
	data `json:"data"`
	Errmsg string `json:"errmsg"`
	Errno int `json:"errno"`
}

type currentCategoryStruct struct {
	BannerUrl    string `json:"banner_url"`
	FrontDesc    string `json:"front_desc"`
	FrontName    string `json:"front_name"`
	IconUrl      string `json:"icon_url"`
	Id           int    `json:"id"`
	ImgUrl       string `json:"img_url"`
	IsShow       int    `json:"is_show"`
	Keywords     string `json:"keywords"`
	Level        string `json:"level"`
	Name         string `json:"name"`
	ParentId     int    `json:"parent_id"`
	ShowIndex    int    `json:"show_index"`
	SortOrder    int    `json:"sort_order"`
	Type         int    `json:"type"`
	WapBannerUrl string `json:"wap_banner_url"`
	SubCategoryList []models.NideshopCategory `json:"subCategoryList"`
}
type data struct {

	NideshopAd []models.NideshopAd `json:"banner"`

	NideshopChannel []models.NideshopChannel `json:"channel"`

	NideshopBrand []models.NideshopBrand `json:"brandList"`

	NideshopTopic []models.NideshopTopic `json:"topicList"`

	NideshopCategory []models.NideshopCategory `json:"categoryList"`

	NideshopHotGoods []models.NideshopGoods `json:"hotGoodsList"`

	NideshopNewGoods []models.NideshopGoods `json:"newGoodsList"`

	BrotherCategory []models.NideshopCategory `json:"brotherCategory"`

	CurrentCategory currentCategoryStruct `json:"currentCategory"`

	ParentCategory models.NideshopCategory `json:"parentCategory"`

	TopicList []models.NideshopTopic `json:"data"`

	TopicCount int `json:"count"`

	TopicCurrentPage int `json:"currentPage"`
	
	GoodsCount int `json:"goodsCount"`
}

type Repository interface {
	DBLoadMainPageData(ctx context.Context) (indexPageJSON, error)
	QueryNideShopAd()([]models.NideshopAd, error)
	QueryNideshopChannel()([]models.NideshopChannel, error)
	QueryNideshopBrand()([]models.NideshopBrand, error)
	QueryNideshopTopic()([]models.NideshopTopic, error)
	QueryNideshopCategory()([]models.NideshopCategory, error)
	QueryNideshopGoods()([]models.NideshopGoods, error)
	QueryNideshopGoodsCount()(goodCount, error)
	QueryNideshopNewGoods()([]models.NideshopGoods, error)
	QueryNideshopHotGoods()([]models.NideshopGoods, error)
	QueryNideshopGoodsCategory(id string)([]models.NideshopCategory, error)
	QueryNideshopTopicList(page int, size int)([]models.NideshopTopic, string, int, int)
	QueryNideshopCatalogIndex()(data, error)
	QueryNideshopCatalogCurrent(parent_id int)(data, error)
}

// Service that provide the service to load main page
type Service interface {
	LoadMainPageData(ctx context.Context) (indexPageJSON, error)
	LoadGoodCount(ctx context.Context) (goodCount, error)
	LoadGoodCategory(ctx context.Context) (goodCategory, error)
	LoadTopicLists(page int, size int) (topicList, error)
	LoadCatalogIndex(ctx context.Context) (catalogIndex, error)
	LoadCatalogCurrent(ctx context.Context, parent_id int) (catalogCurrent, error)
}




