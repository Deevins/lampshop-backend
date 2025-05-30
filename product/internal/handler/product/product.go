package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Заглушки
func (h *Handler) ListProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List products"})
}

func (h *Handler) CreateProduct(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}

func (h *Handler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get product", "id": id})
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Product updated", "id": id})
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted", "id": id})
}

func (h *Handler) GetUploadURL(c *gin.Context) {
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filename is required"})
		return
	}

	//url, err := h.MinioClient.PresignedPutObject(c, h.BucketName, filename, time.Minute*15)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{"url": "https://test.pdf?X-Amz-Algorithm"})
}

func (h *Handler) NotifyUpload(c *gin.Context) {
	var payload struct {
		Filename  string `json:"filename"`
		ProductID string `json:"product_id"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// здесь можно связать файл с товаром, записать в БД и т.п.
	c.JSON(http.StatusOK, gin.H{
		"message":   "File upload recorded",
		"productId": payload.ProductID,
		"filename":  payload.Filename,
	})
}
