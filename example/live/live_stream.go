package main

import (
	"fmt"
	"github.com/visionular/aurora-sdk-go/api"
	"github.com/visionular/aurora-sdk-go/model"
)

func main() {

	client := api.NewAPIClient(api.NewConfiguration(api.WithBasePath("https://aaa.xxx.com"), api.WithBasicAuth("AAAAAAAAAAAAAAAAAAAAAAA", "BBBBBBBBBBBBBBBBBBBBBBB")))

	////========== create-live-stream ==========
	var targets []*model.SimulcastTargets
	request := model.CreateStreamRequest{
		Type:              "",
		ReconnectWindow:   nil,
		UserMetadata:      "userMetadata",
		SimulcastTargets:  targets,
		Policy:            "",
		Record:            false,
		WatermarkSettings: nil,
	}
	res, err := client.LiveStreamsApi.CreateLiveStream(request)
	if err != nil {
		fmt.Println("CreateLiveStream fail", err)
		return
	}
	fmt.Println("CreateLiveStream success", res)

	// ========== disable-live-stream ==========
	resdis, err := client.LiveStreamsApi.DisableLiveStream("aaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		fmt.Println("DisableLiveStream fail", err)
		return
	}

	fmt.Println("DisableLiveStream success", resdis)

	// ========== enable-live-stream ==========
	resen, err := client.LiveStreamsApi.EnableLiveStream("aaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		fmt.Println("EnableLiveStream fail", err)
		return
	}

	fmt.Println("EnableLiveStream success", resen)

	// ========== list-live-stream ==========
	reslist, err := client.LiveStreamsApi.ListLiveStreams()
	if err != nil {
		fmt.Println("ListLiveStreams fail", err)
		return
	}

	fmt.Println("ListLiveStreams success", reslist)

	// ========== get-live-stream-info ==========
	resinfo, err := client.LiveStreamsApi.QueryLiveStreamInfo("aaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		fmt.Println("GetLiveStreamInfo fail", err)
		return
	}

	fmt.Println("GetLiveStreamInfo success", resinfo)
}
