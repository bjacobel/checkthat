package main

import (
	"github.com/jinzhu/gorm"
	"github.com/laurent22/ripple"
	_ "github.com/lib/pq"
	//"github.com/bjacobel/checkthat/controllers"
	"fmt"
	"github.com/bjacobel/checkthat/models"
	"net/http"
	"os"
)

func main() {
	// set up the database
	db, dberr := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@ec2-54-197-241-67.compute-1.amazonaws.com:5432/%s", os.Getenv("PGUSER"), os.Getenv("PGPW"), os.Getenv("PGDB")))
	if dberr != nil {
		panic(dberr)
	}
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Device{})

	fmt.Println(db.DB())

	// Build the REST application
	app := ripple.NewApplication()
	// // userController := controllers.NewUserController()
	// app.RegisterController("users", userController)

	// app.AddRoute(ripple.Route{Pattern: ":_controller/:id/:_action"})
	// app.AddRoute(ripple.Route{Pattern: ":_controller/:id/"})
	// app.AddRoute(ripple.Route{Pattern: ":_controller"})

	// Start the server
	httperr := http.ListenAndServe(":"+os.Getenv("PORT"), app)
	if httperr != nil {
		panic(httperr)
	}
}
