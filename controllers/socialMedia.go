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

func CreateSocialMedia(c *gin.Context) {

	db := database.StartDB()

	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)

}

func GetOneSocialMedia(c *gin.Context) {

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	socialMedia := models.SocialMedia{}

	db := database.StartDB()

	err := db.First(&socialMedia, "id = ?", socialMediaId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Social Media Not Found")
			return
		}
		print("Error Find Social Media : ", err)
	}

	c.JSON(http.StatusOK, socialMedia)

}

func GetSocialMedias(c *gin.Context) {
	var socialMedias []models.SocialMedia
	db := database.StartDB()
	err := db.Model(&models.SocialMedia{}).Find(&socialMedias).Error

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, socialMedias)
}

func DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))
	socialMedia := models.SocialMedia{}
	db := database.StartDB()

	err := db.Where("id = ?", socialMediaId).Delete(&socialMedia).Error

	if err != nil {
		fmt.Println("Error Deleted Social Media : ", err.Error())
	}

	ctx.JSON(http.StatusOK, "Deleted")
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.StartDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{
		Name:           SocialMedia.Name,
		SocialMediaUrl: SocialMedia.SocialMediaUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)

}
