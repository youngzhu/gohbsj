package provider

import "youngzy.com/gohbsj/model"

// Provider Information Provider 信息供应商
type Provider interface {
	GetProducts(keywords string) []model.Product
	//GetProviders() []string
}
