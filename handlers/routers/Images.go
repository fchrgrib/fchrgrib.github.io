package routers

import (
	"github.com/controllers/wallpaperpage"
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/libs/middleware"
	"github.com/models"
)

func Images(routers *gin.Engine) {
	var wallpaper []models.WallpaperCollectionDB

	db, err := database.ConnectDB()
	if err != nil {
		return
	}

	db.Table("wallpaper_collect").Find(&wallpaper)

	rImage := routers.Group("/images")
	rImage.Use(middleware.AuthWithToken)

	if len(wallpaper) != 0 {
		for _, values := range wallpaper {
			rImage.Static(values.ImageId, values.Path)
		}
	}

	rImage.DELETE("/:id/delete", wallpaperpage.DeleteWallpaperController)
}
