package gincategory

import (
	categorymodel "go-api/module/category/model"
	categorystorage "go-api/module/category/storage"
	categorybusiness "go-api/module/category/business"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCatgory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.CategoryCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		// db.Create(&data)
		store := categorystorage.NewSqlStorage(db)
		business := categorybusiness.NewCreateCategoryBusiness(store)

		if err := business.CreateCatgory(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
