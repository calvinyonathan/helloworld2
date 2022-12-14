package main

import (
	"github.com/calvin/helloworld2/api"
	_ "github.com/joho/godotenv/autoload"
)

//	func handler(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"message": "hello world 2",
//		})
//	}
// func (r *repository) handler(c *gin.Context) {
// 	var todos []Todos
// 	res := r.db.Find(&todos)
// 	if res.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": todos,
// 	})
// }

// func (r *repository) postHandler(c *gin.Context) {
// 	var data DataRequest
// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	todo := Todos{
// 		Task: data.Task,
// 		Done: false,
// 	}

// 	res := r.db.Create(&todo)
// 	if res.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "data berhasil terkirim", "data": todo})
// }

// type Todos struct {
// 	gorm.Model
// 	Task string `json:"task"`
// 	Done bool   `json:"done"`
// }
// type repository struct {
// 	db *gorm.DB
// }
// type DataRequest struct {
// 	Task string `json:"task" binding:"required"`
// }

func main() {

	// var db *gorm.DB
	// var err error
	// dbUrl := os.Getenv("DATABASE_URL")

	// if os.Getenv("ENVIRONMENT") == "PROD" {
	// 	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	// } else {
	// 	db, err = gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	// }

	// sqlDB, err := db.DB()
	// if err != nil {
	// 	panic("failed to get database")
	// }

	// if err := sqlDB.Ping(); err != nil {
	// 	panic("failed to ping database")
	// }
	// if err := db.AutoMigrate(&Todos{}); err != nil {
	// 	panic("failed to migrate database")
	// }
	// repo := repository{
	// 	db: db,
	// }
	// r := gin.Default()
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"},
	// }))

	// r.GET("/", repo.handler)
	// r.POST("/send", repo.postHandler)
	// r.Run(":" + os.Getenv("PORT"))
	db, err := api.SetupDb()
	if err != nil {
		panic(err)
	}

	server := api.MakeServer(db)
	server.RunServer()
}
