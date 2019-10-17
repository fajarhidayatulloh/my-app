package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/my-app/helpers"
	"github.com/my-app/infrastructures"
	"github.com/my-app/models"
	"github.com/my-app/repositories"
	"github.com/my-app/services"
	"github.com/thedevsaddam/govalidator"
)

//InitProductController init
func InitProductController() *ProductController {
	productRepository := new(repositories.ProductRepository)
	productRepository.DB = &infrastructures.SQLConnection{}

	productService := new(services.ProductService)
	productService.ProductRepository = productRepository

	productController := new(ProductController)
	productController.ProductService = productService

	return productController
}

// ProductController behaviour
type ProductController struct {
	ProductService services.IProductService
}

// StoreProduct is new product
func (r *ProductController) StoreProduct(res http.ResponseWriter, req *http.Request) {
	var product models.ProductInput

	rules := govalidator.MapData{
		"name":  []string{"required"},
		"sku":   []string{"required"},
		"price": []string{"required"},
	}

	opts := govalidator.Options{
		Request:         req,
		Rules:           rules,
		RequiredDefault: true,
		Data:            &product,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()

	if len(e) != 0 {
		err := map[string]interface{}{"validationError": e}
		helpers.Response(res, http.StatusUnprocessableEntity, err)
		return
	}

	result, err := r.ProductService.StoreProduct(product)
	if err != nil {
		rs := map[string]interface{}{
			"err":      "error_store_product",
			"err_desc": fmt.Sprintf("%s", err),
		}
		helpers.Response(res, http.StatusBadRequest, rs)
		return
	}

	helpers.Response(res, http.StatusCreated, result)
	return
}

// ProductList is list of product
func (r *ProductController) ProductList(res http.ResponseWriter, req *http.Request) {
	product, _ := r.ProductService.ProductList()
	rs := map[string]interface{}{
		"data": product,
	}
	helpers.Response(res, http.StatusOK, rs)
}

// GetProductByID is
func (r *ProductController) GetProductByID(res http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)
	id, err := strconv.Atoi(param["id"])

	if err != nil {
		helpers.Response(res, http.StatusBadRequest, err)
		return
	}

	product, err := r.ProductService.GetProductByID(id)
	if err != nil {
		helpers.Response(res, http.StatusBadRequest, err)
		return
	}

	if product.ID == 0 {
		if err != nil {
			helpers.Response(res, http.StatusBadRequest, err)
			return
		}

		helpers.Response(res, http.StatusOK, err)
		return
	}

	rs := map[string]interface{}{
		"data": product,
	}

	helpers.Response(res, http.StatusOK, rs)
	return
}
