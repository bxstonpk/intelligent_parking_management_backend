package events

type imageFrame struct {
	Timestamp int64  `json:"timestamp"`
	cameraID  string `json:"camera_id"`
	data      []byte `json:"data"`
}

type streamResponse struct {
	status  string `json:"status"`
	message string `json:"message"`
}
