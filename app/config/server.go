package config

import (
	"GoAppOne/app/database/migration"
	"GoAppOne/app/database/seeder"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
	DB     *gorm.DB
	APP_NAME,
	APP_PORT,
	DB_USER,
	DB_PASSWORD,
	DB_NAME string
}

func getEnv(env_key string, falback string) (result string) {
	result = os.Getenv(env_key)
	if result == "" {
		return falback
	}
	return result
}
func (server *Server) initialize() {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
	enverr := godotenv.Load()
	if enverr != nil {
		panic(enverr)
	}
	server.APP_NAME = getEnv("APP_NAME", "GoAppOne")
	server.APP_PORT = getEnv("APP_PORT", "localhost:8080")
	server.DB_NAME = getEnv("DB_NAME", "media_visual_banyuwangi")
	server.DB_USER = getEnv("DB_USER", "root")
	server.DB_PASSWORD = getEnv("DB_PASSWORD", "")
}
func (server *Server) dbLaunch() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", server.DB_USER, server.DB_PASSWORD, server.DB_NAME)
	var err error
	server.DB, err = gorm.Open(mysql.New(mysql.Config{
		DisableWithReturning: true,
		DSN:                  dsn, 
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
}
func (server *Server) serverLaunch() {
	var conf = http.Server{
		Addr:    server.APP_PORT,
		Handler: server.Router,
	}
	fmt.Println("Welcome to", server.APP_NAME)
	fmt.Println("Server launching in", server.APP_PORT)
	err := conf.ListenAndServe()
	panic(err)
}

func (server *Server) DbMigration() {
	for _, model := range migration.RegisterModels() {
		server.DB.AutoMigrate(model.Model)
	}
	seeder.RunSeeder(server.DB)
}

func Run(migrate bool) {
	var server Server = Server{}
	server.initialize()
	server.dbLaunch()
	if migrate {
		server.DbMigration()
	}
	server.serverLaunch()
}
