package main

import (
	"aurora-sdk-go/api"
	"aurora-sdk-go/model"
	"fmt"
)

func main() {

	client := api.NewAPIClient(api.NewConfiguration(api.WithBasePath("http://inner-aws-hub-test.visionular.com"), api.WithBasicAuth("83D7FBC66F284F66B6EE34A42153DF3B", "F3B954FDA1FB4A74B1B8A32BDDB3AA60")))

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
	resdis, err := client.LiveStreamsApi.DisableLiveStream("NTBmNWM0NGY4NTM2OTcyNjlmZjZiYTcxZDY1NzJhNDc")
	if err != nil {
		fmt.Println("DisableLiveStream fail", err)
		return
	}

	fmt.Println("DisableLiveStream success", resdis)

	// ========== enable-live-stream ==========
	resen, err := client.LiveStreamsApi.EnableLiveStream("NTBmNWM0NGY4NTM2OTcyNjlmZjZiYTcxZDY1NzJhNDc")
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
	resinfo, err := client.LiveStreamsApi.QueryLiveStreamInfo("NTBmNWM0NGY4NTM2OTcyNjlmZjZiYTcxZDY1NzJhNDc")
	if err != nil {
		fmt.Println("GetLiveStreamInfo fail", err)
		return
	}

	fmt.Println("GetLiveStreamInfo success", resinfo)
}
