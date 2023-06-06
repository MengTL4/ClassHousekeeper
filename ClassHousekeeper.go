package main

import (
	"fmt"
	"github.com/imroc/req/v3"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

const imprint = "oWRkU0aXAjlcP9IeY9lKaHk0XltI"

var (
	headers = map[string]string{
		"Imprint":         imprint,
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF XWEB/6500",
		"Content-Type":    "application/json",
		"Accept-Language": "zh-CN,zh",
	}
	client = req.C()
)

type UserInfo struct {
	CurrentUser struct {
		ChildClassList []struct {
			MemberID string `json:"member_id"`
		} `json:"child_class_list"`
	} `json:"currentUser"`
}

// GetMemberId 通过Imprint获取member_id
func GetMemberId(imprint string) string {
	var response UserInfo
	client = req.C()
	_, _ = client.R().
		SetHeaders(headers).
		SetBodyJsonMarshal(map[string]string{
			"openid": imprint,
		}).
		SetSuccessResult(&response).
		Post("https://a.welife001.com/getUser")
	return response.CurrentUser.ChildClassList[0].MemberID
}

// WorkInfo 作业的相关信息
type WorkInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		//作业的id
		ID string `json:"_id"`
		//作业的创建时间
		CreateAt string `json:"create_at"`
		Title    string `json:"title"`
	} `json:"data"`
}

// GetWorkInfo 获取当前所有的作业，并筛选出最新一期作业
func GetWorkInfo(memberId string) (string, string, string) {
	var response WorkInfo
	client = req.C()
	_, _ = client.R().
		SetHeaders(headers).
		SetPathParams(map[string]string{
			"members": memberId,
		}).
		SetSuccessResult(&response).
		Get("https://a.welife001.com/info/getParent?onlyMe=false&lookAll=true&page=0&size=10&type=-1&date=-1&members={members}")
	return response.Data[0].Title, response.Data[0].ID, response.Data[0].CreateAt
}

// StudentWorkInfo 通过上一个函数获取的最新一期作业的id来获取更多作业信息
type StudentWorkInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Notify struct {
			Invest struct {
				ID      string `json:"_id"`
				Subject []struct {
					ID          string `json:"_id"`
					ItemDetails []struct {
						ID string `json:"_id"`
					} `json:"item_details"`
				} `json:"subject"`
			} `json:"invest"`
		} `json:"notify"`
	} `json:"data"`
}

func GetStudentWorkInfo(workid string) (string, string, string) {
	var response StudentWorkInfo
	client = req.C()
	_, _ = client.R().
		SetHeaders(headers).
		SetBodyJsonMarshal(map[string]string{
			"_id":  workid,
			"page": "0",
			"size": "10",
		}).
		SetSuccessResult(&response).
		Post("https://a.welife001.com/applet/notify/checkNew2Parent")
	investid := response.Data.Notify.Invest.ID
	subjectid := response.Data.Notify.Invest.Subject[0].ID
	itemdetails1 := response.Data.Notify.Invest.Subject[0].ItemDetails[0].ID
	//itemdetails2 := response.Data.Notify.Invest.Subject[0].ItemDetails[1].ID
	fmt.Println("作业标识："+investid, subjectid, itemdetails1)
	return investid, subjectid, itemdetails1
}

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
func GetPNGInfo() (string, string) {
	var response PNGInfo
	client = req.C()
	body := map[string]bool{"stsGray": true}
	_, _ = client.R().
		SetHeaders(headers).
		SetBody(body).
		SetSuccessResult(&response).
		Post("https://a.welife001.com/applet/api/appGetCosSTS")
	return response.Data.XCosSecurityToken, response.Data.Authorization
}

func SubmitPNG(key string, XCosSecurityToken string, Authorization string) string {
	client = req.C()
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

func SubmitDataStudent(title string, workid string, createAt string, filepath string, memberId string, investid string, subjectid string, itemDetailsid1 string) string {
	timeunix := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	requestbody := map[string]interface{}{
		"extra":       1,
		"id":          workid,
		"daka_day":    "",
		"submit_type": "submit",
		"networkType": "wifi",
		"member_id":   memberId,
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
			"create_at": createAt,
			"update_at": createAt,
			"__v":       0,
			"time":      timeunix,
		},
		"feedback_text": "",
	}
	var response Code
	client = req.C()
	_, _ = client.R().
		SetBodyJsonMarshal(requestbody).
		SetHeaders(headers).
		SetSuccessResult(&response).
		Post("https://a.welife001.com/applet/notify/feedbackWithOss")
	if response.Code == 0 {
		return "成功完成" + title
	} else {
		return "提交失败"
	}
}
