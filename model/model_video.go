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
	ThumbnailSets   *ThumbnailSet     `json:"thumbnail"       validate:"omitempty"`
	Description     string            `json:"description,omitempty"`
	CreatedAt       int64             `json:"created_at,omitempty"`
	BuiltIn         bool              `json:"built_in,omitempty"`
}

type AdaptiveStream struct {
	Vcodec       string  `json:"vcodec"                   validate:"required,oneof=h264 h265 av1 copy auto none ffmpeg_aac"`
	Quality      float64 `json:"quality,omitempty"        validate:"omitempty,lte=10"`
	Resolution   string  `json:"resolution,omitempty"     validate:"omitempty,rsl,max=20"`
	FrameRate    int     `json:"framerate,omitempty"      validate:"omitempty,lte=120"`
	VideoBitrate int     `json:"video_bitrate,omitempty"  validate:"omitempty"`
	AudioBitrate int     `json:"audio_bitrate,omitempty"  validate:"omitempty"`
}

type EncryptConfig struct {
	Provider   string            `json:"provider,omitempty"               validate:"omitempty,oneof=Normal ExpressPlay"`
	Mode       string            `json:"mode,omitempty"                   validate:"omitempty,oneof=AES-128 SAMPLE-AES"`
	DrmSystems []string          `json:"drm_systems,omitempty"            validate:"omitempty,min=0,max=4,dive,oneof=Marlin PlayReady Widevine FairPlay"`
	KeyUrl     string            `json:"key_url"                validate:"required,url,max=512"`
	HttpMethod string            `json:"http_method,omitempty"  validate:"omitempty,oneof=GET POST HEAD PUT"`
	Headers    map[string]string `json:"headers,omitempty"      validate:"omitempty,max=10,dive,keys,gte=1,lte=64,endkeys,lte=256"`
}

type ThumbnailSet struct {
	StartPoint int `json:"start_point,omitempty"      validate:"omitempty,gte=0,lte=10000"`
	Interval   int `json:"interval,omitempty"      validate:"required,gte=1,lte=10000"`
	Number     int `json:"number,omitempty"      validate:"required,gte=1,lte=1000"`
	Width      int `json:"width,omitempty"      validate:"omitempty,numberEven,gte=32,lte=2000"`
	Height     int `json:"height,omitempty"      validate:"omitempty,numberEven,gte=32,lte=2000"`
}

type TemplateApiRequest struct {
	TemplateName    string            `json:"template_name"          validate:"required,idstr"`
	Format          string            `json:"format"                 validate:"omitempty,oneof=mp4 3gp hls flv dash cmaf mpegts image2 auto"`
	Vcodec          string            `json:"vcodec"                 validate:"omitempty,oneof=h264 h265 av1 copy auto"`
	Acodec          string            `json:"acodec"                 validate:"omitempty,oneof=aac ac3 libvo_amrwbenc copy auto none"`
	TransMode       string            `json:"trans_mode"             validate:"omitempty,oneof=normal adaptive"`
	Resolution      string            `json:"resolution"             validate:"omitempty,rsl,max=20"`
	FrameRate       int               `json:"framerate"              validate:"omitempty,lte=120"`
	Quality         float64           `json:"quality"                validate:"omitempty,lte=10"`
	VideoBitrate    int               `json:"video_bitrate"          validate:"omitempty"`
	GopSize         int               `json:"gop_size"               validate:"omitempty,gte=0,lte=1000"`
	GopSeconds      int               `json:"gop_seconds"            validate:"omitempty,gte=0,lte=1000"`
	SegType         string            `json:"seg_type"               validate:"omitempty,max=20"`
	SegTime         int               `json:"seg_time"               validate:"omitempty"`
	AudioBitrate    int               `json:"audio_bitrate"          validate:"omitempty"`
	AudioSampleRate string            `json:"audio_sample_rate"      validate:"omitempty,oneof=96k 88.2k 64k 48k 44.1k 32k 24k 22.05k 16k 12k 11.025k 8k 7.35k"`
	AudioChannels   int               `json:"audio_channels"         validate:"omitempty,lte=10"`
	LogoPath        string            `json:"logo_path"              validate:"omitempty,max=2048"`
	LogoSize        string            `json:"logo_size"              validate:"omitempty,lsz,max=64"`
	LogoOffset      string            `json:"logo_offset"            validate:"omitempty,lof,max=64"`
	AdaptiveStreams *[]AdaptiveStream `json:"adaptive_streams" validate:"omitempty,dive,min=2,max=10"`
	Encryption      *EncryptConfig    `json:"encryption"       validate:"omitempty,dive"`
	ThumbnailSets   *ThumbnailSet     `json:"thumbnail"       validate:"omitempty,dive"`
	Description     string            `json:"description"                 validate:"omitempty,min=1,max=256"`
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
	Type      string `json:"type"                     validate:"required,oneof=s3 oss obs kodo cos ks3 http file auto"`
	Region    string `json:"region"                   validate:"required,idstr,max=20"`
	Bucket    string `json:"bucket"                   validate:"required,idstr"`
	Prefix    string `json:"prefix,omitempty"         validate:"omitempty,max=128"`
	Notify    string `json:"notify,omitempty"         validate:"omitempty,max=1024"`
	StorageAk string `json:"storage_ak,omitempty"     validate:"omitempty,printascii,min=3,max=64"`
	StorageSk string `json:"storage_sk,omitempty"     validate:"omitempty,printascii,min=3,max=128"`
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
