package ClassGJ

import (
	"fmt"
	"github.com/imroc/req/v3"
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
	client := req.C()
	_, _ = client.R().
		SetHeaders(map[string]string{
			"Imprint":         imprint,
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF XWEB/6500",
			"Content-Type":    "application/json",
			"Accept-Language": "zh-CN,zh",
		}).
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
func GetWorkInfo(imprint string, member_id string) (string, string, string) {
	var response WorkInfo
	client := req.C()
	_, _ = client.R().
		SetHeaders(map[string]string{
			"Imprint":         imprint,
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF XWEB/6500",
			"Content-Type":    "application/json",
			"Accept-Language": "zh-CN,zh",
		}).
		SetPathParams(map[string]string{
			"members": member_id,
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

func GetStudentWorkInfo(imprint string, workid string) (string, string, string, string) {
	var response StudentWorkInfo
	client := req.C()
	_, _ = client.R().
		SetHeaders(map[string]string{
			"Imprint":         imprint,
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF XWEB/6500",
			"Content-Type":    "application/json",
			"Accept-Language": "zh-CN,zh",
		}).
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
	itemdetails2 := response.Data.Notify.Invest.Subject[0].ItemDetails[1].ID
	fmt.Println("作业标识："+investid, subjectid, itemdetails1, itemdetails2)
	return investid, subjectid, itemdetails1, itemdetails2
}
