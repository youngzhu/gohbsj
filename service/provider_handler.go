package service

import (
	"log"
	"youngzy.com/gohbsj/model"
	"youngzy.com/gohbsj/search"
)

type ProviderHandler struct {
	Searcher     *search.Searcher
	URLGenerator URLGenerator
}

type providerTemplateContext struct {
	SearchTerm       string
	Providers        []model.Provider
	SelectedProvider string
	ProviderUrlFunc  func(string) string
}

func (handler ProviderHandler) GetButtons(selected, searchTerm string) ActionResult {
	log.Println("ProviderHandler searchTerm:", searchTerm)

	return NewTemplateAction("provider_buttons.html",
		providerTemplateContext{
			SearchTerm:       searchTerm,
			Providers:        handler.Searcher.GetProviders(),
			SelectedProvider: selected,
			ProviderUrlFunc:  handler.createProviderFilterFunction(),
		})
}

func (handler ProviderHandler) createProviderFilterFunction() func(string) string {
	return func(provider string) string {
		url, err := handler.URLGenerator.GenerateUrl(ProductHandler.GetProducts, provider)
		if err != nil {
			panic(err)
		}
		//log.Println(url)
		return url
	}
}
