package search

import (
	"log"
	"sync"
	"time"
	"youngzy.com/gohbsj/model"
)

type Searcher struct {
	SearchTerm string

	products []model.Product

	Cost    time.Duration // 耗时
	Records int           // 搜索到的记录条数
	Hit     bool          // 是否进行了搜索
}

func (s *Searcher) Run() {
	start := time.Now()

	prodChan := make(chan *model.Product)

	var wg sync.WaitGroup
	wg.Add(len(matchers))

	for _, m := range matchers {
		go func(matcher Matcher) {
			match(s.SearchTerm, matcher, prodChan)
			wg.Done()
		}(m)
	}

	go func() {
		wg.Wait()

		close(prodChan)
	}()

	// 每次搜索，结果需要重置
	var searchResult []model.Product
	for p := range prodChan {
		searchResult = append(searchResult, *p)
	}
	// 按价格排序
	sortProductByPrice(searchResult)
	s.products = searchResult

	s.Cost = time.Since(start)
	s.Records = len(s.products)
	s.Hit = true
}

var matchers = make(map[string]Matcher)

func Register(matcherID string, matcher Matcher) {
	log.Println("Register", matcherID, "matcher")
	matchers[matcherID] = matcher
}

func (s *Searcher) GetProviders() (providers []model.Provider) {
	aMap := map[string]model.Provider{}
	for _, p := range s.products {
		if _, ok := aMap[p.ProviderID]; !ok {
			aMap[p.ProviderID] = *p.Provider
			providers = append(providers, *p.Provider)
		}
	}
	log.Println("providers:", len(providers))
	return
}

func (s *Searcher) GetProducts(providerID string) (products []model.Product) {
	if providerID == "" {
		return s.products
	}

	for _, product := range s.products {
		if product.ProviderID == providerID {
			products = append(products, product)
		}
	}
	return
}
