package service

import (
	"pkhub/models"
)

type StorageInterface interface {
	GetCategories() (listCategories []models.Categories, err error)
	GetGoodsByCategory(categoryId string) (goods map[models.Models][]models.Item, err error)
	GetBrandsByCategory(categoryId string) (listBrands []models.Brands, err error)
	GetGoodsByBrand(brandId, categoryId string) (goods map[models.Models][]models.Item, err error)
}
