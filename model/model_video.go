package model

type AddTemplateApiRequest struct {
	TemplateName string  `json:"template_name"`
	Format       string  `json:"format"`
	Quality      float32 `json:"quality"`
	Framerate    int     `json:"framerate"`
	Resolution   string  `json:"resolution"`
	VCodec       string  `json:"vcodec"`
	ACodec       string  `json:"acodec"`
	AudioBitrate int     `json:"audioBitrate"`
}

type TemplateResponse struct {
	Code      int      `json:"code"`
	Msg       string   `json:"msg"`
	Data      struct{} `json:"data,omitempty"`
	RequestID string   `json:"request_id"`
}

type TemplateApiResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      TemplateApi `json:"data,omitempty"`
	RequestID string      `json:"request_id"`
}

type TemplateApi struct {
	TemplateName    string            `json:"template_name"`
	TransMode       string            `json:"trans_mode,omitempty"`
	Format          string            `json:"format"`
	Vcodec          string            `json:"vcodec"`
	Acodec          string            `json:"acodec,omitempty"`
	Resolution      string            `json:"resolution,omitempty"`
	FrameRate       int               `json:"framerate,omitempty"`
	Quality         float64           `json:"quality,omitempty"`
	VideoBitrate    int               `json:"video_bitrate,omitempty"`
	GopSize         int               `json:"gop_size,omitempty"`
	GopSeconds      int               `json:"gop_seconds,omitempty"`
	SegType         string            `json:"seg_type,omitempty"`
	SegTime         int               `json:"seg_time,omitempty"`
	AudioBitrate    int               `json:"audio_bitrate,omitempty"`
	AudioSampleRate string            `json:"audio_sample_rate,omitempty"`
	AudioChannels   int               `json:"audio_channels,omitempty"`
	LogoPath        string            `json:"logo_path,omitempty"`
	LogoSize        string            `json:"logo_size,omitempty"`
	LogoOffset      string            `json:"logo_offset,omitempty"`
	AdaptiveStreams *[]AdaptiveStream `json:"adaptive_streams,omitempty"`
	Encryption      *EncryptConfig    `json:"encryption,omitempty"`
	ThumbnailSets   *ThumbnailSet     `json:"thumbnail"`
	Description     string            `json:"description,omitempty"`
	CreatedAt       int64             `json:"created_at,omitempty"`
	BuiltIn         bool              `json:"built_in,omitempty"`
}

type AdaptiveStream struct {
	Vcodec       string  `json:"vcodec"`
	Quality      float64 `json:"quality,omitempty"`
	Resolution   string  `json:"resolution,omitempty" `
	FrameRate    int     `json:"framerate,omitempty"`
	VideoBitrate int     `json:"video_bitrate,omitempty"`
	AudioBitrate int     `json:"audio_bitrate,omitempty"`
}

type EncryptConfig struct {
	Provider   string            `json:"provider,omitempty"`
	Mode       string            `json:"mode,omitempty"`
	DrmSystems []string          `json:"drm_systems,omitempty"`
	KeyUrl     string            `json:"key_url"`
	HttpMethod string            `json:"http_method,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
}

type ThumbnailSet struct {
	StartPoint int `json:"start_point,omitempty"`
	Interval   int `json:"interval,omitempty"`
	Number     int `json:"number,omitempty"`
	Width      int `json:"width,omitempty"`
	Height     int `json:"height,omitempty"`
}

type TemplateApiRequest struct {
	TemplateName    string            `json:"template_name"`
	Format          string            `json:"format"`
	Vcodec          string            `json:"vcodec"`
	Acodec          string            `json:"acodec"`
	TransMode       string            `json:"trans_mode"`
	Resolution      string            `json:"resolution"`
	FrameRate       int               `json:"framerate"`
	Quality         float64           `json:"quality"`
	VideoBitrate    int               `json:"video_bitrate"`
	GopSize         int               `json:"gop_size"`
	GopSeconds      int               `json:"gop_seconds"`
	SegType         string            `json:"seg_type"`
	SegTime         int               `json:"seg_time"`
	AudioBitrate    int               `json:"audio_bitrate"`
	AudioSampleRate string            `json:"audio_sample_rate"`
	AudioChannels   int               `json:"audio_channels"`
	LogoPath        string            `json:"logo_path"`
	LogoSize        string            `json:"logo_size"`
	LogoOffset      string            `json:"logo_offset"`
	AdaptiveStreams *[]AdaptiveStream `json:"adaptive_streams"`
	Encryption      *EncryptConfig    `json:"encryption"`
	ThumbnailSets   *ThumbnailSet     `json:"thumbnail"`
	Description     string            `json:"description"`
}

type UpdateTemplateResponse struct {
	Code      int      `json:"code"`
	Msg       string   `json:"msg"`
	Data      struct{} `json:"data,omitempty"`
	RequestID string   `json:"request_id"`
}

type TemplateListApiResponse struct {
	Code      int             `json:"code"`
	Msg       string          `json:"msg"`
	Data      TemplateListApi `json:"data,omitempty"`
	RequestID string          `json:"request_id"`
}

type TemplateListApi struct {
	Num  int                   `json:"num"`
	List []TemplateApiResponse `json:"list"`
}

type StorageApiRequest struct {
	Type      string `json:"type"`
	Region    string `json:"region"`
	Bucket    string `json:"bucket"`
	Prefix    string `json:"prefix,omitempty"`
	Notify    string `json:"notify,omitempty"`
	StorageAk string `json:"storage_ak,omitempty"`
	StorageSk string `json:"storage_sk,omitempty"`
}

type StorageApiResponse struct {
	StorageId string `json:"storage_id"`
	Type      string `json:"type"`
	Region    string `json:"region"`
	Bucket    string `json:"bucket"`
	Prefix    string `json:"prefix,omitempty"`
	Notify    string `json:"notify,omitempty"`
	Extra     string `json:"extra,omitempty"`
	NameRule  string `json:"name_rule,omitempty"`
	CreatedAt int64  `json:"created_time"`
}

type StorageListApiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Num  int                  `json:"num"`
		List []StorageApiResponse `json:"list"`
	} `json:"data,omitempty"`
	RequestID string `json:"request_id"`
}

type StorageDelResponse struct {
	Code      int      `json:"code"`
	Msg       string   `json:"msg"`
	Data      struct{} `json:"data,omitempty"`
	RequestID string   `json:"request_id"`
}

type CreateTaskRequest struct {
	TemplateName string `json:"template_name"`
	StorageID    string `json:"storage_id"`
	Input        string `json:"input"`
	Output       string `json:"output"`
}

type TaskResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TaskID string `json:"task_id"`
	} `json:"data,omitempty"`
	RequestID string `json:"request_id"`
}
