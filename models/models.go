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

type Item struct {
	Id       string
	Name     string
	Price    float64
	Currency string
	Rate     float64
	Brand    string
	Model    string
	Amount   int
	Image    string
	Size     string
	Weight   string
}

func (i *Item) ConvertPriceToUAH(price, rate float64) (result float64) {
	return price * rate
}

type Param struct {
	Name  string
	Value string
}
