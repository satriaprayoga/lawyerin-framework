package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/satriaprayoga/lawyerin-framework/controller"
	"github.com/satriaprayoga/lawyerin-framework/data"
	"github.com/satriaprayoga/lawyerin-framework/pkg/app"
	db "github.com/satriaprayoga/lawyerin-framework/pkg/database"
)

type application struct {
	L *app.Lawyerin
	S *data.Store
	H *controller.Handler
}

func (a *application) Start() {
	serv := &http.Server{
		Addr:         fmt.Sprintf(":%s", a.L.Route.Port),
		Handler:      a.L.Route.R,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	defer db.Close(a.L.Driver.DB)

	fmt.Printf("Listening on port %s\n", a.L.Route.Port)
	err := serv.ListenAndServe()
	log.Fatal(err)
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	l := app.InitLawyerin(path)
	s := data.New(l.Driver.DB)
	h := controller.New(l.Route.R)
	app := &application{
		L: l,
		S: s,
		H: h,
	}
	app.Start()

}
