package lead

import (
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/1234arushi/go-fiber-crm-basic/database"
)
type Lead struct{ 
	gorm.Model
	Name string `json:"name"`
	Company string `json:"company"`
	Email string `json:"email"`
	Phone int `json:"phone"`

}
func GetLeads(c *fiber.Ctx){
	db:=database.DBConn
	var leads []Lead//slice of ->almost like an array of objects
    db.Find(&leads)
	c.JSON(leads)

}
func GetLead(c *fiber.Ctx){
	id:=c.Params("id")
	db:=database.DBConn
	var lead Lead
	db.Find(&lead,id)
	c.JSON(lead)//sending response

}
func NewLead(c *fiber.Ctx){
	db:=database.DBConn
	lead:=new(Lead)
	if err:=c.BodyParser(lead);err!=nil{
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
	

}
func DeleteLead(c *fiber.Ctx){
	id:=c.Params("id")
	db:=database.DBConn
	var lead Lead
	db.First(&lead,id)
	if lead.Name==""{
		c.Status(500).Send("No lead found with ID")
		return 
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted")

}