package models

type Categories struct {
	Id             int
	CategoriesName string
}

type Brands struct {
	Id         int
	BrandsName string
}

type Models struct {
	Id         int
	ModelsName string
	Image      string
}

type Goods struct {
	GoodsId   int
	IdSite    int
	Active    int
	Article   string
	GoodsName string
	Model     int
	Brand     int
	Category  int
	Amount    int
	Uah       int
	Usd       int
	Eur       int
	Price     int
}

type Item struct {
	Name   string
	Price  int
	Brand  string
	Model  string
	Amount int
}
