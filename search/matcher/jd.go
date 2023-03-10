package matcher

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"youngzy.com/gohbsj/model"
	"youngzy.com/gohbsj/search"
)

func init() {
	var matcher jdMatcher
	search.Register("jd", matcher)
}

type jdMatcher struct{}

var header = http.Header{
	//":authority":                {"search.jd.com"},
	//":method":                   {"GET"},
	//":path":                     {"/Search?keyword=%E8%9A%95%E4%B8%9D%E8%A2%AB&enc=utf-8&spm=a.0.0&pvid=eaeabb3cc04e4c07bf09da2684c471d8"},
	//":scheme":                   {"https"},
	"accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
	"accept-encoding":           {"gzip, deflate, br"},
	"accept-language":           {"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6"},
	"referer":                   {"https://www.jd.com/"},
	"sec-ch-ua":                 {"\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Microsoft Edge\";v=\"110\""},
	"sec-ch-ua-mobile":          {"?0"},
	"sec-ch-ua-platform":        {"\"Windows\""},
	"sec-fetch-dest":            {"document"},
	"sec-fetch-mode":            {"navigate"},
	"sec-fetch-site":            {"same-site"},
	"sec-fetch-user":            {"?1"},
	"upgrade-insecure-requests": {"1"},
	"user-agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.63"},
}

const searchUrlFormat = "https://search.jd.com/Search?%s&enc=utf-8&%s&pvid=edcd0d2fd3894407958588dea052ca80"

func (m jdMatcher) Search(searchTerm string) ([]*model.Product, error) {
	keyword := url.Values{}
	keyword.Set("keyword", searchTerm)
	wq := url.Values{}
	wq.Set("wq", searchTerm)

	request, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf(searchUrlFormat, keyword.Encode(), wq.Encode()), nil)
	if err != nil {
		return nil, err
	}

	// 用不用cookie，查询结果还不一样
	header.Set("cookie", viper.GetString("jd.cookie"))
	request.Header = header
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Println("resp.Status:", resp.Status)

	// 注意！！！这里放开，会影响下面的解析
	//respBody, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, err
	//}
	//log.Println("respBody:", string(respBody))

	return parseJDHtml(resp.Body)
}

func parseJDHtml(r io.Reader) ([]*model.Product, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var goodsList *goquery.Selection
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("id")
		if id == "J_goodsList" {
			goodsList = s
			return
		}
	})

	var items []*goquery.Selection
	goodsList.Find("li").Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")
		if class == "gl-item" {
			items = append(items, s)
		}
	})

	var name, vendor, url string
	var price float64

	products := make([]*model.Product, 0, len(items))
	for _, item := range items {
		item.Find("div").Each(func(i int, s *goquery.Selection) {
			class, _ := s.Attr("class")

			if strings.Contains(class, "price") {
				priceStr := s.Find("i").Text()
				price, _ = strconv.ParseFloat(priceStr, 64)
				//fmt.Println("price:", price)
			} else if strings.Contains(class, "name") {
				text := s.Find("em").Text()
				tag := regexp.MustCompile("<[^>]+>")
				name = tag.ReplaceAllString(text, "")
				//fmt.Println("text:", name)
				//
				url, _ = s.Find("a").Attr("href")
				//fmt.Println("href:", url)

			} else if strings.Contains(class, "shop") {
				vendor = s.Find("a").Text()
				//fmt.Println("shop:", vendor)
			}
		})
		products = append(products, newJDProduct(name, price, vendor, url))
	}

	return products, nil
}

func newJDProduct(name string, price float64,
	vendor, url string) *model.Product {
	return &model.Product{
		Name:        name,
		Price:       price,
		Vendor:      vendor,
		OriginalURL: url,
		Provider:    model.JD,
	}
}
