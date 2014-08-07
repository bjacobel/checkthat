package main

import (
	"fmt"
	"github.com/bjacobel/checkthat/controllers"
	"github.com/bjacobel/checkthat/models"
	"github.com/jinzhu/gorm"
	"github.com/laurent22/ripple"
	"net/http"
	"os"
)

var db gorm.DB

func main() {
	// Set up the DB
	db, dberr := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@ec2-54-197-241-67.compute-1.amazonaws.com:5432/%s", os.Getenv("PGUSER"), os.Getenv("PGPW"), os.Getenv("PGDB")))

	if dberr != nil {
		panic(dberr)
	}

	db.AutoMigrate(models.Device{})
	db.AutoMigrate(models.User{})

	// Build the REST application
	app := ripple.NewApplication()

	deviceController := controllers.NewDeviceController(db)
	app.RegisterController("devices", deviceController)

	userController := controllers.NewUserController(db)
	app.RegisterController("users", userController)

	app.AddRoute(ripple.Route{Pattern: ":_controller/:_action"})
	app.AddRoute(ripple.Route{Pattern: ":_controller/:id/"})
	app.AddRoute(ripple.Route{Pattern: ":_controller"})

	// Start the server
	httperr := http.ListenAndServe(":"+os.Getenv("PORT"), app)
	if httperr != nil {
		panic(httperr)
	}
}
