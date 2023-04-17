package controller

import (
	"OctavianoRyan25/GO-JWT/database"
	"OctavianoRyan25/GO-JWT/helper"
	"OctavianoRyan25/GO-JWT/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var(
	appJSON = "application/json"
)

// Register godoc
// @Summary      Register user
// @Description  Post data and register user
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        mode.User   body      model.User  true  "Register user"
// @Success      200  {object}  model.User
// @Router       /user/register [post]
func UserRegister(ctx *gin.Context)  {
	db := database.GetDB()
	contentType := helper.GetContentType(ctx)
	_, _ = db, contentType
	User := model.User{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status"	: "Account succesfully created",
		"username"	: User.Username,
		"email"		: User.Email,
		"age"		: User.Age,
	})
}

// Login godoc
// @Summary      Login user
// @Description  Post data and login user
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        model.User   body      model.User  true  "Login user"
// @Success      200  {object}  model.User
// @Router       /user/login [post]
func UserLogin(ctx *gin.Context)  {
	db := database.GetDB()
	contentType := helper.GetContentType(ctx)
	_, _ = db, contentType
	User := model.User{}
	password := ""

	if contentType == appJSON {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}
	
	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	comparePass := helper.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	} 
	token := helper.GenerateToken(User.ID, User.Email)

	ctx.JSON(http.StatusCreated, gin.H{
		"status"	: "Login succesfully",
		"ID"		: User.ID,
		"token"		: token,
	})
}