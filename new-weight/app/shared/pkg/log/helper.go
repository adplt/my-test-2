package log

// ParseMessage godoc.
func ParseMessage(msg string) string {
	if len(msg) > MAX_LENGTH {
		msg = msg[:MAX_LENGTH-3] + "..."
	}
	return msg
}
