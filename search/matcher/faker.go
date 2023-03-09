package matcher

import (
	"github.com/spf13/viper"
	"log"
	"strings"
	"youngzy.com/gohbsj/model"
	"youngzy.com/gohbsj/search"
)

func init() {
	var matcher fakeMatcher
	search.Register("fake", matcher)
}

type fakeMatcher struct{}

func (m fakeMatcher) Search(searchTerm string) ([]*model.Product, error) {
	log.Println("===" + viper.GetString("faker.info") + "===")
	log.Println("===" + viper.GetString("system.name") + "===")
	log.Println("===" + viper.GetString("test") + "===")

	var products []*model.Product

	for _, prod := range fakeProducts {
		if strings.Contains(prod.Name, searchTerm) {
			products = append(products, prod)
		}
	}

	return products, nil
}

var fakeProducts = []*model.Product{
	{Name: "iPhone", Price: 8999.99, Vendor: "苹果旗舰店", Provider: model.JD, OriginalURL: "www.youngzy.com"},
	{Name: "iPhone", Price: 7999.99, Vendor: "苹果旗舰店", Provider: model.TaoBao, OriginalURL: "www.youngzy.com"},
	{Name: "iPhone Pro 14", Price: 11999.99, Vendor: "京东自营", Provider: model.JD, OriginalURL: "www.youngzy.com"},
	{Name: "华为 保时捷", Price: 16888.99, Vendor: "京东自营", Provider: model.JD, OriginalURL: "www.youngzy.com"},
}
