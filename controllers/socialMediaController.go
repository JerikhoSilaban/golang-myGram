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

// CreateSocialMedia godoc
// @Summary Post a new social media
// @Description post details of a new social media based on current user
// @Tags social media
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param models.SocialMedia body models.SocialMedia true "create a social media"
// @Succes 201 {object} models.SocialMedia
// @Router /socialMedia [post]
func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&SocialMedia)
	} else {
		ctx.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = uint(userData["id"].(float64))

	err := db.Debug().Create(&SocialMedia).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": SocialMedia,
	})
}

// GetSocialMedia godoc
// @Summary Get social media details for the given id
// @Description Get details of a social media corresponding to the input id
// @Tags social media
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the social media"
// @Succes 200 {object} models.SocialMedia
// @Router /socialMedia/{ID} [get]
func GetSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType

	SocialMedia := models.SocialMedia{}
	socialMediaID, _ := strconv.Atoi(ctx.Param("socialMediaID"))

	err := db.Debug().Find(&SocialMedia, "id = ?", uint(socialMediaID)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Social Media Not Found",
				"message": err.Error(),
			})

			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": SocialMedia,
	})
}

// GetMySocialMedias godoc
// @Summary Get all social media corresponding user
// @Description Get all social media data corresponding user
// @Tags social media
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Succes 200 {object} models.SocialMedia
// @Router /socialMedia/my [get]
func GetMySocialMedias(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)

	SocialMedias := []models.SocialMedia{}

	userID := uint(userData["id"].(float64))

	err := db.Where("user_id = ?", userID).Find(&SocialMedias).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": SocialMedias,
	})

}

// GetSocialMedias godoc
// @Summary Get all social media
// @Description Get all social media data
// @Tags social media
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Succes 200 {object} models.SocialMedia
// @Router /socialMedia [get]
func GetSocialMedias(ctx *gin.Context) {
	db := database.GetDB()

	SocialMedias := []models.SocialMedia{}

	err := db.Debug().Find(&SocialMedias).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": SocialMedias,
	})
}

// UpdateSocialMedia godoc
// @Summary Update social media for the given id
// @Description Update details of a social media corresponding to the input id
// @Tags social media
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the social media"
// @Succes 200 {object} models.Socialmedia
// @Router /socialMedia/{ID} [put]
func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType

	SocialMedia := models.SocialMedia{}
	socialMediaID, _ := strconv.Atoi(ctx.Param("socialMediaID"))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&SocialMedia)
	} else {
		ctx.ShouldBind(&SocialMedia)
	}

	err := db.Debug().Model(&SocialMedia).Where("id = ?", uint(socialMediaID)).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaURL: SocialMedia.SocialMediaURL}).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully update social media",
	})
}

// DeleteSocialMedia godoc
// @Summary Delete social media details for the given id
// @Description Delete details of a social media corresponding to the input id
// @Tags social media
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the social media"
// @Succes 200 {object} models.SocialMedia
// @Router /socialMedia/{ID} [delete]
func DeleteSocialMedia(ctx *gin.Context) {
	db := database.GetDB()

	SocialMedia := models.SocialMedia{}
	socialMediaID, _ := strconv.Atoi(ctx.Param("socialMediaID"))

	err := db.Debug().Where("id = ?", uint(socialMediaID)).Delete(&SocialMedia).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete social media",
	})
}
