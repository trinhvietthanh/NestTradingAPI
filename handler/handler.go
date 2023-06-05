package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}

	// Create an account group
	g := c.R.Group("/api/account")

	g.GET("/me", h.ME)
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/signout", h.Signout)
	g.POST("/token", h.Tokens)

	g.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"hello": "space persons",
		})
	})
}

func (h *Handler) ME(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

func (h *Handler) Signup(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"signup": "it's signup",
	})
}

func (h *Handler) Signin(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"signin": "it's signin",
	})
}

func (h *Handler) Signout(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "it's sigout",
	})
}

func (h *Handler) Tokens(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "it's token",
	})
}
