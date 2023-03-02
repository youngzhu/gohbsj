package service

import (
	"log"
	"youngzy.com/gohbsj/search"
)

type ProviderHandler struct {
	search.Searcher
	URLGenerator
}

type categoryTemplateContext struct {
	Providers        []string
	SelectedProvider string
	ProviderUrlFunc  func(string) string
}

func (handler ProviderHandler) GetButtons(selected string) ActionResult {
	handler.Searcher.Run("")

	return NewTemplateAction("provider_buttons.html",
		categoryTemplateContext{
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
