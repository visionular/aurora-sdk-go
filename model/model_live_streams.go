package model

// CreateStreamRequest  Request
type CreateStreamRequest struct {
	Type              string                   `json:"type" binding:"oneof='' transfer normal"`
	ReconnectWindow   *int                     `json:"reconnect_window"`
	UserMetadata      string                   `json:"user_metadata"`
	SimulcastTargets  []*SimulcastTargets      `json:"simulcast_targets"`
	Policy            string                   `json:"policy"`
	Record            bool                     `json:"record"`
	WatermarkSettings *StreamWatermarkSettings `json:"watermark_settings,omitempty"`
}

type SimulcastTargets struct {
	StreamKey    string `json:"stream_key"`
	Url          string `json:"url"`
	UserMetadata string `json:"user_metadata"`
}

type StreamWatermarkSettings struct {
	Url             string                `json:"url" binding:"url"`
	OverlaySettings StreamOverlaySettings `json:"overlay_settings"`
}

type StreamOverlaySettings struct {
	Width            string `json:"width" binding:"omitempty,pic_percent|pic_pixel"`
	Height           string `json:"height" binding:"omitempty,pic_percent|pic_pixel"`
	VerticalMargin   string `json:"vertical_margin" binding:"omitempty,margin_percent|margin_pixel"`
	VerticalAlign    string `json:"vertical_align" binding:"omitempty,oneof=top middle bottom"`
	HorizontalMargin string `json:"horizontal_margin" binding:"omitempty,margin_percent|margin_pixel"`
	HorizontalAlign  string `json:"horizontal_align" binding:"omitempty,oneof=left center right"`
	Opacity          string `json:"opacity" binding:"omitempty,pic_percent"`
}

// LiveStreamResponse Response
type LiveStreamResponse struct {
	Code      int        `json:"code"`
	Msg       string     `json:"msg"`
	Data      LiveStream `json:"data,omitempty"`
	RequestID string     `json:"request_id"`
}

type LiveStream struct {
	StreamKey       string `json:"stream_key"`
	Status          string `json:"status"`
	ReconnectWindow int    `json:"reconnect_window"`
	ID              string `json:"id"`
	UserMetadata    string `json:"user_metadata"`
	CreateAt        int    `json:"create_at"`
	PlaybackIDS     []struct {
		Policy string `json:"policy"`
		ID     string `json:"id"`
	} `json:"playback_ids"`
	Type   string `json:"type"`
	Record bool   `json:"record"`
}

// DisableLiveStreamResponse Request
type DisableLiveStreamResponse struct {
	Code      int      `json:"code"`
	Msg       string   `json:"msg"`
	Data      struct{} `json:"data,omitempty"`
	RequestID string   `json:"request_id"`
}

// EnableLiveStreamResponse Request
type EnableLiveStreamResponse struct {
	Code      int      `json:"code"`
	Msg       string   `json:"msg"`
	Data      struct{} `json:"data,omitempty"`
	RequestID string   `json:"request_id"`
}

// ListLiveStreamsResponse  Response
type ListLiveStreamsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		StreamKey         string                   `json:"stream_key"`
		Status            string                   `json:"status"`
		ReconnectWindow   int                      `json:"reconnect_window"`
		UserMetadata      string                   `json:"user_metadata"`
		ID                string                   `json:"id"`
		Type              string                   `json:"type"`
		PlaybackIDs       []*Playback              `json:"playback_ids,omitempty"`
		CreatedAt         int64                    `json:"created_at"`
		Record            bool                     `json:"record"`
		WatermarkSettings *StreamWatermarkSettings `json:"watermark_settings,omitempty"`
	} `json:"data,omitempty"`
	RequestID string `json:"request_id"`
}

type Playback struct {
	Policy string `json:"policy"`
	ID     string `json:"id"`
}
