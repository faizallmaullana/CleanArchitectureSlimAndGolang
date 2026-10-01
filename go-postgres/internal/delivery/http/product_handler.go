package httpdelivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/faizallmaullana/rekrutment-gbu-go/internal/domain"
	"github.com/faizallmaullana/rekrutment-gbu-go/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *usecase.ProductService
}

func NewProductHandler(service *usecase.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (handler *ProductHandler) RegisterRoutes(router *gin.Engine) {
	products := router.Group("/api/products")
	products.GET("", handler.list)
	products.GET("/:id", handler.get)
	products.POST("", handler.create)
	products.PUT("/:id", handler.update)
	products.DELETE("/:id", handler.delete)
}

func (handler *ProductHandler) list(context *gin.Context) {
	products, err := handler.service.List(context.Request.Context())
	if err != nil {
		writeServiceError(context, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"products": products})
}

func (handler *ProductHandler) get(context *gin.Context) {
	id, err := parseID(context.Param("id"))
	if err != nil {
		writeError(context, http.StatusBadRequest, err)
		return
	}

	product, err := handler.service.Get(context.Request.Context(), id)
	if err != nil {
		writeServiceError(context, err)
		return
	}

	context.JSON(http.StatusOK, product)
}

func (handler *ProductHandler) create(context *gin.Context) {
	input, err := decodeProductInput(context)
	if err != nil {
		writeError(context, http.StatusBadRequest, err)
		return
	}

	product, err := handler.service.Create(context.Request.Context(), input.Name, *input.Price)
	if err != nil {
		writeServiceError(context, err)
		return
	}

	context.JSON(http.StatusCreated, product)
}

func (handler *ProductHandler) update(context *gin.Context) {
	id, err := parseID(context.Param("id"))
	if err != nil {
		writeError(context, http.StatusBadRequest, err)
		return
	}

	input, err := decodeProductInput(context)
	if err != nil {
		writeError(context, http.StatusBadRequest, err)
		return
	}

	product, err := handler.service.Update(context.Request.Context(), id, input.Name, *input.Price)
	if err != nil {
		writeServiceError(context, err)
		return
	}

	context.JSON(http.StatusOK, product)
}

func (handler *ProductHandler) delete(context *gin.Context) {
	id, err := parseID(context.Param("id"))
	if err != nil {
		writeError(context, http.StatusBadRequest, err)
		return
	}

	if err := handler.service.Delete(context.Request.Context(), id); err != nil {
		writeServiceError(context, err)
		return
	}

	context.Status(http.StatusNoContent)
}

type productInput struct {
	Name  string   `json:"name" binding:"required"`
	Price *float64 `json:"price" binding:"required"`
}

func decodeProductInput(context *gin.Context) (productInput, error) {
	var input productInput
	if err := context.ShouldBindJSON(&input); err != nil {
		return productInput{}, errors.New("request body must contain a valid name and price")
	}

	return input, nil
}

func parseID(value string) (int64, error) {
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("product id must be a positive integer")
	}

	return id, nil
}

func writeServiceError(context *gin.Context, err error) {
	if errors.Is(err, domain.ErrNotFound) {
		writeError(context, http.StatusNotFound, err)
		return
	}
	if errors.Is(err, domain.ErrInvalid) {
		writeError(context, http.StatusBadRequest, err)
		return
	}

	writeError(context, http.StatusInternalServerError, err)
}

func writeError(context *gin.Context, status int, err error) {
	context.JSON(status, gin.H{"error": err.Error()})
}
