package nideshop

/**Implementation of the interface Repository**/
import (
	"context"
	"database/sql"
	"errors"
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

var RepoErr = errors.New("Unable to handle Repo Request")
type repo struct {
	db *sql.DB
	logger log.Logger
}

type data struct {
	GoodCount int `json:"goodsCount"`

	NideshopAd []models.NideshopAd `json:"banner"`

	NideshopChannel []models.NideshopChannel `json:"channel"`

	NideshopBrand []models.NideshopBrand `json:"brandList"`

	NideshopTopic []models.NideshopTopic `json:"topicList"`

	NideshopCategory []models.NideshopCategory `json:"categoryList"`

	NideshopHotGoods []models.NideshopGoods `json:"hotGoodsList"`

	NideshopNewGoods []models.NideshopGoods `json:"newGoodsList"`

	BrotherCategory []models.NideshopCategory `json:"brotherCategory"`

	CurrentCategory []models.NideshopCategory `json:"currentCategory"`

	ParentCategory models.NideshopCategory `json:"parentCategory"`
}

func (r repo)getNideshopCategoryList(rows *sql.Rows, err error) []models.NideshopCategory {
	if err != nil {
		r.logger.Log("error with querying category list")
	}
	var data []models.NideshopCategory
	for rows.Next() {
		row := models.NideshopCategory{}
		err := rows.Scan(&row.Id, &row.Name, &row.Keywords, &row.FrontDesc, &row.ParentId, &row.SortOrder, &row.ShowIndex, &row.IsShow, &row.BannerUrl, &row.IconUrl, &row.ImgUrl, &row.WapBannerUrl, &row.Level, &row.Type, &row.FrontName)
		if err != nil {
			r.logger.Log("error with single query good category")
		}
		data = append(data, row)
	}
	return data
}
//func (r repo) QueryNideshopGoodsCategory(id string) (goodCategory, error) {
//	sql1 := "SELECT * FROM nideshop_category WHERE ( id = $1 )"
//	//sql2 := "SELECT * FROM `nideshop_category` WHERE ( `id` = 0 ) LIMIT 1" // it is empty set!!
//	sql3 := "SELECT * FROM `nideshop_category` WHERE ( `parent_id` = 0 )"
//	sql3 := "SELECT `category_id` FROM `nideshop_goods` LIMIT 10000"
//	sql4 := "SELECT `parent_id` FROM `nideshop_category` WHERE ( `id` IN (1005007, 1005007, 1005007, 1005007, 1005007, 1005007, 1005007, 1005007, 1005007, 1008001, 1008001, 1008001, 1008001, 1008001, 1008001, 1008001, 1008001, 1008001, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008002, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008008, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008009, 1008015, 1008015, 1008015, 1008015, 1008015, 1008015, 1008015, 1008015, 1008015, 1008016, 1008016, 1008016, 1010002, 1010002, 1010002, 1010002, 1010002, 1010002, 1010002, 1010002, 1010002, 1011004, 1011004, 1011004, 1011004, 1011004, 1011004, 1011004, 1011004, 1011004, 1011004, 1011004, 1011004, 1011004, 1011004, 1011004, 1012001, 1012001, 1012001, 1012001, 1012001, 1012001, 1012001, 1012003, 1012003, 1012003, 1012003, 1012003, 1012003, 1012003, 1012003, 1012003, 1015000, 1015000, 1015000, 1015000, 1015000, 1015000, 1015000, 1015000, 1015000, 1015000, 1015000, 1015000, 1015000, 1015000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1017000, 1020003, 1020003, 1020003, 1020003, 1020003, 1020003, 1020003, 1020003, 1020003, 1032000, 1032000, 1032000, 1032000, 1032000, 1032000, 1032000, 1032000, 1036000, 1036000, 1036000, 1036000, 1036000, 1036000, 1036000, 1036000, 1036000, 1036000, 1036000, 1036000) ) LIMIT 10000"
//	sql5 := "SELECT `id`,`name` FROM `nideshop_category` WHERE ( `id` IN (1005001, 1013001, 1005000, 1005000, 1005000, 1005002, 1005000, 1010000, 1005000, 1008000, 1012000, 1005000, 1005000, 1011000, 1019000, 1005000) ) ORDER BY `sort_order` asc"
//	sql6 := "SELECT `id` FROM `nideshop_category` WHERE ( `parent_id` = '$1' ) LIMIT 10000"
//
//	var goodsCategory goodCategory
//	var currentCategoryList []models.NideshopCategory // id = $1
//	var categoryList []models.NideshopCategory // parent_id = 0
//	var subCategoryList []models.NideshopCategory // parent_id = $2
//
//	rows, err := r.db.Query(sql1, id)
//
//	currentCategoryList = r.getNideshopCategoryList(rows, err)
//
//	rows, err = r.db.Query(sql3, id)
//
//	categoryList = r.getNideshopCategoryList(rows, err)
//
//	rows, err = r.db.Query(sql6, id)
//
//
//}

func (r repo) QueryNideshopNewGoods() ([]models.NideshopGoods, error) {
	sql1 := "SELECT `id`,`name`,`list_pic_url`,`retail_price` FROM `nideshop_goods` WHERE ( `is_new` = 1 ) LIMIT 4;"

	var newGoods []models.NideshopGoods
	rows, err := r.db.Query(sql1)
	if err != nil {

	}
	for rows.Next() {
		row  := models.NideshopGoods{}
		err := rows.Scan(&row.Id, &row.Name, &row.ListPicUrl, &row.RetailPrice)
		if err != nil {
		}
		newGoods = append(newGoods, row)
	}
	return newGoods, nil
}

func (r repo) QueryNideshopHotGoods() ([]models.NideshopGoods, error) {
	// with hot goods and new goods
	sql1 := "SELECT `id`,`name`,`list_pic_url`,`retail_price`,`goods_brief` FROM `nideshop_goods` WHERE ( `is_hot` = 1 ) LIMIT 3;"

	var hotGoods []models.NideshopGoods
	rows, err := r.db.Query(sql1)
	if err != nil {
	}
	for rows.Next() {
		row  := models.NideshopGoods{}
		err := rows.Scan(&row.Id, &row.Name, &row.ListPicUrl, &row.RetailPrice, &row.GoodsBrief)
		if err != nil {
		}
		hotGoods = append(hotGoods, row)
	}
	return hotGoods, nil
}


func (r repo) QueryNideshopGoodsCount() (goodCount, error) {
	sql := "SELECT COUNT(`id`) AS think_count FROM `nideshop_goods` WHERE ( `is_delete` = 0 ) AND ( `is_on_sale` = 1 ) LIMIT 1;"

	rows, err := r.db.Query(sql)
	if err != nil {
	}
	var goodsCount goodCount
	for rows.Next() {
		row := 0
		err := rows.Scan(&row)
		if err != nil {
		}
		goodsCount = goodCount{data{
			GoodCount:        row,
		}}
	}
	return goodsCount, nil
}

func (r repo) QueryNideshopGoods() ([]models.NideshopGoods, error) {
	panic("")
}

func (r repo) DBLoadMainPageData(ctx context.Context) (indexPageJSON, error) {
	hotGoods, err := r.QueryNideshopHotGoods()
	if err != nil {
	}
	newGoods, err := r.QueryNideshopNewGoods()
	if err != nil {
	}
	banners, err := r.QueryNideShopAd()
	if err != nil {
	}
	brandList, err := r.QueryNideshopBrand()
	if err != nil {
	}
	channels, err := r.QueryNideshopChannel()
	if err != nil {
	}
	topicList, err := r.QueryNideshopTopic()
	if err != nil {
	}
	cataList, err := r.QueryNideshopCategory()
	if err != nil {
	}

	var indexOutput = data {
		NideshopAd:       banners,
		NideshopChannel:  channels,
		NideshopBrand:    brandList,
		NideshopTopic:    topicList,
		NideshopCategory: cataList,
		NideshopHotGoods: hotGoods,
		NideshopNewGoods: newGoods,
	}

	return indexPageJSON{data:indexOutput, Errno:0}, nil
}

func (r repo) QueryNideshopChannel() ([]models.NideshopChannel, error) {
	const sql = "SELECT * FROM `nideshop_channel` ORDER BY `sort_order` asc"

	var data []models.NideshopChannel
	rows, err := r.db.Query(sql)
	if err != nil {
	}
	for rows.Next() {
		row := models.NideshopChannel{}
		err := rows.Scan(&row.Id, &row.Name, &row.Url, &row.IconUrl, &row.SortOrder)
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}

func (r repo) QueryNideshopBrand() ([]models.NideshopBrand, error) {
	sql := "SELECT * FROM `nideshop_brand` WHERE ( `is_new` = 1 ) ORDER BY `new_sort_order` asc LIMIT 4;"

	rows, err := r.db.Query(sql)
	var data []models.NideshopBrand
	if err != nil {
	}
	for rows.Next() {
		row := models.NideshopBrand{}
		err := rows.Scan(&row.Id, &row.Name, &row.ListPicUrl, &row.SimpleDesc, &row.PicUrl, &row.SortOrder, &row.IsShow, &row.FloorPrice, &row.AppListPicUrl, &row.IsNew, &row.NewPicUrl, &row.NewSortOrder)
		if err != nil {
			return nil,err
		}
		data = append(data, row)
	}
	return data, nil
}

func (r repo) QueryNideshopTopic() ([]models.NideshopTopic, error) {
	sql := "SELECT * FROM `nideshop_topic` LIMIT 3"
	var data []models.NideshopTopic
	rows, err := r.db.Query(sql)
	if err != nil {
	}
	for rows.Next() {
		row := models.NideshopTopic{}
		err := rows.Scan(&row.Id, &row.Title, &row.Content, &row.Avatar, &row.ItemPicUrl, &row.Subtitle, &row.TopicCategoryId, &row.PriceInfo, &row.ReadCount, &row.ScenePicUrl, &row.TopicTemplateId, &row.TopicTagId, &row.SortOrder, &row.IsShow)
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}

func (r repo) QueryNideshopCategory() ([]models.NideshopCategory, error) {
	sql := `SELECT * FROM nideshop_category WHERE ( parent_id = 0 ) AND ( name != '推荐' )`

	var data []models.NideshopCategory
	rows, err := r.db.Query(sql)
	if err != nil {
	}
	for rows.Next() {
		row := models.NideshopCategory{}
		err := rows.Scan(&row.Id, &row.Name, &row.Keywords, &row.FrontDesc, &row.ParentId, &row.SortOrder, &row.ShowIndex, &row.IsShow, &row.BannerUrl, &row.IconUrl, &row.ImgUrl, &row.WapBannerUrl, &row.Level, &row.Type, &row.FrontName)
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}

/**
	AdPositionId int    `json:"ad_position_id"`
	Content      string `json:"content"`
	Enabled      int    `json:"enabled"`
	EndTime      int    `json:"end_time"`
	Id           int    `json:"id"`
	ImageUrl     string `json:"image_url"`
	Link         string `json:"link"`
	MediaType    int    `json:"media_type"`
	Name         string `json:"name"`
**/

func (r repo) QueryNideShopAd() ([]models.NideshopAd, error) {
	const sql1 = `
		SELECT * FROM nideshop_ad WHERE ( ad_position_id = 1 )
		`

	var data []models.NideshopAd
	rows, err := r.db.Query(sql1)
	if err != nil {
	}
	for rows.Next() {
		row := models.NideshopAd{}
		err := rows.Scan(&row.Id, &row.AdPositionId, &row.MediaType, &row.Name, &row.Link, &row.ImageUrl, &row.Content, &row.EndTime, &row.Enabled)
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}


func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}
