package httpx

import (
	"go-hex-temp/internal/adapters/in/httpx/driver"
	"go-hex-temp/internal/infrastructure/logx"
	"go-hex-temp/internal/ports/input"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService input.UserService
}
type userRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Bio    string `json:"bio"`
	Avatar string `json:"avatar"`
}

func NewUserHandler(useCase input.UserService) *userHandler {
	return &userHandler{useCase}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	req := new(userRequest)
	if err := c.Bind(req); err != nil {
		logx.Errorf("Bind failed: %v", err)
		return
	}

	c.JSON(http.StatusOK, req)
}

func (h *userHandler) FindUsers(c *gin.Context) {

	query := driver.ClaimQuery(c)
	users, err := h.userService.GetUsers(query)
	if err != nil {
		logx.Errorf("compile query failed: %v", err)
		return
	}
	_ = users
	c.JSON(200, query)

}
