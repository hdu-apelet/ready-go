package api

import (
	"errors"
	"fmt"
	"net/http"
	"readygo/internal/database"
	"readygo/internal/route"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	route.Register("/api",
		func(i gin.IRouter) {
			i.GET("/ping", ping())

			i.POST("/register", func(c *gin.Context) {
				reqbody := struct {
					Username string `json:"username" form:"username" binding:"required"`
					Password string `json:"password" form:"password" binding:"required"`
				}{}

				if err := c.BindJSON(&reqbody); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"code": 400,
						"msg":  fmt.Sprintf("invalid payload: %v", err),
					})
					return
				}

				user := database.User{}
				db := database.Get()
				if err := db.Where("username = ?", reqbody.Username).First(&user).Error; err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						c.JSON(http.StatusInternalServerError, gin.H{
							"code": 500,
							"msg":  fmt.Sprintf("database error: %v", err),
						})
						return
					}
				}

				if user.ID != 0 {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 400,
						"msg":  "username already exists",
					})
					return
				}

				user.Username = reqbody.Username
				// md5sum
				user.Password = reqbody.Password

				if err := db.Create(&user).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg":  fmt.Sprintf("create user: %v", err),
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "success",
					"data": user,
				})
			})
		},
	)
}

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	}
}
