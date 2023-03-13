package search

import (
	"sort"
	"youngzy.com/gohbsj/model"
)

func sortProductByPrice(products []model.Product) {
	sort.Sort(productSliceOrderByPrice(products))
}

type productSliceOrderByPrice []model.Product

func (p productSliceOrderByPrice) Len() int {
	return len(p)
}

func (p productSliceOrderByPrice) Less(i, j int) bool {
	return p[i].Price < p[j].Price
}

func (p productSliceOrderByPrice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
