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

func CreateComment(c *gin.Context) {

	db := database.StartDB()

	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)

}

func GetOneComment(c *gin.Context) {

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	comment := models.Comment{}

	db := database.StartDB()

	err := db.First(&comment, "id = ?", commentId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Photo Not Found")
			return
		}
		print("Error Find Comment : ", err)
	}

	c.JSON(http.StatusOK, comment)

}

func GetComments(c *gin.Context) {
	var comments []models.Comment
	db := database.StartDB()
	err := db.Model(&models.Comment{}).Find(&comments).Error

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, comments)
}

func DeleteComment(ctx *gin.Context) {
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))
	comment := models.Comment{}
	db := database.StartDB()

	err := db.Where("id = ?", commentId).Delete(&comment).Error

	if err != nil {
		fmt.Println("Error Deleted comment : ", err.Error())
	}

	ctx.JSON(http.StatusOK, "Deleted")
}

func UpdateComment(c *gin.Context) {
	db := database.StartDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{
		PhotoID: Comment.PhotoID,
		Message: Comment.Message,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)

}
