package main

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
	"youngzy.com/gohbsj/logging"
	"youngzy.com/gohbsj/search"
	_ "youngzy.com/gohbsj/search/matcher"
	"youngzy.com/gohbsj/service"
	"youngzy.com/gohbsj/service/basic"
	"youngzy.com/gohbsj/template"
)

func registerServices() {
	var err error

	err = service.AddSingleton(func() logging.Logger {
		return logging.NewDefaultLogger()
	})
	if err != nil {
		panic(err)
	}

	err = service.AddSingleton(
		func() template.TemplateExecutor {
			err = template.LoadTemplates()
			return &template.LayoutTemplateProcessor{}
		})
	if err != nil {
		panic(err)
	}

	err = service.AddSingleton(func() *search.Searcher {
		return &search.Searcher{}
	})
	if err != nil {
		panic(err)
	}
}

func createPipeline() service.RequestPipeline {
	return service.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		service.NewRouter(
			service.HandlerEntry{"", service.ProductHandler{}},
			service.HandlerEntry{"", service.ProviderHandler{}},
		).AddMethodAlias("/", service.ProductHandler.GetProducts, "").
			AddMethodAlias("/search", service.ProductHandler.PostSearch),
		//.AddMethodAlias("/products[/]?[A-z0-9]*?",
		//	handler.ProductHandler.GetProducts, 0, 1),
	)
}

func initConfig() {
	viper.SetConfigName(".gohbsj") // 不要后缀
	viper.AddConfigPath("$HOME")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	//log.Println("main===" + viper.GetString("test") + "===")
}

func main() {
	initConfig()
	registerServices()

	results, err := service.Call(service.Serve, createPipeline())
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
