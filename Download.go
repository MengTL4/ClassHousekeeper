package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
	"io"
	"log"
	"net/http"
	"regexp"
)

func Download() {
	// 发送HTTP请求并获取HTML内容
	response, err := http.Get("https://news.cyol.com/gb/channels/vrGlAKDl/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	// 使用GoQuery解析HTML
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 声明一个切片来存储href
	var hrefs []string
	// 查找所有<h3>标签内的<a>标签，并提取href属性
	doc.Find("h3 a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			hrefs = append(hrefs, href)
		}
	})
	regex := regexp.MustCompile(`\/([^/]+)\/m.html$`)
	match := regex.FindStringSubmatch(hrefs[0])
	client := req.C()
	get, err := client.R().SetOutputFile("test.png").
		SetPathParams(map[string]string{
			"id": match[1],
		}).
		Get("https://h5.cyol.com/special/daxuexi/{id}/images/end.jpg")
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(get.Body)
}
