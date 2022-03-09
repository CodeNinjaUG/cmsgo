package main

import (
	"fmt"
	"github.com/codeninjaug/crm/database"
    "github.com/codeninjaug/crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_"gorm.io/driver/mysql"
)

func setUpRoutes(app *fiber.App){
	app.Get("api/v1/leads",lead.GetLeads)
	app.Get("api/v1/lead/:id",lead.GetLead)
	app.Post("api/v1/lead/",lead.NewLead)
	app.Delete("api/v1/lead/:id",lead.DeleteLead)
}
func initDatabase(){
  var err error
  database.DBConn , err = gorm.Open("mysql","root@tcp(localhost)/leads?charset=utf8&parseTime=True&loc=Local")
  if err!=nil {
	  panic("failed to connect to database")
  }
  fmt.Println("connection opened to database")
  database.DBConn.AutoMigrate(&lead.Lead{})
  fmt.Println("connected to database")
}
func main()  {
  app:=fiber.New()
  initDatabase()
  //defer database.DBConn.Close()
  setUpRoutes(app)
  app.Listen(5200)
  defer database.DBConn.Close()
}