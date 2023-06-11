package handler

import (
	"context"
	"fmt"
	"net/http"
	"task/models"
	"task/security"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Create user
// @ID create_user
// @Router /register [POST]
// @Summary Create user
// @Description Регистрация пользователя
// @Tags User
// @Accept json
// @Produce json
// @Param agent body models.CreateUserRequest true "Request body"
// @Success 201 {object} models.CreateUserResponse "Response body"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *Handler) Create(c *gin.Context) {
	var user models.CreateUserRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"storage.user.create": err.Error()})
		return
	}

	resp, err := h.storages.User().Create(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"storage.user.create": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Login godoc
// @ID login
// @Router /auth [POST]
// @Tags User
// @Summary  login
// @Description  Авторизация пользователя
// @Accept json
// @Param credentials body models.AuthUserRequest true "credentials"
// @Produce json
// @Success 200 {object} models.AuthUserResponse "Response body"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *Handler) Login(c *gin.Context) {
	var login models.AuthUserRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Авторизация": err.Error()})
		return
	}

	user, err := h.storages.User().Login(context.Background(), &login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Авторизация": err.Error()})
		return
	}


	if user.Login == "" || !ComparePasswords([]byte(user.Password), login.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
		return
	}

	token, err := security.GenerateToken(user.UserId, user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := security.GenerateRefreshToken(user.UserId) // Generate a new refresh token for the user
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.AuthUserResponse{
		UserId: user.UserId,
		Login: user.Login,
		Token: token,
		RefreshToken: refreshToken,
	}

	c.SetCookie("SESSTOKEN", token, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, response)

}


func ComparePasswords(hashedPassword []byte, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(plainPassword))
	return err == nil
}


// Get user by name
// @ID get_user_by_name
// @Router /user/{name} [GET]
// @Summary Get user by name
// @Description Получение пользователя по имени
// @Tags User
// @Accept json
// @Produce json
// @Param name path string true "name"
// @Success 200 {object} []models.GetUserByNameResponse "Success Request"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *Handler) GetByName(c *gin.Context) {
	name := c.Param("name")

	resp, err := h.storages.User().GetByName(context.Background(), &models.GetUserByNameRequest{Name: name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"storage.user.GetByName": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}


func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.URL.Path == "/doc.json" {
			c.Next()
			return
		}

		sesstoken, err := c.Cookie("SESSTOKEN")
		if err != nil || sesstoken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
			c.Abort()
			return
		}

		
		// Verify the JWT token
		token, err := jwt.Parse(sesstoken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("kawasaki"), nil 
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
			c.Abort()
			return
		}

		c.Next()
	}
}


