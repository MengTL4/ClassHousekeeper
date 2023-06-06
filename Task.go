package main

import (
	"fmt"
)

// Task 个人作业
func Task() {
	Download()
	//获取member_id
	memberId := GetMemberId(imprint)
	//获取新一期大学习名称，作业id，创建日期
	title, workid, createat := GetWorkInfo(memberId)
	//获取作业的各项标识
	investid, subjectid, itemdetails1 := GetStudentWorkInfo(workid)
	//生成图片随机标识
	pngkey := GenerateRandomString(4)
	//获取图片认证信息
	XCosSecurityToken, Authorization := GetPNGInfo()
	//获取上传的图片路径
	pngfilepath := SubmitPNG(pngkey, XCosSecurityToken, Authorization)
	//提交作业
	result := SubmitDataStudent(title, workid, createat, pngfilepath, memberId, investid, subjectid, itemdetails1)
	fmt.Println(result)
	_, err := fmt.Scanln()
	if err != nil {
		return
	}

}
