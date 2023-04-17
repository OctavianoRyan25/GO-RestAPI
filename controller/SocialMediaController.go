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

// Create social media godoc
// @Summary      Create social media
// @Description  Post data and create social media using given data
// @Tags         social media
// @Accept       json
// @Produce      json
// @Param        model.SocialMedia   body      model.SocialMedia  true  "Create social media"
// @Success      200  {object}  model.SocialMedia
// @Router       /socialmedia/create [post]
func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	SocialMedia := model.SocialMedia{}
	userID := uint(userData["id"].(float64)) 


	if contentType == appJSON {
		ctx.ShouldBindJSON(&SocialMedia)
	} else {
		ctx.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Photo": SocialMedia,
	})
}

// Update social media godoc
// @Summary      Update social media
// @Description  Put data and update social media using given data
// @Tags         social media
// @Accept       json 
// @Produce      json
// @Param        model.SocialMedia   body      model.SocialMedia  true  "Update social media"
// @Success      200  {object}  model.SocialMedia
// @Router       /socialmedia/update/{id} [put]
func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	SocialMedia := model.SocialMedia{}
	socialmediaId, _ := strconv.Atoi(ctx.Param("socialmediaId"))
	userID := uint(userData["id"].(float64)) 


	if contentType == appJSON {
		ctx.ShouldBindJSON(&SocialMedia)
	} else {
		ctx.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialmediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialmediaId).Updates(model.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, SocialMedia)
}

// Get all social media godoc
// @Summary      Get all social media
// @Description  Get data social media
// @Tags         social media
// @Accept       json
// @Produce      json
// @Param        model.SocialMedia   body      model.SocialMedia  true  "get all social media"
// @Success      200  {object}  model.SocialMedia
// @Router       /socialmedia/get [get]
func GetAllSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	//userData := ctx.MustGet("userData").(jwt.MapClaims)

	var AllSocialMedia []model.SocialMedia
	//photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	//userID := uint(userData["id"].(float64)) 

	err := db.Find(&AllSocialMedia).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, AllSocialMedia)
}

// Get social media by id godoc
// @Summary      Get social media by id social media
// @Description  Get data social media by given id
// @Tags         social media
// @Accept       json
// @Produce      json
// @Param        model.SocialMedia   body      model.SocialMedia  true  "get social media by id"
// @Success      200  {object}  model.SocialMedia
// @Router       /socialmedia/get/{id} [get]
func GetSocialmediaById(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)

	SocialMedia := model.SocialMedia{}
	socialmediaId, _ := strconv.Atoi(ctx.Param("socialmediaId"))
	userID := uint(userData["id"].(float64)) 

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialmediaId)

	err := db.First(&SocialMedia).Where("id = ?", socialmediaId).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, SocialMedia)
}

// Delete social media by id godoc
// @Summary      Delete social media by id social media
// @Description  Delete social media by given id
// @Tags         social media
// @Accept       json
// @Produce      json
// @Param        model.SocialMedia   body      model.SocialMedia  true  "delete social media"
// @Success      200  {object}  model.SocialMedia
// @Router       /socialmedia/delete/{id} [delete]
func Deletesocialmedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)

	SocialMedia := model.SocialMedia{}
	socialmediaId, _ := strconv.Atoi(ctx.Param("socialmediaId"))
	userID := uint(userData["id"].(float64)) 

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialmediaId)

	err := db.First(&SocialMedia).Where("id = ?", socialmediaId).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	db.Delete(&socialmediaId)

	ctx.JSON(http.StatusOK, gin.H{
		"message"	: "Data has been deleted",
		"data"		: socialmediaId,
	})
}