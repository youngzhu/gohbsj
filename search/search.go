package search

import (
	"log"
	"sync"
	"youngzy.com/gohbsj/model"
)

type Searcher struct {
	SearchTerm string

	result []model.Product
}

func (s *Searcher) Run() {
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

	for p := range prodChan {
		s.result = append(s.result, *p)
	}
}

var matchers = make(map[string]Matcher)

func Register(matcherID string, matcher Matcher) {
	log.Println("Register", matcherID, "matcher")
	matchers[matcherID] = matcher
}

func (s *Searcher) GetProviders() (providers []model.Provider) {
	aMap := map[string]model.Provider{}
	for _, p := range s.result {
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
		return s.result
	}

	for _, product := range s.result {
		if product.ProviderID == providerID {
			products = append(products, product)
		}
	}
	return
}
