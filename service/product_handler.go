package service

import (
	"log"
	"youngzy.com/gohbsj/model"
	"youngzy.com/gohbsj/search"
)

type ProductHandler struct {
	search.Searcher
	URLGenerator
}

type ProductTemplateContext struct {
	Products         []model.Product
	SelectedProvider string
}

func (handler ProductHandler) GetProducts(providerID string) ActionResult {
	log.Println("provider:", providerID)

	handler.Searcher.Run("")

	//log.Println("products:", len(handler.Searcher.GetProducts()))

	return NewTemplateAction("product_list.html",
		ProductTemplateContext{
			Products:         handler.Searcher.GetProducts(providerID),
			SelectedProvider: providerID,
		})
}
