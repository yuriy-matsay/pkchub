package models

type Categories struct {
	Id             int    `json:"id"`
	CategoriesName string `json:"categoriesName"`
}

type Brands struct {
	Id         int    `json:"id"`
	BrandsName string `json:"brandsName"`
}

type Models struct {
	Id         int    `json:"id"`
	ModelsName string `json:"modelsName"`
	Image      string `json:"image"`
}

type Goods struct {
	GoodsId   int    `json:"goodsId"`
	IdSite    int    `json:"-"`
	Active    int    `json:"active"`
	Article   string `json:"-"`
	GoodsName string `json:"goodsName"`
	Model     int    `json:"model"`
	Brand     int    `json:"brand"`
	Category  int    `json:"category"`
	Amount    int    `json:"amount"`
	Uah       int    `json:"uah"`
	Usd       int    `json:"usd"`
	Eur       int    `json:"eur"`
	Price     int    `json:"price"`
}

type Product struct {
	Name   string
	Price  int
	Brand  string
	Model  string
	Amount int
}
