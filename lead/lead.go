package lead

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/saeedmzr/go_crm/database"
	_ "github.com/saeedmzr/go_crm/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	err := c.JSON(leads)
	if err != nil {
		return err
	}
	return nil

}

func GetLead(c *fiber.Ctx) error {
	id := c.Params(":id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	err := c.JSON(lead)
	if err != nil {
		return err
	}
	return nil
}
func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503)
		return err
	}
	db.Create(&lead)
	err := c.JSON(lead)
	if err != nil {
		return err
	}
	return nil

}
func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		err := c.Status(404).Send([]byte("no lead found"))
		if err != nil {
			return err
		}
		return nil
	}
	db.Delete(id)
	err := c.Send([]byte("test"))
	if err != nil {
		return err

	}
	return nil
}
