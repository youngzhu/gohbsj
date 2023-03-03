package matcher

import (
	"youngzy.com/gohbsj/model"
	"youngzy.com/gohbsj/search"
)

func init() {
	var matcher fakeMatcher
	search.Register("fake", matcher)
}

type fakeMatcher struct{}

func (m fakeMatcher) Search(keywords string) ([]*model.Product, error) {

	return fakeProducts, nil
}

var fakeProducts = []*model.Product{
	{Name: "iPhone", Price: 8999.99, Vendor: "苹果旗舰店", Provider: model.JD, OriginalURL: "www.youngz.com"},
	{Name: "iPhone", Price: 7999.99, Vendor: "苹果旗舰店", Provider: model.TaoBao, OriginalURL: "www.youngz.com"},
	{Name: "iPhone Pro 14", Price: 11999.99, Vendor: "京东自营", Provider: model.JD, OriginalURL: "www.youngz.com"},
	{Name: "华为 保时捷", Price: 16888.99, Vendor: "京东自营", Provider: model.JD, OriginalURL: "www.youngz.com"},
}
