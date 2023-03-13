package matcher

import (
	"net/http"
	"reflect"
	"testing"
)

// key值得大小写有影响？
// 从jd的调用情况来看，是的。new就可以正常返回，Set不行
func TestRequestHeader(t *testing.T) {
	// new Header
	newHeader := http.Header{
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
	req1, _ := http.NewRequest("", "", nil)
	req1.Header = newHeader

	// Header.Set()
	req2, _ := http.NewRequest("", "", nil)
	req2.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req2.Header.Set("accept-encoding", "gzip, deflate, br")
	req2.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req2.Header.Set("referer", "https://www.jd.com/")
	req2.Header.Set("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Microsoft Edge\";v=\"110\"")
	req2.Header.Set("sec-ch-ua-mobile", "?0")
	req2.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req2.Header.Set("sec-fetch-dest", "document")
	req2.Header.Set("sec-fetch-mode", "navigate")
	req2.Header.Set("sec-fetch-site", "same-site")
	req2.Header.Set("sec-fetch-user", "?1")
	req2.Header.Set("upgrade-insecure-requests", "1")
	req2.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.63")

	// Header.Add()
	req3, _ := http.NewRequest("", "", nil)
	req3.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req3.Header.Add("accept-encoding", "gzip, deflate, br")
	req3.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req3.Header.Add("referer", "https://www.jd.com/")
	req3.Header.Add("sec-ch-ua", "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Microsoft Edge\";v=\"110\"")
	req3.Header.Add("sec-ch-ua-mobile", "?0")
	req3.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req3.Header.Add("sec-fetch-dest", "document")
	req3.Header.Add("sec-fetch-mode", "navigate")
	req3.Header.Add("sec-fetch-site", "same-site")
	req3.Header.Add("sec-fetch-user", "?1")
	req3.Header.Add("upgrade-insecure-requests", "1")
	req3.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.63")

	t.Logf("new:\n%s", req1.Header)
	t.Logf("Set:\n%s", req2.Header)
	t.Logf("Add:\n%s", req3.Header)

	if !reflect.DeepEqual(req1.Header, req2.Header) {
		t.Log("'new' is different from 'Set'")
	} else {
		t.Log("'new' is the same as 'Set'")
	}

	if !reflect.DeepEqual(req3.Header, req2.Header) {
		t.Log("'Add' is different from 'Set'")
	} else {
		t.Log("'Add' is the same as 'Set'")
	}
}
