package sqlite

import (
	"database/sql"
	"log"
	"os"
	"pkhub/models"

	_ "modernc.org/sqlite"
)

type Sqlite struct {
	db *sql.DB
}

func NewSqliteDB() *Sqlite {

	sqlitedb, err := OpenDatabase()
	if err != nil {
		panic(err)
	}
	storage := &Sqlite{
		db: sqlitedb,
	}
	return storage
}

func OpenDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", os.Getenv("DB_PATH"))
	if err == nil {
		log.Print("Opened database")
	}

	// if err := migrationsUp(db); err != nil {
	// 	log.Print("migration failed")
	// }

	return
}

func (s *Sqlite) GetCategories() (listCategories []models.Categories, err error) {
	rows, err := s.db.Query(`
		SELECT DISTINCT categories.id, categories.categoriesName
		FROM goods
		INNER JOIN categories ON goods.category = categories.id
		WHERE goods.active = 1
		`)
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		category := models.Categories{}
		if err = rows.Scan(&category.Id, &category.CategoriesName); err != nil {
			log.Print(err)
			return
		}
		listCategories = append(listCategories, category)
	}
	return
}

func (s *Sqlite) GetGoodsByCategory(categoryId string) (goods map[models.Models][]models.Item, err error) {

	listModel := s.getListModelsByCategory(categoryId)

	goods = make(map[models.Models][]models.Item)

	for _, i := range listModel {
		rows, _ := s.db.Query(`
			SELECT goods.goodsName, goods.uah, goods.amount, brands.brandsName, models.modelsName
			FROM goods
			INNER JOIN brands ON goods.brand = brands.id
			INNER JOIN models ON goods.model = models.id
			WHERE goods.model = ? AND goods.active = 1 AND goods.amount > 0
			`, i.Id)

		products := []models.Item{}
		for rows.Next() {
			product := models.Item{}
			if err = rows.Scan(&product.Name, &product.Price, &product.Amount, &product.Brand, &product.Model); err != nil {
				log.Print(err)
				return
			}

			products = append(products, product)
		}
		if len(products) > 0 {
			goods[i] = products
		}
	}
	return
}

func (s *Sqlite) GetBrandsByCategory(categoryId string) (listBrands []models.Brands, err error) {
	rows, err := s.db.Query(`
		SELECT DISTINCT goods.brand, brands.brandsName
		FROM goods
		INNER JOIN brands ON goods.brand = brands.id
		WHERE category = ? AND goods.active = 1
		`, categoryId)
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		brandItem := models.Brands{}
		if err = rows.Scan(&brandItem.Id, &brandItem.BrandsName); err != nil {
			log.Print(err)
			return
		}
		listBrands = append(listBrands, brandItem)
	}
	return
}

func (s *Sqlite) GetGoodsByBrand(brandId, categoryId string) (goods map[models.Models][]models.Item, err error) {

	listModel := s.getListModelsByBrand(brandId, categoryId)

	goods = make(map[models.Models][]models.Item)

	for _, i := range listModel {
		rows, _ := s.db.Query(`
			SELECT goods.goodsName, goods.uah, goods.amount, brands.brandsName, models.modelsName
			FROM goods
			INNER JOIN brands ON goods.brand = brands.id
			INNER JOIN models ON goods.model = models.id
			WHERE goods.model = ? AND goods.active = 1
			`, i.Id)

		products := []models.Item{}
		for rows.Next() {
			product := models.Item{}
			if err = rows.Scan(&product.Name, &product.Price, &product.Amount, &product.Brand, &product.Model); err != nil {
				log.Print(err)
				return
			}
			products = append(products, product)
		}
		goods[i] = products
	}
	return
}

func (s *Sqlite) getListModelsByCategory(category string) (modelsImageList []models.Models) {
	rows, err := s.db.Query(`
		SELECT DISTINCT goods.model, models.image
		FROM goods
		INNER JOIN models ON goods.model = models.id
		WHERE category = ? AND goods.active = 1
		`, category)
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		modelitem := models.Models{}
		if err = rows.Scan(&modelitem.Id, &modelitem.Image); err != nil {
			log.Print(err)
			return
		}
		modelsImageList = append(modelsImageList, modelitem)
	}
	return
}

func (s *Sqlite) getListModelsByBrand(brand, category string) (modelsImageList []models.Models) {
	rows, err := s.db.Query(`
		SELECT DISTINCT goods.model, models.image
		FROM goods
		INNER JOIN models ON goods.model = models.id
		WHERE brand = ? AND category = ? AND goods.active = 1
		`, brand, category)
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		modelitem := models.Models{}
		if err = rows.Scan(&modelitem.Id, &modelitem.Image); err != nil {
			log.Print(err)
			return
		}
		modelsImageList = append(modelsImageList, modelitem)
	}
	return
}
