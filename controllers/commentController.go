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

// CreateComment godoc
// @Summary Post a comment
// @Description post details of a new comment based on current user
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param models.Comment body models.Comment true "create a Comment"
// @Succes 201 {object} models.Comment
// @Router /comments [post]
func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Comment := models.Comment{}

	contentType := helpers.GetContentType(ctx)
	if contentType == appJSON {
		ctx.ShouldBindJSON(&Comment)
	} else {
		ctx.ShouldBind(&Comment)
	}

	Photo := models.Photo{}
	errPhoto := db.First(&Photo, "id = ?", Comment.PhotoID).Error
	if errPhoto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": errPhoto.Error(),
		})

		return
	}

	Comment.UserID = uint(userData["id"].(float64))

	errComment := db.Debug().Create(&Comment).Error
	if errComment != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": errComment.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Comment,
	})
}

// GetComment godoc
// @Summary Get comment details for the given id
// @Description Get details of a comment corresponding to the input id
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the comment"
// @Succes 200 {object} models.Comment
// @Router /comments/{ID} [get]
func GetComment(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType

	Comment := models.Comment{}
	commentID, _ := strconv.Atoi(ctx.Param("commentID"))

	err := db.Debug().First(&Comment, "id = ?", uint(commentID)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Comment Not Found",
				"message": err.Error(),
			})

			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Comment,
	})
}

// GetMyComments godoc
// @Summary Get all comment corresponding user
// @Description Get all comment data corresponding user
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Succes 200 {object} models.Comment
// @Router /comments/my [get]
func GetMyComments(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Comments := []models.Comment{}

	userID := uint(userData["id"].(float64))

	err := db.Where("user_id = ?", userID).Find(&Comments).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Comments,
	})
}

// GetComments godoc
// @Summary Get all comment
// @Description Get all comment data
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Succes 200 {object} models.Comment
// @Router /comments [get]
func GetComments(ctx *gin.Context) {
	db := database.GetDB()

	Comments := []models.Comment{}

	err := db.Debug().Find(&Comments).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Comments,
	})
}

// UpdateComment godoc
// @Summary Update comment for the given id
// @Description Update details of a comment corresponding to the input id
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the comment"
// @Succes 200 {object} models.Comment
// @Router /comments/{ID} [put]
func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()

	Comment := models.Comment{}
	commentID, _ := strconv.Atoi(ctx.Param("commentID"))

	contentType := helpers.GetContentType(ctx)
	if contentType == appJSON {
		ctx.ShouldBindJSON(&Comment)
	} else {
		ctx.ShouldBind(&Comment)
	}

	comment := models.Comment{}

	errComment := db.First(&comment, "id = ?", uint(commentID)).Error
	if errComment != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": errComment.Error(),
		})
	}

	Comment.PhotoID = comment.PhotoID

	Photo := models.Photo{}
	errPhoto := db.First(&Photo, "id = ?", Comment.PhotoID).Error
	if errPhoto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": errPhoto.Error(),
		})

		return
	}

	err := db.Model(&Comment).Where("id = ?", uint(commentID)).Updates(models.Comment{Message: Comment.Message}).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully update comment",
	})
}

// DeleteComments godoc
// @Summary Delete comment details for the given id
// @Description Delete details of a comment corresponding to the input id
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the comment"
// @Succes 200 {object} models.Comment
// @Router /comments/{ID} [delete]
func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()

	Comment := models.Comment{}
	commentID, _ := strconv.Atoi(ctx.Param("commentID"))

	err := db.Debug().Where("id = ?", uint(commentID)).Delete(&Comment).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted comment",
	})
}
