package matcher

import (
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
