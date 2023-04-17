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

// Create comment godoc
// @Summary      Create comment
// @Description  Post data and create comment using given data
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param        model.Comment   body      model.Comment  true  "Create comment"
// @Success      200  {object}  model.Comment
// @Router       /comment/create [post]
func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	Comment := model.Comment{}
	userID := uint(userData["id"].(float64)) 


	if contentType == appJSON {
		ctx.ShouldBindJSON(&Comment)
	} else {
		ctx.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Photo": Comment,
	})
}

// Update comment godoc
// @Summary      Update comment
// @Description  Put data and update comment using given data
// @Tags         comment
// @Accept       json 
// @Produce      json
// @Param        model.Comment   body      model.Comment  true  "Update comment"
// @Success      200  {object}  model.Comment
// @Router       /comment/update/{id} [put]
func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)

	Comment := model.Comment{}
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))
	userID := uint(userData["id"].(float64)) 


	if contentType == appJSON {
		ctx.ShouldBindJSON(&Comment)
	} else {
		ctx.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(model.Comment{Message: Comment.Message}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, Comment)
}

// Get all comment godoc
// @Summary      Get all comment
// @Description  Get data comment
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param        model.Comment   body      model.Comment  true  "get all comment"
// @Success      200  {object}  model.Comment
// @Router       /comment/get [get]
func GetAllComment(ctx *gin.Context) {
	db := database.GetDB()
	//userData := ctx.MustGet("userData").(jwt.MapClaims)

	var AllComment []model.Comment
	//photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	//userID := uint(userData["id"].(float64)) 

	err := db.Find(&AllComment).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, AllComment)
}

// Get comment by id godoc
// @Summary      Get comment by id comment
// @Description  Get data comment by given id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param        model.Comment   body      model.Comment  true  "get comment by id"
// @Success      200  {object}  model.Comment
// @Router       /comment/get/{id} [get]
func GetCommentbyID(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Comment := model.Comment{}
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))
	userID := uint(userData["id"].(float64)) 

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.First(&Comment).Where("id = ?", commentId).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, Comment)
}

// Delete comment by id godoc
// @Summary      Delete comment by id comment
// @Description  Delete comment by given id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param        model.Comment   body      model.Comment  true  "delete comment"
// @Success      200  {object}  model.Comment
// @Router       /comment/delete/{id} [delete]
func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Comment := model.Comment{}
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))
	userID := uint(userData["id"].(float64)) 

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.First(&Comment).Where("id = ?", commentId).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err" 		: "Bad Request",
			"message"	: err.Error(),
		})
		return
	}

	db.Delete(&Comment)

	ctx.JSON(http.StatusOK, gin.H{
		"message"	: "Data has been deleted",
		"data"		: Comment,
	})
}