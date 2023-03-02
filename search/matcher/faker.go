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

var fakeProducts = []*model.Product{
	{Name: "iPhone", Price: 8999.99, Vendor: "苹果旗舰店", Provider: "京东", OriginalURL: "www.youngz.com"},
	{Name: "iPhone", Price: 7999.99, Vendor: "苹果旗舰店", Provider: "淘宝", OriginalURL: "www.youngz.com"},
	{Name: "iPhone Pro 14", Price: 11999.99, Vendor: "京东自营", Provider: "京东", OriginalURL: "www.youngz.com"},
	{Name: "华为 保时捷", Price: 16888.99, Vendor: "京东自营", Provider: "京东", OriginalURL: "www.youngz.com"},
}

func (m fakeMatcher) Search(keywords string) ([]*model.Product, error) {

	return fakeProducts, nil
}
