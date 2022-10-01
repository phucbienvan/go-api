package main

import (
	"go-api/module/category/transport/gincategory"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	Id int `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Description string `json:"description" gorm:"column:description;"`
	Image string `json:"image" gorm:"column:image;"`
}

type CategoryUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Description *string `json:"description" gorm:"column:description;"`
	Image *string `json:"image" gorm:"column:image;"`
}

func (Category) TableName() string {return "categories"}

func main() {
	dsn := os.Getenv("MYSQL_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		log.Println(db)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})

	v1 := r.Group("v1")
	categories := v1.Group("categories")

	categories.POST("", gincategory.CreateCatgory(db))

	categories.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})

			return
		}

		var data Category
		db.Where("id = ?", id).First(&data)

		c.JSON(http.StatusOK, gin.H{
			"data" : data,
		})
	})

	categories.GET("/", func(c *gin.Context) {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})

			return
		}

		var data []Category

		type Paging struct {
			Page int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}

		var pagingData Paging
		
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
			
			return
		}

		if pagingData.Limit <= 0 {
			pagingData.Limit = 5
		}

		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}
		db.Offset((pagingData.Page - 1) * pagingData.Limit).
			Order("id desc").Limit(pagingData.Limit).
			Find(&data)
		
		c.JSON(http.StatusOK, gin.H{
			"data" : data,
		})
	})

	categories.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})

			return
		}

		var data CategoryUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
			
			return
		}

		db.Where("id = ?", id).Updates(&data)

		c.JSON(http.StatusOK, gin.H{
			"data" : data,
		})
	})

	categories.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})

			return
		}

		var data Category
		db.Where("id = ?", id).Delete(&data)

		c.JSON(http.StatusOK, gin.H{
			"data" : "done",
		})
	})
	
	r.Run()
}
