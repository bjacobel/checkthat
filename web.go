package main

import (
    "github.com/laurent22/ripple"
    "github.com/bjacobel/checkthat/controllers"
    "net/http"
    "os"
)

func main() {
    // Build the REST application

    app := ripple.NewApplication()

    userController := controllers.NewUserController()
    app.RegisterController("users", userController)

    app.AddRoute(ripple.Route{Pattern: ":_controller/:id/:_action"})
    app.AddRoute(ripple.Route{Pattern: ":_controller/:id/"})
    app.AddRoute(ripple.Route{Pattern: ":_controller"})

    // Start the server

    err := http.ListenAndServe(":"+os.Getenv("PORT"), app)
    if err != nil {
        panic(err)
    }
}