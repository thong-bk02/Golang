package handler

import (
	"fmt"
	"go-web/internal/service"
	"go-web/migrations"
	"net/http"

	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	svc service.AccountService
}

func NewAccountHandler(svc service.AccountService) *accountHandler {
	return &accountHandler{svc: svc}
}

func (h *accountHandler) Create(c *gin.Context) {
	var input migrations.Account
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu gửi lên không đúng định dạng"})
		return
	}

	if err := h.svc.SubmitForm(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu dữ liệu vào YugabyteDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tạo tài khoản thành công!"})
}

func (h *accountHandler) SelectAll(c *gin.Context) {
	accounts, err := h.svc.GetAllAccounts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, accounts)
}
func (h *accountHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID không hợp lệ"})
		return
	}
	if err := h.svc.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Xóa tài khoản thành công!"})
}

func (h *accountHandler) Select(c *gin.Context) {
	account, error := h.svc.Select(1)
	if error != nil {
		c.JSON(500, gin.H{"error": error.Error()})
		return
	}
	c.JSON(200, account)
}
