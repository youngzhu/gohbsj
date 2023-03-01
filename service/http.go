package service

import (
	"fmt"
	"net/http"
	"sync"
	"youngzy.com/gohbsj/logging"
)

type pipelineAdaptor struct {
	RequestPipeline
}

func (p pipelineAdaptor) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	p.ProcessRequest(request, writer)
}

func Serve(pl RequestPipeline, logger logging.Logger) *sync.WaitGroup {
	wg := sync.WaitGroup{}

	adaptor := pipelineAdaptor{RequestPipeline: pl}

	httpPort := 8000
	logger.Debugf("Starting HTTP server on port %v", httpPort)
	wg.Add(1)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%v", httpPort), adaptor)
		if err != nil {
			panic(err)
		}
	}()

	return &wg
}
