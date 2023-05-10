package ClassGJ

import (
	"github.com/imroc/req/v3"
)

type WXOpenIdInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Feedbacks []struct {
			WxOpenid string `json:"wx_openid"`
		} `json:"feedbacks"`
	} `json:"data"`
}

// AllSubmit 用来提交一个班作业的函数
func AllSubmit(imprint string) []string {
	var response WXOpenIdInfo
	client := req.C()
	_, _ = client.R().
		SetHeaders(map[string]string{
			"Imprint":         imprint,
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF XWEB/6500",
			"Content-Type":    "application/json",
			"Accept-Language": "zh-CN,zh",
		}).
		SetBodyJsonMarshal(map[string]string{
			"_id":  "6421027c7cc4b507438663ea",
			"page": "0",
			"size": "10",
		}).
		SetSuccessResult(&response).
		Post("https://a.welife001.com/applet/notify/checkNew2Parent")
	// 创建一个通道来保存结果
	resultChan := make(chan string)
	go func() {
		// 异步执行for循环
		for i := 0; i <= 42; i++ {
			resultChan <- response.Data.Feedbacks[i].WxOpenid
		}
		close(resultChan)
	}()
	// 从通道中读取结果并保存到数组
	var results []string
	for result := range resultChan {
		results = append(results, result)
	}
	return results
}
