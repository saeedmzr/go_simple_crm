package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/saeedmzr/go_crm/database"
	"github.com/saeedmzr/go_crm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead/", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead/", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("fails to connect to db.")
	}
	fmt.Println("connection was successful to db")
	database.DBConn.AutoMigrate(&lead.Lead{})
}

func main() {

	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	err := app.Listen("3000")
	if err != nil {
		return
	}
	defer func(DBConn *gorm.DB) {
		err := DBConn.Close()
		if err != nil {

		}
	}(database.DBConn)
}
