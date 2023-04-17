package controller

import (
	"OctavianoRyan25/GO-JWT/database"
	"OctavianoRyan25/GO-JWT/helper"
	"OctavianoRyan25/GO-JWT/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Create photo godoc
// @Summary      Create photo
// @Description  Post data and create photo using given data
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param        model.Photo   body      model.Photo  true  "Create photo"
// @Success      200  {object}  model.Photo
// @Router       /photo/create [post]
func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	Photo := model.Photo{}
	userID := uint(userData["id"].(float64)) 


	if contentType == appJSON {
		ctx.ShouldBindJSON(&Photo)
	} else {
		ctx.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Photo": Photo,
	})
}

// Update photo godoc
// @Summary      Update photo
// @Description  Put data and update photo using given data
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param        model.Photo   body      model.Photo  true  "Update photo"
// @Success      200  {object}  model.Photo
// @Router       /photo/update/{id} [put]
func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	Photo := model.Photo{}
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	userID := uint(userData["id"].(float64)) 


	if contentType == appJSON {
		ctx.ShouldBindJSON(&Photo)
	} else {
		ctx.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(model.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, Photo)
}

// Get all photo godoc
// @Summary      Get all photo
// @Description  Get data photo
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param        model.Photo   body      model.Photo  true  "get all photo"
// @Success      200  {object}  model.Photo
// @Router       /photo/get [get]
func GetAllPhoto(ctx *gin.Context) {
	db := database.GetDB()
	//userData := ctx.MustGet("userData").(jwt.MapClaims)

	var AllPhoto []model.Photo
	//photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	//userID := uint(userData["id"].(float64)) 

	err := db.Find(&AllPhoto).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, AllPhoto)
}

// Get photo by id godoc
// @Summary      Get photo by id photo
// @Description  Get data photo by given id
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param        model.Photo   body      model.Photo  true  "get photo by id"
// @Success      200  {object}  model.Photo
// @Router       /photo/get/{id} [get]
func GetPhotobyID(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Photo := model.Photo{}
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	userID := uint(userData["id"].(float64)) 

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.First(&Photo).Where("id = ?", photoId).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, Photo)
}

// Delete photo by id godoc
// @Summary      Delete photo by id photo
// @Description  Delete photo by given id
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param        model.Photo   body      model.Photo  true  "delete photo"
// @Success      200  {object}  model.Photo
// @Router       /photo/delete/{id} [delete]
func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Photo := model.Photo{}
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	userID := uint(userData["id"].(float64)) 

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.First(&Photo).Where("id = ?", photoId).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	db.Delete(&Photo)

	ctx.JSON(http.StatusOK, gin.H{
		"message"	: "Data has been deleted",
		"data"		: Photo,
	})
}