package main

import (
	"fmt"
	"github.com/bjacobel/checkthat/controllers"
	"github.com/bjacobel/checkthat/models"
	"github.com/bjacobel/twilio-go"
	"github.com/jinzhu/gorm"
	"github.com/laurent22/ripple"
	"net/http"
	"os"
	"strings"
)

var db gorm.DB
var chttp = http.NewServeMux()

func main() {
	// Set up the DB
	db, dberr := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@ec2-54-197-241-67.compute-1.amazonaws.com:5432/%s", os.Getenv("PGUSER"), os.Getenv("PGPW"), os.Getenv("PGDB")))

	if dberr != nil {
		panic(dberr)
	}

	db.AutoMigrate(models.Device{})
	db.AutoMigrate(models.User{})

	// Set up the Twilio client
	twclient := twilio.NewTwilioRestClient(os.Getenv("TWILIO_SID"), os.Getenv("TWILIO_TOKEN"))

	// Build the REST application
	app := ripple.NewApplication()

	deviceController := controllers.NewDeviceController(db)
	app.RegisterController("devices", deviceController)

	userController := controllers.NewUserController(db, twclient)
	app.RegisterController("users", userController)

	app.AddRoute(ripple.Route{Pattern: ":_controller/:_action"})
	app.AddRoute(ripple.Route{Pattern: ":_controller/:id/"})
	app.AddRoute(ripple.Route{Pattern: ":_controller"})

	// serve the go app on /api/v1
	app.SetBaseUrl("/api/v1/")
	http.HandleFunc("/api/v1", app.ServeHTTP)

	//serve the js app on /
	http.HandleFunc("/", HomeHandler)

	httperr := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if httperr != nil {
		panic(httperr)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".") {
		chttp.ServeHTTP(w, r)
	} else {
		fmt.Fprintf(w, "HomeHandler")
	}
}
