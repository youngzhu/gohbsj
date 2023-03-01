package service

import "log"

type ProviderHandler struct {
	URLGenerator
}

type categoryTemplateContext struct {
	Providers        []string
	SelectedProvider string
	ProviderUrlFunc  func(string) string
}

func (handler ProviderHandler) GetButtons(selected string) ActionResult {
	return NewTemplateAction("provider_buttons.html",
		categoryTemplateContext{
			Providers: []string{"all"},
			//SelectedProvider: selected,
			SelectedProvider: "all",
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
