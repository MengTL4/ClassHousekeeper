package JiangSuGTT

import "github.com/imroc/req/v3"

func Download(id string) {
	client := req.C()
	client.R().SetOutputFile("test.png").
		SetPathParams(map[string]string{
			"id": id,
		}).
		Get("https://h5.cyol.com/special/daxuexi/{id}/images/end.jpg")
}
