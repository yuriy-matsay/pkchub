package models

type Category struct {
	Id           int
	CategoryName string
}

type Brand struct {
	Id        int
	BrandName string
}

type Model struct {
	Id        int
	ModelName string
	Image     string
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
