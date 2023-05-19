package controllers

import (
	"DTSGolang/FinalProject/database"
	"DTSGolang/FinalProject/helpers"
	"DTSGolang/FinalProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

type registerInput struct {
	Username string `gorm:"not null;UniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null;UniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Password hasto have a minimun length 6 characters"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required, numeric~Age must be numeric"`
}

type loginInput struct {
	Email    string `gorm:"not null;UniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Password hasto have a minimun length 6 characters"`
}

// UserRegister godoc
// @Summary Register a new user
// @Description Register a new user using email, username, and password
// @Tags user
// @Accept json
// @Produce json
// @Param Input body registerInput true "register a user"
// @Success 201 {object} models.User
// @Router /users/register [post]
func UserRegister(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType

	User := models.User{}
	registerInput := registerInput{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&registerInput)
	} else {
		ctx.ShouldBind(&registerInput)
	}

	User.Email = registerInput.Email
	User.Password = registerInput.Password
	User.Username = registerInput.Username
	User.Age = registerInput.Age

	err := db.Debug().Create(&User).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"username": User.Username,
		"email":    User.Email,
		"password": User.Password,
		"age":      User.Age,
	})
}

// UserLogin godoc
// @Summary Login an existing user
// @Description Register an existing user using email, and password
// @Tags user
// @Accept json
// @Produce json
// @Param Input body loginInput true "login an user"
// @Success 200 {object} models.User
// @Router /users/login [post]
func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	User := models.User{}
	loginInput := loginInput{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&loginInput)
	} else {
		ctx.ShouldBind(&loginInput)
	}

	err := db.Debug().Where("email = ?", loginInput.Email).Take(&User).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized -- Email",
			"message": "invalid email/password",
		})

		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(loginInput.Password))

	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized -- Pass",
			"message": "invalid email/password",
		})

		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
