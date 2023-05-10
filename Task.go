package main

import (
	"DaXueXiAuto/ClassGJ"
	"DaXueXiAuto/JiangSuGTT"
	"fmt"
	"sync"
)

// Task 个人作业
func Task() {
	//读取配置文件
	cookie, imprint, identity := ReadConfig()
	//完成新一期大学习，并下载截图
	lessonid := JiangSuGTT.GetLesson(cookie)
	//获取截图
	pngid := JiangSuGTT.DoLesson(lessonid, cookie)
	//下载图片
	JiangSuGTT.Download(pngid)
	if identity == "0" || identity == "" {
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
		result := ClassGJ.SubmitDataStudent(title, workid, createat, pngfilepath, member_id, imprint, investid, subjectid, itemdetails1, itemdetails2)
		fmt.Println(result)
		fmt.Scanln()
	} else {
		//获取member_id
		member_id := ClassGJ.GetMemberId(imprint)
		//获取新一期大学习名称，作业id，创建日期
		title, workid, _ := ClassGJ.GetWorkInfo(imprint, member_id)
		//生成图片随机标识
		pngkey := ClassGJ.GenerateRandomString(4)
		//获取图片认证信息
		XCosSecurityToken, Authorization := ClassGJ.GetPNGInfo(imprint)
		//获取上传的图片路径
		pngfilepath := ClassGJ.SubmitPNG(pngkey, XCosSecurityToken, Authorization)
		//提交作业
		ClassGJ.SpecialSubmit(title, workid, pngfilepath, member_id, imprint)
		fmt.Scanln()
	}
}

// TaskAll 全部作业
func TaskAll() {
	//读取配置文件
	cookie, imprint, identity := ReadConfig()
	//完成新一期大学习，并下载截图
	lessonid := JiangSuGTT.GetLesson(cookie)
	//获取截图
	pngid := JiangSuGTT.DoLesson(lessonid, cookie)
	//下载图片
	JiangSuGTT.Download(pngid)
	if identity == "0" || identity == "" {
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
		//=======================下面是提交一个班所有的作业，暂时只能自行适配(修改AllSubmit函数)=======================
		//由于请求过多，如果从头开始使用并发，会消耗过度资源，使用任何一份正确参数均可提交，所有不采用从开始的并发操作
		wxopenids := ClassGJ.AllSubmit(imprint)
		var wg sync.WaitGroup
		for _, wxopenid := range wxopenids {
			wg.Add(1)
			go func(wxopenid string) {
				defer wg.Done()
				ClassGJ.SubmitDataStudent(title, workid, createat, pngfilepath, member_id, wxopenid, investid, subjectid, itemdetails1, itemdetails2)
			}(wxopenid)
		}
		go func() {
			wg.Wait()
		}()
		fmt.Println("全部作业已完成")
		//==================================================================================
		fmt.Scanln()
	} else {
		//获取member_id
		member_id := ClassGJ.GetMemberId(imprint)
		//获取新一期大学习名称，作业id，创建日期
		title, workid, _ := ClassGJ.GetWorkInfo(imprint, member_id)
		//生成图片随机标识
		pngkey := ClassGJ.GenerateRandomString(4)
		//获取图片认证信息
		XCosSecurityToken, Authorization := ClassGJ.GetPNGInfo(imprint)
		//获取上传的图片路径
		pngfilepath := ClassGJ.SubmitPNG(pngkey, XCosSecurityToken, Authorization)
		//=======================下面是提交一个班所有的作业，暂时只能自行适配(修改AllSubmit函数)=======================
		//由于请求过多，如果从头开始使用并发，会消耗过度资源，使用任何一份正确参数均可提交，所有不采用从开始的并发操作
		wxopenids := ClassGJ.AllSubmit(imprint)
		var wg sync.WaitGroup
		for _, wxopenid := range wxopenids {
			wg.Add(1)
			go func(wxopenid string) {
				defer wg.Done()
				ClassGJ.SpecialSubmit(title, workid, pngfilepath, member_id, wxopenid)
			}(wxopenid)
		}
		go func() {
			wg.Wait()
		}()
		fmt.Println("全部作业已完成")
		//==================================================================================
		fmt.Scanln()
	}
}
