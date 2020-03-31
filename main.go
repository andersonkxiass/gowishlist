package main

import (
	"WishlistApi/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	_ "os"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("DB_URL"))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.List{})
	db.AutoMigrate(&models.Wish{})

	return db
}

func main() {

	db := initDB()
	defer db.Close()

	wishAPI := initWishAPI(db)
	listAPI := initListAPI(db)

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/wishes", wishAPI.FindAll)
	r.GET("/list", listAPI.ListALl)

	err := r.Run()

	if err != nil {
		panic(err)
	}
}
