# Aurora Go

![img.png](img.png)

Official Aurora API wrapper for golang projects, supporting both AuroraCloud Live and AuroraCloud.

[AuroraCloud Live](https://docs.visionular.com/auroralive) is a managed live streaming solution that provides high-definition and uninterrupted live video services that are quick and easy to set up, with low latency and supports a high number of concurrent viewers.

[AuroraCloud VOD](https://docs.visionular.com/auroracloud) is an audio and video media processing service based on the intelligent encoding and artificial intelligence content adaptive video processing technology of Visionular. AuroraCloud provides a scalable, reliable, easy-to-use and maintenance-free cloud service for transcoding media files.

## Installation

```
go get github.com/visionular/aurora-sdk-go
```

## Getting Started

### Overview

Aurora Go is a code generated lightweight wrapper around the Aurora REST API and reflects them accurately. This has a few consequences you should watch out for:

For almost all API responses, the object you're looking for will be in the `data` field on the API response object, as in the example below. This is because we designed our APIs with similar concepts to the [JSON:API](https://jsonapi.org/) standard. This means we'll be able to return more metadata from our API calls (such as related entities) without the need to make breaking changes to our APIs. We've decided not to hide that in this library.

### Authentication

To use the Aurora API, you'll need an access token and a secret. [documentation.](https://docs.visionular.com/auroracloud/apireference/#section/Overview/Signature-method)

Its up to you to manage your token and secret. In our examples, we read them from `AccessKey` and `SecretKey` in your environment.

### Example Usage

Below is a quick example of using Aurora-go to list the Video assets stored in your Aurora account.

Be sure to also checkout the [example directory](example/).

```go
package main

import (
	"fmt"

	"github.com/visionular/aurora-sdk-go/api"
	"github.com/visionular/aurora-sdk-go/model"
)

func main() {

	client := api.NewAPIClient(api.NewConfiguration(api.WithBasePath("https://api.visionular.com"), api.WithBasicAuth("AccessKey", "SecretKey")))

	//========== create-live-stream ==========
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
		fmt.Printf("CreateLiveStream fail:  body: %v,  err: %v: \n", res, err)
		return
	}
	fmt.Println("CreateLiveStream success", res)
}

```

### AuroraCloud Live Api List

| Interfaces          | Reference                                                                                                                                                         |
|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| CreateLiveStream    | [create-a-live-stream](https://docs.visionular.com/auroralive/apireference/#tag/Live-Stream/paths/~1live~1v1~1live-streams/post)                                  |
| DisableLiveStream   | [disabling-a-live-stream](https://docs.visionular.com/auroralive/apireference/#tag/Live-Stream/paths/~1live~1v1~1live-streams~1%7Blive_stream_id%7D~1disable/put) |
| EnableLiveStream    | [resuming-a-live-stream](https://docs.visionular.com/auroralive/apireference/#tag/Live-Stream/paths/~1live~1v1~1live-streams~1%7Blive_stream_id%7D~1enable/put)   |
| ListLiveStreams     | [list-live-streams](https://docs.visionular.com/auroralive/apireference/#tag/Live-Stream/paths/~1live~1v1~1live-streams/get)                                      |
| QueryLiveStreamInfo | [query-a-live-stream](https://docs.visionular.com/auroralive/apireference/#tag/Live-Stream/paths/~1live~1v1~1live-streams~1%7Blive_stream_id%7D/get)              |

### AuroraCloud VOD Api List

| Interfaces    | Reference                                                                                                                        |
|---------------|----------------------------------------------------------------------------------------------------------------------------------|
| QueryTemplate | [query_template](https://docs.visionular.com/auroracloud/apireference/#tag/Template/paths/~1vodencoding~1v1~1query_template/get) |
| ListTemplate  | [list_template](https://docs.visionular.com/auroracloud/apireference/#tag/Template/paths/~1vodencoding~1v1~1list_template/get)   |
| AddStorage    | [add_storage](https://docs.visionular.com/auroracloud/apireference/#tag/Storage/paths/~1vodencoding~1v1~1add_storage/post)       |
| ListStorage   | [list_storage](https://docs.visionular.com/auroracloud/apireference/#tag/Storage/paths/~1vodencoding~1v1~1list_storage/get)      |
| DeleteStorage | [del_storage](https://docs.visionular.com/auroracloud/apireference/#tag/Storage/paths/~1vodencoding~1v1~1del_storage/delete)     |
| CreateTask    | [create_task](https://docs.visionular.com/auroracloud/apireference/#tag/Task/paths/~1vodencoding~1v1~1create_task/post)          |
| QueryTask     | [query_task](https://docs.visionular.com/auroracloud/apireference/#tag/Task/paths/~1vodencoding~1v1~1query_task/get)             |
| ListTask      | [list_task](https://docs.visionular.com/auroracloud/apireference/#tag/Task/paths/~1vodencoding~1v1~1list_task/get)               |

### Errors & Error Handling

All API calls return an err as their final return value. Below is documented the errors you might want to check for. You can check error.code or error.msg on all errors to see the full HTTP response.

### BadParam

BadParam is returned when a you make a bad request to Aurora, this likely means you've passed in an invalid parameter to the API call.

### AuthFail

AuthFail is returned when Aurora cannot authenticate your request.  [You should check you have configured your credentials correctly.](#Authentication)

### EmptyResult

EmptyResult is returned when a resource is not found. This is useful when trying to get an entity by its ID.

### InternalError

InternalError is returned when Aurora returns a HTTP 5XX Status Code. If you encounter this reproducibly, please get in touch
