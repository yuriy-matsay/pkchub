package service

import (
	"pkhub/models"
)

type StorageInterface interface {
	GetCategories() (listCategories []models.Category, err error)
	GetGoodsByCategory(categoryId string) (goods []map[models.Model][]models.Item, err error)
	GetBrandsByCategory(categoryId string) (listBrands []models.Brand, err error)
	GetGoodsByBrand(brandId, categoryId string) (goods map[models.Model][]models.Item, err error)
}
