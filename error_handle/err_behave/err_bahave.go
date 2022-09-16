package errbehave

// An Error represents a network error.
type ServiceErr interface {
	error
	Timeout() bool   // Is the error a timeout?
	Temporary() bool // Is the error temporary?
}

func analysisError(e error) {
	if e == nil {
		return
	}

	if err, ok := e.(ServiceErr); ok {
		if err.Timeout() {
			// do something
		}
		if err.Temporary() {
			// do something
		}
	}

}
