package middlewares

import (
	"DTSGolang/FinalProject/database"
	"DTSGolang/FinalProject/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		photoID, err := strconv.Atoi(ctx.Param("photoID"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid photo parameter",
			})

			return
		}

		Photo := models.Photo{}
		err = db.Debug().First(&Photo, uint(photoID)).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Photo Not Found",
				"message": "photo doesn't exist",
			})

			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		admin := userData["admin"].(bool)
		if !admin {
			userID := uint(userData["id"].(float64))
			if Photo.UserID != userID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to update or delete this photo",
				})

				return
			}
		}

		ctx.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		commentID, err := strconv.Atoi(ctx.Param("commentID"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid comment parameter",
			})

			return
		}

		Comment := models.Comment{}
		err = db.Debug().First(&Comment, uint(commentID)).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Comment Not Found",
				"message": "comment doesn't exist",
			})

			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		admin := userData["admin"].(bool)
		if !admin {
			userID := uint(userData["id"].(float64))
			if Comment.UserID != userID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to update or delete this comment",
				})

				return
			}
		}

		ctx.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaID"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid social media parameter",
			})

			return
		}

		SocialMedia := models.SocialMedia{}
		err = db.Debug().First(&SocialMedia, uint(socialMediaID)).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Social Media Not Found",
				"message": "social media doesn't exist",
			})

			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		admin := userData["admin"].(bool)
		if !admin {
			userID := uint(userData["id"].(float64))
			if SocialMedia.UserID != userID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthoized",
					"message": "you are not allowed to update or delete this social media",
				})

				return
			}
		}

		ctx.Next()
	}
}
