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

func (s *Sqlite) GetCategories() (listCategories []models.Category, err error) {
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
		category := models.Category{}
		if err = rows.Scan(&category.Id, &category.CategoryName); err != nil {
			log.Print(err)
			return
		}
		listCategories = append(listCategories, category)
	}
	return
}

func (s *Sqlite) GetBrandsByCategory(categoryId string) (listBrands []models.Brand, err error) {
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
		brand := models.Brand{}
		if err = rows.Scan(&brand.Id, &brand.BrandName); err != nil {
			log.Print(err)
			return
		}
		listBrands = append(listBrands, brand)
	}
	return
}

func (s *Sqlite) GetGoodsByCategory(categoryId string) (goods []map[models.Model][]models.Item, err error) {

	listModel := s.getListModelsByCategory(categoryId)

	stmtUAH, _ := s.db.Prepare(`
			SELECT goods.goodsName, goods.uah, goods.amount, brands.brandsName, models.modelsName
			FROM goods
			INNER JOIN brands ON goods.brand = brands.id
			INNER JOIN models ON goods.model = models.id
			WHERE goods.model = ? AND goods.active = 1 AND goods.amount > 0
			`)
	stmtUSD, _ := s.db.Prepare(`
			SELECT goods.goodsName, goods.usd, goods.amount, brands.brandsName, models.modelsName
			FROM goods
			INNER JOIN brands ON goods.brand = brands.id
			INNER JOIN models ON goods.model = models.id
			WHERE goods.model = ? AND goods.active = 1 AND goods.amount > 0
			`)

	var stmt *sql.Stmt

	if categoryId == "6" {
		stmt = stmtUSD
	} else {
		stmt = stmtUAH
	}

	for _, i := range listModel {
		// rows, _ := s.db.Query(`
		// 	SELECT goods.goodsName, goods.uah, goods.amount, brands.brandsName, models.modelsName
		// 	FROM goods
		// 	INNER JOIN brands ON goods.brand = brands.id
		// 	INNER JOIN models ON goods.model = models.id
		// 	WHERE goods.model = ? AND goods.active = 1 AND goods.amount > 0
		// 	`, i.Id)
		rows, _ := stmt.Query(i.Id)

		items := []models.Item{}
		itemListByModel := make(map[models.Model][]models.Item)
		for rows.Next() {
			item := models.Item{}
			if err = rows.Scan(&item.Name, &item.Price, &item.Amount, &item.Brand, &item.Model); err != nil {
				log.Print(err)
				return
			}
			items = append(items, item)
		}
		if len(items) > 0 {
			itemListByModel[i] = items
			goods = append(goods, itemListByModel)
		}
	}
	if len(goods) == 0 {
		goods, err = s.getAllGoodsByCategory(categoryId)
	}
	return
}

func (s *Sqlite) GetGoodsByBrand(brandId, categoryId string) (goods map[models.Model][]models.Item, err error) {

	listModel := s.getListModelsByBrand(brandId, categoryId)

	goods = make(map[models.Model][]models.Item)

	for _, i := range listModel {
		rows, _ := s.db.Query(`
			SELECT goods.goodsName, goods.uah, goods.amount, brands.brandsName, models.modelsName
			FROM goods
			INNER JOIN brands ON goods.brand = brands.id
			INNER JOIN models ON goods.model = models.id
			WHERE goods.model = ? AND goods.active = 1
			`, i.Id)

		items := []models.Item{}
		for rows.Next() {
			item := models.Item{}
			if err = rows.Scan(&item.Name, &item.Price, &item.Amount, &item.Brand, &item.Model); err != nil {
				log.Print(err)
				return
			}
			items = append(items, item)
		}
		goods[i] = items
	}
	return
}

func (s *Sqlite) getListModelsByCategory(category string) (modelsImageList []models.Model) {
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
		modelitem := models.Model{}
		if err = rows.Scan(&modelitem.Id, &modelitem.Image); err != nil {
			log.Print(err)
			return
		}
		modelsImageList = append(modelsImageList, modelitem)
	}
	return
}

func (s *Sqlite) getListModelsByBrand(brand, category string) (modelsImageList []models.Model) {
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
		modelitem := models.Model{}
		if err = rows.Scan(&modelitem.Id, &modelitem.Image); err != nil {
			log.Print(err)
			return
		}
		modelsImageList = append(modelsImageList, modelitem)
	}
	return
}

func (s *Sqlite) getAllGoodsByCategory(categoryId string) (goods []map[models.Model][]models.Item, err error) {

	listModel := s.getListModelsByCategory(categoryId)

	for _, i := range listModel {
		rows, _ := s.db.Query(`
			SELECT goods.goodsName, goods.uah, goods.amount, brands.brandsName, models.modelsName
			FROM goods
			INNER JOIN brands ON goods.brand = brands.id
			INNER JOIN models ON goods.model = models.id
			WHERE goods.model = ? AND goods.active = 1
			`, i.Id)

		items := []models.Item{}
		itemListByModel := make(map[models.Model][]models.Item)
		for rows.Next() {
			item := models.Item{}
			if err = rows.Scan(&item.Name, &item.Price, &item.Amount, &item.Brand, &item.Model); err != nil {
				log.Print(err)
				return
			}
			items = append(items, item)
		}
		itemListByModel[i] = items
		goods = append(goods, itemListByModel)
	}
	return
}
