package main
import(
	"fmt"
	"github.com/gofiber/fiber"//this import statement+"go mod tidy" will install the package
	"github.com/1234arushi/go-fiber-crm-basic/database"
	"github.com/1234arushi/go-fiber-crm-basic/lead"
	"github.com/jinzhu/gorm"
)
func setupRoutes(app *fiber.App){//we want to call app by reference hence,' * '
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("api/v1/lead/:id",lead.GetLead)
	app.Post("api/v1/lead",lead.NewLead)
	app.Delete("api/v1/lead/:id",lead.DeleteLead)
}
func initDatabase(){
	var err error
	database.DBConn,err=gorm.Open("sqlite3","leads.db")
	if err!=nil{
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Datbase Migrated")

}
func main(){
	//what is the "app" that is being called in setupRoutes function
	//ans-> app is basically an instance of fiber
	app:=fiber.New()
    setupRoutes(app)
	initDatabase()//start the db connection
	app.Listen(3000)//start of your server,port = 3000
    defer database.DBConn.Close()//defer->last line to be run


}