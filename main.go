package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 初始化模板引擎
	engine := html.New("./views", ".html")
	// 初始化数据库
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		panic(err)
	}
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./static")
	app.Get("/index", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello,World!",
		}, "layouts/main")
	})
	app.Listen(":3000")
}
