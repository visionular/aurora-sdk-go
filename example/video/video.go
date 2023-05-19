package video

import (
	"encoding/json"
	"fmt"
	"github.com/visionular/aurora-sdk-go/api"
	"github.com/visionular/aurora-sdk-go/model"
)

func main() {

	client := api.NewAPIClient(api.NewConfiguration(api.WithBasePath("https://aaa.xxx.com"), api.WithBasicAuth("AAAAAAAAAAAAAAAAAAAAAAA", "BBBBBBBBBBBBBBBBBBBBBBB")))

	//========== QueryTemplate ==========

	resQuery, err := client.VideoViewsApi.QueryTemplate("aaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		fmt.Println("QueryTemplate fail", err)
		return
	}
	fmt.Println("QueryTemplate success", resQuery)
	rq, _ := json.Marshal(resQuery)
	fmt.Println("QueryTemplate success", string(rq))

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

	resDs, err := client.VideoViewsApi.DeleteStorage("aaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		fmt.Println("DeleteStorage fail", err)
		return
	}
	rds, _ := json.Marshal(resDs)
	fmt.Println("DeleteStorage success", string(rds))
}
