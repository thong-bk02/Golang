package main

import (
	"go-web/internal/handler"
	"go-web/internal/repository"
	"go-web/internal/service"
	"go-web/migrations"
	"go-web/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	db := utils.InitDB()
	db.AutoMigrate(&migrations.Account{})
	AccountRepo := repository.NewAccountRepository(db)
	accountSvc := service.NewAccountService(AccountRepo)
	accountHandler := handler.NewAccountHandler(accountSvc)
	r := gin.Default()
	router := r.Group("/api/test")
	{
		router.POST("/account/insert", accountHandler.Create)
		router.DELETE("/account/delete/:id", accountHandler.Delete)
		router.GET("/account/select/:id", accountHandler.Select)
		router.GET("/account/selectAll", accountHandler.SelectAll)
	}
	r.Run(":8080")
}
