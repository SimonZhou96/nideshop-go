package nideshop

import (
	"context"
	"nideshop-go/models"
)

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
}

// Service that provide the service to load main page
type Service interface {
	LoadMainPageData(ctx context.Context) (indexPageJSON, error)
	LoadGoodCount(ctx context.Context) (goodCount, error)
	LoadGoodCategory(ctx context.Context) (goodCategory, error)
}




