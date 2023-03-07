package matcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// 这是ChatGPT给出的代码

func main() {
	// 设置请求参数
	url := "https://api.jd.com/routerjson"
	appKey := "your_app_key"
	//appSecret := "your_app_secret"
	method := "jd.union.search.goods.query"
	format := "json"
	page := 1
	pageSize := 20
	sortName := "price"
	sort := "asc"
	keyword := "iPhone"

	// 构建请求URL
	reqURL := fmt.Sprintf("%s?app_key=%s&v=2.0&method=%s&format=%s&page=%d&page_size=%d&sort_name=%s&sort=%s&keyword=%s&sign_method=md5&timestamp=%s",
		url, appKey, method, format, page, pageSize, sortName, sort, keyword, timestamp())

	// 发送HTTP GET请求
	resp, err := http.Get(reqURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 解析JSON响应
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印结果
	fmt.Println(result)
}

// 获取当前时间戳字符串
func timestamp() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}
