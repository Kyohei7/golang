package controllers

import (
	"errors"
	"fmt"
	"latihan-simple-rest-api-middleware/database"
	"latihan-simple-rest-api-middleware/helpers"
	"latihan-simple-rest-api-middleware/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePhoto(c *gin.Context) {

	db := database.StartDB()

	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)

}

func GetOnePhoto(c *gin.Context) {

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	photo := models.Photo{}

	db := database.StartDB()

	err := db.First(&photo, "id = ?", photoId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Photo Not Found")
			return
		}
		print("Error Find Photo : ", err)
	}

	c.JSON(http.StatusOK, photo)

}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	db := database.StartDB()
	err := db.Model(&models.Photo{}).Find(&photos).Error

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, photos)
}

func DeletePhoto(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	photo := models.Photo{}
	db := database.StartDB()

	err := db.Where("id = ?", photoId).Delete(&photo).Error

	if err != nil {
		fmt.Println("Error Deleted photo : ", err.Error())
	}

	ctx.JSON(http.StatusOK, "Deleted")
}

func UpdatePhoto(c *gin.Context) {
	db := database.StartDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{
		Title:    Photo.Title,
		Caption:  Photo.Caption,
		PhotoUrl: Photo.PhotoUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)

}
