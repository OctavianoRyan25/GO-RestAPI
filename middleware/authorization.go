package middleware

import (
	"OctavianoRyan25/GO-JWT/database"
	"OctavianoRyan25/GO-JWT/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		photoID, err := strconv.Atoi(ctx.Param("photoId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error"   : "bad parameter",
				"message" : "invalid parameter",
			})
			return
		}
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := model.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoID)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error"		: "Data Not Found",
				"message"	: "data dosen't exist",
			})
			return
		}

		if Photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error"		: "Unauthorized",
				"message"	: "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func AuthorizationComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		commentID, err := strconv.Atoi(ctx.Param("commentId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error"   : "bad parameter",
				"message" : "invalid parameter",
			})
			return
		}
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := model.Comment{}

		err = db.Select("user_id").First(&Comment, uint(commentID)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error"		: "Data Not Found",
				"message"	: "data dosen't exist",
			})
			return
		}

		if Comment.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error"		: "Unauthorized",
				"message"	: "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func AuthorizationSocialMedia() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		socialmediaID, err := strconv.Atoi(ctx.Param("socialmediaId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error"   : "bad parameter",
				"message" : "invalid parameter",
			})
			return
		}
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := model.SocialMedia{}

		err = db.Select("user_id").First(&SocialMedia, uint(socialmediaID)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error"		: "Data Not Found",
				"message"	: "data dosen't exist",
			})
			return
		}

		if SocialMedia.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error"		: "Unauthorized",
				"message"	: "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}
