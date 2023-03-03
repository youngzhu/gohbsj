package model

// Product 商品信息
type Product struct {
	Name        string  // 产品名称
	Price       float64 // 价格
	Vendor      string  // 卖家
	OriginalURL string  // 原始网址（可点击查看详情）
	*Provider           // 信息供应商（来源）
}
