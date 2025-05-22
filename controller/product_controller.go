package controller

import (
	"api-go/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}
