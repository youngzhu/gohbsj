package provider

import "youngzy.com/gohbsj/model"

// Information Provider 信息供应商

func GetProviders(products []model.Product) (providers []string) {
	aMap := map[string]string{}
	for _, p := range products {
		if _, ok := aMap[p.Provider]; !ok {
			aMap[p.Provider] = p.Provider
			providers = append(providers, p.Provider)
		}
	}
	return
}
