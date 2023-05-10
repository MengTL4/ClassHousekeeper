package JiangSuGTT

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"io"
	"regexp"
	"strconv"
)

// GetLessonResponse  getlesson接口响应的json数据结构体
type GetLessonResponse struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Redirect string `json:"redirect"`
	Data     []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		ShowConfirm bool   `json:"show_confirm"`
		Cover       string `json:"cover"`
	} `json:"data"`
}

func GetLesson(cookie string) string {
	var response GetLessonResponse
	client := req.C()
	lessondata, _ := client.R().
		SetHeaders(map[string]string{
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x6308011a) XWEB/6500",
			"Content-Type":    "application/json",
			"Origin":          "https://service.jiangsugqt.org",
			"Referer":         "https://service.jiangsugqt.org/youth-h5/",
			"Accept-Language": "zh-CN,zh",
			"Connection":      "close",
			"Cookie":          cookie,
		}).
		SetBodyJsonMarshal(map[string]string{
			"page":  "1",
			"limit": "5",
		}).
		SetSuccessResult(&response).
		Post("https://service.jiangsugqt.org/api/lessons")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("关闭连接失败")
		}
	}(lessondata.Body)
	return strconv.Itoa(response.Data[0].ID)
}

// DoLessonResponse  dolesson接口响应的json数据
type DoLessonResponse struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Redirect string `json:"redirect"`
	Data     struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	} `json:"data"`
}

func DoLesson(lessonId string, cookie string) string {
	var response DoLessonResponse
	client := req.C()
	lessondata, _ := client.R().
		SetHeaders(map[string]string{
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x6308011a) XWEB/6500",
			"Content-Type":    "application/json",
			"Origin":          "https://service.jiangsugqt.org",
			"Referer":         "https://service.jiangsugqt.org/youth-h5/",
			"Accept-Language": "zh-CN,zh",
			"Connection":      "close",
			"Cookie":          cookie,
		}).
		SetBodyJsonMarshal(map[string]string{
			"lesson_id": lessonId,
		}).
		Post("https://service.jiangsugqt.org//api/doLesson")
	lessondataString := lessondata.String()
	err := json.Unmarshal([]byte(lessondataString), &response)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	pattern := "daxuexi/(.*?)/index.html"
	re, _ := regexp.Compile(pattern)
	match := re.FindStringSubmatch(response.Data.URL)[1]
	return match
}
