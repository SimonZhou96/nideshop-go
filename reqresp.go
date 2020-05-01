package nideshop

import "nideshop-go/models"

type loadMainPageDataRequest struct {
}

type loadMainPageDataResponse struct {
	indexMainPageData indexPageJSON
}
type goodCountRequest struct {

}
type goodCountResponse struct {
	goodCount goodCount
}

type goodCategoryRequest struct {
	id int
	parent_id int
}

type goodCategoryResponse struct {
	brotherCategory []models.NideshopCategory
	currentCategory models.NideshopCategory
	parentCategory []models.NideshopCategory
}

type topicListRequest struct {
	page int
	size int
}

type topicListResponse struct {
	topicList topicList
}

type catalogIndexRequest struct {
}

type catalogIndexResponse struct {
	catelogIndex catalogIndex
}

type catalogCurrentRequest struct {
	parent_id int
}

type catalogCurrentResponse struct {
	catalogCurrent catalogCurrent
}