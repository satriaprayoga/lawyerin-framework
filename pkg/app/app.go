package app

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/satriaprayoga/lawyerin-framework/pkg/config"
	db "github.com/satriaprayoga/lawyerin-framework/pkg/database"
	"github.com/satriaprayoga/lawyerin-framework/pkg/route"
	"gorm.io/gorm"
)

type Lawyerin struct {
	Driver   db.Driver
	RootPath string
	Debug    bool
	Route    route.Route
}

func InitLawyerin(rootPath string) *Lawyerin {

	var (
		dbCon *gorm.DB
		err   error
		debug bool
	)

	err = config.CheckDotEnv(rootPath)
	if err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

	if os.Getenv("DB_TYPE") != "" {
		dbCon, err = db.OpenDB(os.Getenv("DB_TYPE"), db.BuildConnDB())
		if err != nil {
			log.Fatalf("%v", err)
			os.Exit(1)
		}
	}

	debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	dr := db.Driver{
		DbType: os.Getenv("DB_TYPE"),
		DB:     dbCon,
	}

	route := route.New(os.Getenv("APP_PORT"), debug)

	app := &Lawyerin{
		RootPath: rootPath,
		Driver:   dr,
		Debug:    debug,
		Route:    route,
	}

	return app

}
