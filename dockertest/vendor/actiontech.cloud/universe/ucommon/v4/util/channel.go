package util

// IsClosed is used when you have a channel that you want to check if it is closed
// if the channel is nil, it will also return true to avoid closing on nil pointer
// WARNING: You have to make sure that no struct has been sent to this channel
func IsClosed(ch <-chan struct{}) bool {
	if ch == nil {
		return true
	}
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func CloseTry(ch chan struct{}) {
	if !IsClosed(ch) {
		close(ch)
	}
}
