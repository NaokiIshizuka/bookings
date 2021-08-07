package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	config2 "github.com/tsawler/bookings-app/internal/config"
	handlers2 "github.com/tsawler/bookings-app/internal/handlers"
	render2 "github.com/tsawler/bookings-app/internal/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config2.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {
	// change this to true when in production
	app.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render2.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers2.NewRepo(&app)
	handlers2.NewHandlers(repo)

	render2.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
