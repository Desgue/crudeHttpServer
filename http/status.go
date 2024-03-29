package http

const (
	StatusOk = 200
)

func StatusToText(status int) string {
	switch status {
	case StatusOk:
		return "OK"
	default:
		return "Unknown"
	}
}
