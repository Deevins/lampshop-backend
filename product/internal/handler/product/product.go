package product

import (
	"github.com/Deevins/lampshop-backend/product/internal/service/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

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

func (h *Handler) ListProducts(c *gin.Context) {
	products, err := h.service.ListProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot list products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	product, err := h.service.GetProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var req *sql.CreateProductParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if req.ID == uuid.Nil {
		req.ID = uuid.New()
	}
	if req.CreatedAt.IsZero() {
		req.CreatedAt = req.UpdatedAt
	}
	if req.UpdatedAt.IsZero() {
		req.UpdatedAt = req.CreatedAt
	}

	err := h.service.CreateProduct(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "product created"})
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	var req *sql.UpdateProductParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	req.ID = id

	err = h.service.UpdateProduct(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product updated"})
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	err = h.service.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

// --- Категории ---

func (h *Handler) ListCategories(c *gin.Context) {
	categories, err := h.service.ListCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot list categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var req *sql.CreateCategoryParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if req.ID == uuid.Nil {
		req.ID = uuid.New()
	}

	err := h.service.CreateCategory(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create category"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "category created"})
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	err = h.service.DeleteCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}
