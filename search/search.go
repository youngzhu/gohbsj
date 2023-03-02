package search

import (
	"log"
	"sync"
	"youngzy.com/gohbsj/model"
)

type Searcher struct {
	result []model.Product
}

func (s *Searcher) Run(searchTerm string) {
	prodChan := make(chan *model.Product)

	var wg sync.WaitGroup
	wg.Add(len(matchers))

	for _, m := range matchers {
		go func(matcher Matcher) {
			match(searchTerm, matcher, prodChan)
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

func (s *Searcher) GetProviders() (providers []string) {
	aMap := map[string]string{}
	for _, p := range s.result {
		if _, ok := aMap[p.Provider]; !ok {
			aMap[p.Provider] = p.Provider
			providers = append(providers, p.Provider)
		}
	}
	log.Println("providers:", len(providers))
	return
}

func (s *Searcher) GetProducts() []model.Product {
	return s.result
}
