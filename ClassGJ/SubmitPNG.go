package ClassGJ

import (
	"fmt"
	"github.com/imroc/req/v3"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type PNGInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ExpiredTime       int    `json:"expiredTime"`
		RequestID         string `json:"requestId"`
		XCosSecurityToken string `json:"XCosSecurityToken"`
		Authorization     string `json:"Authorization"`
	} `json:"data"`
}

// GenerateRandomString 图片标识随机化，保证图片唯一性
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.NewSource(time.Now().UnixNano())
	result := ""
	for i := 0; i < length; i++ {
		index := rand.Intn(len(charset))
		result += string(charset[index])
	}
	key := "11oWRkU0aXAjlcP9IeY9lKaHk0XltI_img/" + result + "_cos@513.png"
	return key
}

// GetPNGInfo 获取上传图片前需要的必须凭证
func GetPNGInfo(imprint string) (string, string) {
	var response PNGInfo
	client := req.C()
	body := map[string]bool{"stsGray": true}
	_, _ = client.R().
		SetHeaders(map[string]string{
			"Imprint":         imprint,
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF XWEB/6500",
			"Content-Type":    "application/json",
			"Accept-Language": "zh-CN,zh",
		}).
		SetBody(body).
		SetSuccessResult(&response).
		Post("https://a.welife001.com/applet/api/appGetCosSTS")
	return response.Data.XCosSecurityToken, response.Data.Authorization
}

func SubmitPNG(key string, XCosSecurityToken string, Authorization string) string {
	client := req.C()
	submit, _ := client.R().
		SetFile("file", "test.png").
		SetFormData(map[string]string{
			"key":                   key,
			"success_action_status": "200",
			"Signature":             Authorization,
			"Content-Type":          "application/json",
			"x-cos-security-token":  XCosSecurityToken,
		}).
		Post("https://img-1302562365.cos.ap-beijing.myqcloud.com")
	Location := submit.Response.Header.Get("Location")
	re := regexp.MustCompile(`http://\S+/(\S+/\S+)`)
	match := re.FindStringSubmatchIndex(Location)
	return Location[match[2]:match[3]]
}

type Code struct {
	Code int `json:"code"`
}

func SubmitDataStudent(title string, workid string, create_at string, filepath string, member_id string, imprint string, investid string, subjectid string, itemDetailsid1 string, itemDetailsid2 string) {
	timeunix := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	requestbody := map[string]interface{}{
		"extra":       1,
		"id":          workid,
		"daka_day":    "",
		"submit_type": "submit",
		"networkType": "wifi",
		"member_id":   member_id,
		"op":          "add",
		"invest": map[string]interface{}{
			"is_tmp":     true,
			"is_private": false,
			"_id":        investid,
			"subject": []map[string]interface{}{
				{
					"seq":          0,
					"cate":         5,
					"inputs_count": 0,
					"_id":          subjectid,
					"inputs":       []map[string]interface{}{},
					"item_details": []map[string]interface{}{
						{
							"seq":          0,
							"checks_count": 0,
							"rate":         0,
							"_id":          itemDetailsid1,
							"file":         []map[string]interface{}{},
							"checkedlist":  []map[string]interface{}{},
							"name":         "",
						},
						{
							"seq":          1,
							"checks_count": 0,
							"rate":         0,
							"_id":          itemDetailsid2,
							"file":         []map[string]interface{}{},
							"checkedlist":  []map[string]interface{}{},
							"name":         "",
						},
					},
					"title":    "学习完成截图",
					"required": true,
					"input": map[string]interface{}{
						"file": []map[string]interface{}{
							{
								"file":     "wxfile://temp/test.png",
								"cate":     "img",
								"new_name": filepath,
								"type":     1,
								"size":     152658,
								"uploaded": true,
								"id":       filepath,
							},
						},
					},
					"valid": true,
				},
			},
			"create_at": create_at,
			"update_at": create_at,
			"__v":       0,
			"time":      timeunix,
		},
		"feedback_text": "",
	}
	var response Code
	client := req.C()
	_, _ = client.R().
		SetBodyJsonMarshal(requestbody).
		SetHeaders(map[string]string{
			"Imprint":         imprint,
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF XWEB/6500",
			"Content-Type":    "application/json",
			"Accept-Language": "zh-CN,zh",
		}).
		SetSuccessResult(&response).
		Post("https://a.welife001.com/applet/notify/feedbackWithOss")
	if response.Code == 0 {
		fmt.Println("成功完成" + title)
	}
}
