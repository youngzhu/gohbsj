package matcher

import (
	"fmt"
	"net/url"
	"os"
	"testing"
	"youngzy.com/gohbsj/model"
)

//func TestJdMatcher_Search(t *testing.T) {
//	jd := jdMatcher{}
//
//	_, err := jd.Search("美的BCD-217WTMA")
//	if err != nil {
//		t.Error(err)
//	}
//}

func TestJdMatcher_Search_parseHtml(t *testing.T) {
	f, err := os.Open("testdata/jd.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	products, _ := parseJDHtml(f)
	firstItem := products[0]

	want := model.Product{
		Name:        "美的（Midea）217升小冰箱 新风冷无霜节能低音三开门三门三温 小型家用电冰箱 BCD-217WTMA 以旧换新",
		Price:       1899.00,
		Vendor:      "南百冰洗电器京东自营专区",
		OriginalURL: "//item.jd.com/100030552515.html",
	}

	if firstItem.Name != want.Name {
		t.Errorf("Product name, want: %s, got: %s", want.Name, firstItem.Name)
	}
	if firstItem.Price != want.Price {
		t.Errorf("Product price, want: %.2f, got: %.2f", want.Price, firstItem.Price)
	}
	if firstItem.Vendor != want.Vendor {
		t.Errorf("Product vendor, want: %s, got: %s", want.Vendor, firstItem.Vendor)
	}
	if firstItem.OriginalURL != want.OriginalURL {
		t.Errorf("Product originalUrl, want: %s, got: %s", want.OriginalURL, firstItem.OriginalURL)
	}
}

// 产品的<li>里面还有<li>，导致解析出的数据重复
func TestJdMatcher_Search_parseHtml_duplicate(t *testing.T) {
	f, err := os.Open("testdata/jd_duplicate.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	products, _ := parseJDHtml(f)
	if len(products) != 30 {
		t.Errorf("Products size, want: 30, got: %d", len(products))
	}

}

func TestJdMatcher_Search_encode(t *testing.T) {
	keyword := "华为 保时捷"

	before := "https://search.jd.com/Search?%s&enc=utf-8&spm=a.0.0&pvid=eaeabb3cc04e4c07bf09da2684c471d8"

	v := url.Values{}
	v.Set("keyword", keyword)
	after := fmt.Sprintf(before, v.Encode())
	t.Log(after)
}
