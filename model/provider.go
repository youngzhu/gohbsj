package model

// Provider Information Provider 信息供应商
type Provider struct {
	ProviderID   string
	ProviderName string
}

var (
	JD     = &Provider{ProviderID: "jd", ProviderName: "京东"}
	TaoBao = &Provider{ProviderID: "taobao", ProviderName: "淘宝"}
)
