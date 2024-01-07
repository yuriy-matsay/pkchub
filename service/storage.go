package service

import (
	"pkhub/models"
)

type StorageInterface interface {
	GetCurrencies() (currencies string)
	GetCategories() (listCategories []models.Category, err error)
	GetGoodsByCategory(categoryId string) (goods []map[models.Model][]models.Item, err error)
	GetBrandsByCategory(categoryId string) (listBrands []models.Brand, err error)
	GetGoodsByBrand(brandId, categoryId string) (goods map[models.Model][]models.Item, err error)
	GetItemParams(itemId string) (data []models.Param)
	GetItemInfo(itemId string) (data models.Item)
}
