package controllers

import (
	"DTSGolang/FinalProject/database"
	"DTSGolang/FinalProject/helpers"
	"DTSGolang/FinalProject/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreatePhoto godoc
// @Summary Post a new photo
// @Description post details of a new photo based on current user
// @Tags photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param models.Photo body models.Photo true "create a photo"
// @Succes 201 {object} models.Photo
// @Router /photos [post]
func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	User := models.User{}
	errA := db.Debug().First(&User, "id = ?", userID).Error
	if errA != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "User not Found",
			"message": errA.Error(),
		})

		return
	}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Photo)
	} else {
		ctx.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": Photo,
	})
}

// GetPhoto godoc
// @Summary Get photo details for the given id
// @Description Get details of a photo corresponding to the input id
// @Tags photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the photo"
// @Succes 200 {object} models.Photo
// @Router /photos/{ID} [get]
func GetPhoto(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	Photo := models.Photo{}

	photoID, _ := strconv.Atoi(ctx.Param("photoID"))

	err := db.Debug().First(&Photo, "id = ?", uint(photoID)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Photo Not Found",
				"message": err.Error(),
			})
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Photo,
	})
}

// GetMyPhotos godoc
// @Summary Get all photo corresponding user
// @Description Get all photo data corresponding user
// @Tags photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Succes 200 {object} models.Photo
// @Router /photos/my [get]
func GetMyPhotos(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Photos := []models.Photo{}

	userID := uint(userData["id"].(float64))

	err := db.Where("user_id", userID).Find(&Photos).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Photos,
	})
}

// GetPhotos godoc
// @Summary Get all photo
// @Description Get all photo data
// @Tags photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Succes 200 {object} models.Photo
// @Router /photos [get]
func GetPhotos(ctx *gin.Context) {
	db := database.GetDB()

	Photos := []models.Photo{}

	err := db.Debug().Find(&Photos).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Photos,
	})
}

// UpdatePhoto godoc
// @Summary Update photo for the given id
// @Description Update details of a photo corresponding to the input id
// @Tags photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the photo"
// @Succes 200 {object} models.Socialmedia
// @Router /photos/{ID} [put]
func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	Photo := models.Photo{}

	photoID, _ := strconv.Atoi(ctx.Param("photoID"))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Photo)
	} else {
		ctx.ShouldBind(&Photo)
	}

	err := db.Debug().Model(&Photo).Where("id = ?", uint(photoID)).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully update photo",
	})
}

// DeletePhoto godoc
// @Summary Delete photo details for the given id
// @Description Delete details of a photo corresponding to the input id
// @Tags photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the photo"
// @Succes 200 {object} models.Photo
// @Router /photos/{ID} [delete]
func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()

	Photo := models.Photo{}
	photoID, _ := strconv.Atoi(ctx.Param("photoID"))

	Comments := []models.Comment{}

	// Must delete all related comments first
	errComments := db.Where("photo_id = ?", photoID).Delete(&Comments).Error
	if errComments != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": errComments.Error(),
		})
	}

	err := db.Debug().Where("id = ?", uint(photoID)).Delete(&Photo).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted photo",
	})
}
