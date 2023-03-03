package service

import (
	"log"
	"youngzy.com/gohbsj/model"
	"youngzy.com/gohbsj/search"
)

type ProviderHandler struct {
	search.Searcher
	URLGenerator
}

type providerTemplateContext struct {
	Providers        []model.Provider
	SelectedProvider string
	ProviderUrlFunc  func(string) string
}

func (handler ProviderHandler) GetButtons(selected string) ActionResult {
	handler.Searcher.Run("")

	return NewTemplateAction("provider_buttons.html",
		providerTemplateContext{
			Providers:        handler.Searcher.GetProviders(),
			SelectedProvider: selected,
			ProviderUrlFunc:  handler.createProviderFilterFunction(),
		})
}

func (handler ProviderHandler) createProviderFilterFunction() func(string) string {
	return func(provider string) string {
		url, err := handler.GenerateUrl(ProductHandler.GetProducts, provider)
		if err != nil {
			panic(err)
		}
		log.Println(url)
		return url
	}
}
