package video

import (
	"Aurora-go/api"
	"Aurora-go/model"
	"encoding/json"
	"fmt"
)

func main() {

	client := api.NewAPIClient(api.NewConfiguration(api.WithBasePath("http://inner-aws-hub-test.visionular.com"), api.WithBasicAuth("83D7FBC66F284F66B6EE34A42153DF3B", "F3B954FDA1FB4A74B1B8A32BDDB3AA60")))

	//========== AddTemplate ==========

	request := model.AddTemplateApiRequest{
		TemplateName: "i_test_create_template",
		Format:       "mp4",
		Quality:      0,
		Framerate:    0,
		Resolution:   "",
		VCodec:       "",
		ACodec:       "",
		AudioBitrate: 0,
	}
	res, err := client.VideoViewsApi.AddTemplate(request)
	if err != nil {
		fmt.Println("AddTemplate fail", err)
		return
	}
	fmt.Println("AddTemplate success", res)

	//========== QueryTemplate ==========

	resQuery, err := client.VideoViewsApi.QueryTemplate("Template1647248420")
	if err != nil {
		fmt.Println("QueryTemplate fail", err)
		return
	}
	fmt.Println("QueryTemplate success", resQuery)
	rq, _ := json.Marshal(resQuery)
	fmt.Println("QueryTemplate success", string(rq))

	//========== UpdateTemplate ==========

	requestUpdate := model.TemplateApiRequest{
		TemplateName:    "Template1647248420",
		Format:          "mp4",
		Description:     "hhhhhhhhhhh",
		Acodec:          "aac",
		AudioBitrate:    0,
		AudioChannels:   0,
		AudioSampleRate: "48k",
		FrameRate:       25,
		GopSize:         50,
		LogoOffset:      "x=(W-w)/2:y=(H-h)/2",
		LogoPath:        "https://cloudhub.oss-cn-zhangjiakou.aliyuncs.com/eg/logo.png",
		LogoSize:        "50x50",
		Quality:         6,
		Resolution:      "0x0",
		SegTime:         10,
		SegType:         "fmp4",
		Vcodec:          "h265",
		VideoBitrate:    3000,
	}

	resUpdate, err := client.VideoViewsApi.UpdateTemplate(requestUpdate)
	if err != nil {
		fmt.Println("UpdateTemplate fail", err)
		return
	}
	fmt.Println("UpdateTemplate success", resUpdate)

	//========== DeleteTemplate ==========

	resDel, err := client.VideoViewsApi.DeleteTemplate("i_test_create_template")
	if err != nil {
		fmt.Println("DeleteTemplate fail", err)
		return
	}
	rd, _ := json.Marshal(resDel)
	fmt.Println("DeleteTemplate success", string(rd))

	//========== Add Storage ==========

	storageReq := model.StorageApiRequest{
		Type:      "oss",
		Region:    "cn-beijing",
		Bucket:    "mytestbucket_sdk",
		Prefix:    "",
		Notify:    "http://test.com/callback",
		StorageAk: "aaaaaaaaaaaaaaaaaaaaa",
		StorageSk: "ssssssssssssssssssssssssssssssssssssssssss",
	}
	resAdd, err := client.VideoViewsApi.AddStorage(storageReq)
	if err != nil {
		fmt.Println("AddStorage fail", err)
		return
	}
	ra, _ := json.Marshal(resAdd)
	fmt.Println("AddStorage success", string(ra))

	//========== ListStorage ==========

	resList, err := client.VideoViewsApi.ListStorage()
	if err != nil {
		fmt.Println("ListStorage fail", err)
		return
	}
	rl, _ := json.Marshal(resList)
	fmt.Println("ListStorage success", string(rl))

	//========== DelStorage ==========

	resDs, err := client.VideoViewsApi.DeleteStorage("f5da9ed069ee201d")
	if err != nil {
		fmt.Println("DeleteStorage fail", err)
		return
	}
	rds, _ := json.Marshal(resDs)
	fmt.Println("DeleteStorage success", string(rds))
}
