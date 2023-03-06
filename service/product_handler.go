package service

import (
	"log"
	"youngzy.com/gohbsj/model"
	"youngzy.com/gohbsj/search"
)

type ProductHandler struct {
	Searcher     *search.Searcher
	URLGenerator URLGenerator
}

type ProductTemplateContext struct {
	SearchTerm       string
	Products         []model.Product
	SelectedProvider string
}

func (handler ProductHandler) GetProducts(searchTerm, providerID string) ActionResult {
	log.Println("provider:", providerID)

	//handler.Searcher.Run("")

	//log.Println("products:", len(handler.Searcher.GetProducts()))

	return NewTemplateAction("product_list.html",
		ProductTemplateContext{
			SearchTerm:       searchTerm,
			Products:         handler.Searcher.GetProducts(providerID),
			SelectedProvider: providerID,
		})
}

// SearchRef 居然是这样接POST参数。。
type SearchRef struct {
	SearchTerm string
}

func (handler ProductHandler) PostSearch(sr SearchRef) ActionResult {
	log.Println("searchTerm:", sr.SearchTerm)

	handler.Searcher.Run(sr.SearchTerm)

	return NewTemplateAction("product_list.html",
		ProductTemplateContext{
			SearchTerm:       sr.SearchTerm,
			Products:         handler.Searcher.GetProducts(""),
			SelectedProvider: "",
		})
}
