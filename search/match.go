package search

import (
	"log"
	"youngzy.com/gohbsj/model"
)

type Matcher interface {
	Search(keywords string) ([]*model.Product, error)
}

func match(searchTerm string,
	matcher Matcher, ch chan<- *model.Product) {
	results, err := matcher.Search(searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range results {
		ch <- item
	}
}
