package main

import (
	"bytes"
	"context"
	"github.com/Deevins/lampshop-backend/admin/internal/handler"
	"github.com/Deevins/lampshop-backend/admin/internal/infra"
	"github.com/Deevins/lampshop-backend/admin/internal/repository"
	"github.com/Deevins/lampshop-backend/admin/internal/service"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	dbPool, err := infra.NewPostgresPool(ctx)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbPool.Close()

	authRepo := repository.NewAuthRepository(dbPool)
	attrRepo := repository.NewAttributeRepository(dbPool)
	authService := service.NewAuthService(authRepo)
	categoryService := service.NewCategoryService(attrRepo)
	authHandler := handler.NewAuthHandler(authService)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	orderServiceURL := "http://order-service:8082"
	productServiceURL := "http://product-service:8081"

	orderClient := infra.NewOrderServiceClient(orderServiceURL)
	productClient := infra.NewProductServiceClient(productServiceURL)

	router := gin.Default()
	//router.Use(corsMiddleware())

	router.POST("/login", authHandler.Login)
	router.GET("/categories", categoryHandler.GetCategories)
	router.GET("/categories/:id/attributes", categoryHandler.GetAttributeOptions)

	protected := router.Group("/")
	//protected.Use(AuthMiddleware())

	protected.GET("/orders", func(c *gin.Context) {
		//token := extractBearerToken(c.Request.Header.Get("Authorization"))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		resp, err := orderClient.GetAllOrders(token)
		infra.ProxyResponse(c, resp, err)
	})

	protected.PUT("/orders/:id/status", func(c *gin.Context) {
		//token := extractBearerToken(c.Request.Header.Get("Authorization"))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		orderID := c.Param("id")
		bodyBuf := new(bytes.Buffer)
		_, _ = io.Copy(bodyBuf, c.Request.Body)

		resp, err := orderClient.UpdateOrderStatus(token, orderID, bodyBuf)
		infra.ProxyResponse(c, resp, err)
	})

	protected.GET("/orders/active", func(c *gin.Context) {
		//token := extractBearerToken(c.Request.Header.Get("Authorization"))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		resp, err := orderClient.GetActiveOrders(token)
		infra.ProxyResponse(c, resp, err)
	})
	protected.GET("/orders/:id/status", func(c *gin.Context) {
		//token := extractBearerToken(c.Request.Header.Get("Authorization"))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		orderID := c.Param("id")
		resp, err := orderClient.GetOrderStatus(token, orderID)
		infra.ProxyResponse(c, resp, err)
	})
	protected.POST("/checkout", func(c *gin.Context) {
		//token := extractBearerToken(c.Request.Header.Get("Authorization"))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		bodyBuf := new(bytes.Buffer)
		_, _ = io.Copy(bodyBuf, c.Request.Body)
		resp, err := orderClient.CreateOrder(token, bodyBuf)
		infra.ProxyResponse(c, resp, err)
	})

	protected.GET("/products", func(c *gin.Context) {
		//token := extractBearerToken(c.Request.Header.Get("Authorization"))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		resp, err := productClient.ListProducts(token)
		infra.ProxyResponse(c, resp, err)
	})
	protected.POST("/products", func(c *gin.Context) {
		//token := extractBearerToken(c.Request.Header.Get("Authorization"))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		bodyBuf := new(bytes.Buffer)
		_, _ = io.Copy(bodyBuf, c.Request.Body)
		resp, err := productClient.CreateProduct(token, bodyBuf)
		infra.ProxyResponse(c, resp, err)
	})
	protected.GET("/products/:id", func(c *gin.Context) {
		//token := extractBearerToken(c.Request.Header.Get("Authorization"))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		productID := c.Param("id")
		resp, err := productClient.GetProductByID(token, productID)
		infra.ProxyResponse(c, resp, err)
	})
	protected.PUT("/products/:id", func(c *gin.Context) {
		//token := extractBearerToken(c.Request.Header.Get("Authorization"))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		productID := c.Param("id")
		bodyBuf := new(bytes.Buffer)
		_, _ = io.Copy(bodyBuf, c.Request.Body)
		resp, err := productClient.UpdateProduct(token, productID, bodyBuf)
		infra.ProxyResponse(c, resp, err)
	})
	protected.DELETE("/products/:id", func(c *gin.Context) {
		//token := extractBearerToken(c.Request.Header.Get("Authorization"))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		productID := c.Param("id")
		resp, err := productClient.DeleteProduct(token, productID)
		infra.ProxyResponse(c, resp, err)
	})

	protected.GET("/products/upload-url", func(c *gin.Context) {
		//token := extractBearerToken(c.Request.Header.Get("Authorization"))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		resp, err := productClient.GetUploadURL(token)
		infra.ProxyResponse(c, resp, err)
	})
	protected.POST("/products/notify-upload", func(c *gin.Context) {
		//token := extractBearerToken(c.Request.Header.Get("Authorization"))
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
		bodyBuf := new(bytes.Buffer)
		_, _ = io.Copy(bodyBuf, c.Request.Body)
		resp, err := productClient.NotifyUpload(token, bodyBuf)
		infra.ProxyResponse(c, resp, err)
	})

	router.Run(":" + "8083")
}

// AuthMiddleware проверяет заголовок Authorization Bearer <token>
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token required"})
			c.Abort()
			return
		}
		tokenString := authHeader[7:]
		username, err := infra.ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Next()
	}
}

// extractBearerToken убирает "Bearer " префикс и возвращает сам токен.
func extractBearerToken(headerValue string) string {
	if len(headerValue) > 7 && headerValue[:7] == "Bearer " {
		return headerValue[7:]
	}
	return ""
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		hdr := c.Writer.Header()

		if _, ok := hdr["Access-Control-Allow-Origin"]; !ok {
			// Нет такого ключа — ставим
			hdr.Set("Access-Control-Allow-Origin", "*")
		}
		// Аналогично для других CORS-заголовков:
		if _, ok := hdr["Access-Control-Allow-Methods"]; !ok {
			hdr.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		}
		if _, ok := hdr["Access-Control-Allow-Headers"]; !ok {
			hdr.Set("Access-Control-Allow-Headers", "*")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
