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
	SearchTerm string

	Products         []model.Product
	SelectedProvider string

	Cost float64
}

func (handler ProductHandler) GetProducts(providerID string) ActionResult {
	//log.Println("provider:", providerID)
	return NewTemplateAction("product_list.html",
		ProductTemplateContext{
			SearchTerm:       handler.Searcher.SearchTerm,
			Products:         handler.Searcher.GetProducts(providerID),
			SelectedProvider: providerID,
			Cost:             handler.Searcher.Cost.Seconds(),
		})
}

// SearchRef 居然是这样接POST参数。。
type SearchRef struct {
	SearchTerm string
}

func (handler ProductHandler) PostSearch(sr SearchRef) ActionResult {
	log.Println("searchTerm:", sr.SearchTerm)

	handler.Searcher.SearchTerm = sr.SearchTerm
	handler.Searcher.Run()

	return NewTemplateAction("product_list.html",
		ProductTemplateContext{
			SearchTerm:       sr.SearchTerm,
			Products:         handler.Searcher.GetProducts(""),
			SelectedProvider: "",
			Cost:             handler.Searcher.Cost.Seconds(),
		})
}
