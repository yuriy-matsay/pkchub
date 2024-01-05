package main

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"log"
	"pkhub/handler"
	"pkhub/service"
	"pkhub/sqlite"
)

type Template struct {
	templates map[string]*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "index", data)
}

func main() {
	loadEnv()

	db := sqlite.NewSqliteDB()
	srvc := service.NewService(db)
	hdl := handler.NewHandler(srvc)

	e := echo.New()

	templates := make(map[string]*template.Template)

	templates["startpage"] = template.Must(template.ParseFiles("view/index.html", "view/startpage.html"))
	templates["categoryitems"] = template.Must(template.ParseFiles("view/index.html", "view/categoryitems.html"))
	templates["branditems"] = template.Must(template.ParseFiles("view/index.html", "view/branditems.html"))

	e.Renderer = &Template{
		templates: templates,
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hdl.GetCategories)
	e.GET("/categories/:id", hdl.GetGoodsByCategory)
	e.GET("/brands/:id", hdl.GetGoodsByBrand)

	e.Logger.Fatal(e.Start(":8000"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}
	log.Print("load from .env successfully")
	return
}
