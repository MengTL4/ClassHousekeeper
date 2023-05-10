package main

import (
	"DaXueXiAuto/ClassGJ"
	"DaXueXiAuto/JiangSuGTT"
)

func main() {
	//读取配置文件
	cookie, imprint, _ := ReadConfig()
	//完成新一期大学习，并下载截图
	lessonid := JiangSuGTT.GetLesson(cookie)
	//获取截图
	pngid := JiangSuGTT.DoLesson(lessonid, cookie)
	//下载图片
	JiangSuGTT.Download(pngid)
	//获取member_id
	member_id := ClassGJ.GetMemberId(imprint)
	//获取新一期大学习名称，作业id，创建日期
	title, workid, createat := ClassGJ.GetWorkInfo(imprint, member_id)
	//获取作业的各项标识
	investid, subjectid, itemdetails1, itemdetails2 := ClassGJ.GetStudentWorkInfo(imprint, workid)
	//生成图片随机标识
	pngkey := ClassGJ.GenerateRandomString(4)
	//获取图片认证信息
	XCosSecurityToken, Authorization := ClassGJ.GetPNGInfo(imprint)
	//获取上传的图片路径
	pngfilepath := ClassGJ.SubmitPNG(pngkey, XCosSecurityToken, Authorization)
	//提交作业
	ClassGJ.SubmitDataStudent(title, workid, createat, pngfilepath, member_id, imprint, investid, subjectid, itemdetails1, itemdetails2)
	fmt.Scanln()
}
