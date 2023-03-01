package service

type ProviderHandler struct {
	URLGenerator
}

type categoryTemplateContext struct {
	Providers        []string
	SelectedProvider int
	CategoryUrlFunc  func(int) string
}

func (handler ProviderHandler) GetButtons(selected int) ActionResult {
	return NewTemplateAction("provider_buttons.html",
		categoryTemplateContext{
			Providers:        []string{"test"},
			SelectedProvider: selected,
			CategoryUrlFunc:  handler.createCategoryFilterFunction(),
		})
}

func (handler ProviderHandler) createCategoryFilterFunction() func(int) string {
	return func(category int) string {
		url, _ := handler.GenerateUrl(ProductHandler.GetProducts,
			category, 1)
		return url
	}
}
