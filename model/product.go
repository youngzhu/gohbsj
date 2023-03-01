package model

type Product struct {
	Name        string  // 产品名称
	Price       float64 // 价格
	Vendor      string  // 卖家
	Provider    string  // 信息供应商（来源）
	OriginalURL string  // 原始网址（可点击查看详情）
}

var TestProducts = []Product{
	{Name: "iPhone", Price: 999.99, Vendor: "苹果旗舰店", Provider: "京东", OriginalURL: "www.youngz.com"},
}
