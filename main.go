package main

import (
	"time"

	schemas "diab/database"
	routes "diab/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var defaultHTTPPort = ":8080"
var defaultDbName = "localDatabase.db"

func main() {
	// Opening SQLite database before moving to server hosting
	db, err := gorm.Open(sqlite.Open(defaultDbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Initialising automigration
	schemas.InitAutoMigrate(db)
	// Configuring connection pool
	sqlDB, err := db.DB()
	// Setting max of IDLEing connections
	sqlDB.SetMaxIdleConns(10)
	// Limiting max open connections
	sqlDB.SetMaxOpenConns(100)
	// Limiting lifetime of each connection
	sqlDB.SetConnMaxLifetime(time.Second * 180)
	// Creating default router
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	// C Routing
	router.POST("/user", func(c *gin.Context) { routes.NewUser(c, db) })
	router.POST("/chat", func(c *gin.Context) { routes.NewChat(c, db) })
	router.POST("/message", func(c *gin.Context) { routes.NewMessage(c, db) })
	router.POST("/test", func(c *gin.Context) { routes.NewTest(c, db) })
	// R Routing
	router.GET("/chats", func(c *gin.Context) { routes.GetChats(c, db) })
	router.GET("/history", func(c *gin.Context) { routes.GetHistory(c, db) })
	router.GET("/tests", func(c *gin.Context) { routes.GetTests(c, db) })
	router.GET("/test", func(c *gin.Context) { routes.GetTest(c, db) })
	// U Routing
	router.POST("/users", func(c *gin.Context) { routes.EditUser(c, db) })
	// Start listening
	router.Run(defaultHTTPPort)
}
