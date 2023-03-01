package service

import (
	"log"
	"youngzy.com/gohbsj/model"
)

type ProductHandler struct {
	URLGenerator
}

type ProductTemplateContext struct {
	Products         []model.Product
	SelectedProvider string
}

func (handler ProductHandler) GetProducts(provider string) ActionResult {
	log.Println("provider:", provider)

	return NewTemplateAction("product_list.html",
		ProductTemplateContext{
			Products:         model.TestProducts,
			SelectedProvider: provider,
		})
}
