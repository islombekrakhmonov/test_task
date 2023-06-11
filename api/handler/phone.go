package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"task/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Create phone
// @ID create_phone
// @Router /phone [POST]
// @Summary Create phone
// @Description Регистрация пользователя. Можно оставить пустым строку user_id или даже убрать строку, в любом случае будет использоваться текущий пользователь.
// @Tags Phone
// @Accept json
// @Produce json
// @Param agent body models.CreatePhoneRequest true "Request body"
// @Success 201 {object} models.CreatePhoneResponse "Response body"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *Handler) CreatePhone(c *gin.Context) {
	// Авторизация middleware чтобы проверить SESSTOKEN cookie
	err := h.Authenticate(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован", "message": err.Error()})
		return
	}

	var phone models.CreatePhoneRequest
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, ok := h.getUserIDFromToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
		return
	}

	createPhoneReq := models.CreatePhoneRequest{
		UserId:      userID,
		PhoneNumber: phone.PhoneNumber,
		Description: phone.Description,
		IsFax:       phone.IsFax,
	}
	// Проверка на дубликат номера телефона
	exists, err := h.storages.Phone().CheckDuplicatePhoneNumber(context.Background(), phone.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Такой номер уже существует"})
		return
	}

	resp, err := h.storages.Phone().Create(context.Background(), &createPhoneReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) Authenticate(c *gin.Context) error {
	sesstoken, err := c.Cookie("SESSTOKEN")
	if err != nil || sesstoken == "" {
		return errors.New("Не авторизован")
	}

	token, err := jwt.Parse(sesstoken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("kawasaki"), nil // Use the same secret key used for token generation
	})
	if err != nil || !token.Valid {
		return errors.New("Недействительный токен")
	}

	return nil
}



func (h *Handler) getUserIDFromToken(c *gin.Context) (string, bool) {
	tokenString, err := c.Cookie("SESSTOKEN")
	if err != nil {
		return "", false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("kawasaki"), nil
	})
	if err != nil {
		return "", false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", false
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", false
	}

	return userID, true
}

// Get By PhoneNumber Phone godoc
// @ID get_by_phoneNumber_phone
// @Router /phone [GET]
// @Summary Get By PhoneNumber 
// @Description Получение данных по номеру телефона
// @Tags Phone
// @Accept json
// @Produce json
// @Param q query string true "phone_number"
// @Success 200 {object} []models.GetByPhoneResponse "Response body"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *Handler) GetByPhone(c *gin.Context) {
	//middleware авторизации для проверки SESSTOKEN cookie
	err := h.Authenticate(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован", "message": err.Error()})
		return
	}

	// Получение номера телефона из параметра запроса
	phone := c.Query("q") 
	req := models.GetByPhoneRequest{PhoneNumber: phone}

	// Получение списка пользователей с указанным номером телефона
	resp, err := h.storages.Phone().GetByPhone(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": resp})
}

// Update Phone godoc
// @ID update_phone
// @Router /phone [PUT]
// @Summary Update Phone
// @Description Обновление данных номера
// @Tags Phone
// @Accept json
// @Produce json
// @Param phone body models.UpdatePhoneRequest true "Request body"
// @Success 202 {object} models.UpdatePhoneResponse "Response body
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *Handler) UpdatePhone(c *gin.Context) {
	var updatePhone models.UpdatePhoneRequest

	err := c.ShouldBindJSON(&updatePhone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"update phone": err.Error()})
		return
	}

	resp, err := h.storages.Phone().Update(context.Background(),&updatePhone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"update phone": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DELETE Phone godoc
// @ID delete_phone
// @Router /phone/{phone_id} [DELETE]
// @Summary Delete Phone
// @Description Удаление номера по phone_id
// @Tags Phone
// @Accept json
// @Produce json
// @Param phone_id path string true "phone_id"
// @Success 204 {object} string "Success Request"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *Handler) DeletePhone(c *gin.Context) {
	phone_id := c.Param("phone_id")
	
	err := h.storages.Phone().Delete(context.Background(),&models.PhonePKey{PhoneId: phone_id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"delete phone": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"delete phone": "Удаление прошло успешно"})
}
