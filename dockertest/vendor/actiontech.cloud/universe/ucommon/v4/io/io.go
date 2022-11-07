package io

import (
	"errors"
	"fmt"
	"io"
	"time"
)

type ErrTimeout string

func (e ErrTimeout) Error() string {
	return "Err: time out: " + string(e)
}

//CopyWithTimeout is an Enhancement of io.Copy in golang std lib
func CopyWithTimeout(dst io.Writer, src io.Reader, timeoutSeconds int) (written int64, err error) {
	return copyBufferWithTimeout(dst, src, nil, timeoutSeconds)
}

// Enhancement of io.copyBuffer in golang std lib
func copyBufferWithTimeout(dst io.Writer, src io.Reader, buf []byte, timeoutSeconds int) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	if wt, ok := src.(io.WriterTo); ok {
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(io.ReaderFrom); ok {
		return rt.ReadFrom(src)
	}
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*io.LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	begin := time.Now()
	timeoutChan := make(chan struct{}, 1)
	endFromTimeout := begin.Add(time.Second * time.Duration(timeoutSeconds))
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = errors.New("invalid write result")
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}

		if time.Now().After(endFromTimeout) {
			timeoutChan <- struct{}{}
		}

		select {
		case <-timeoutChan:
			err = ErrTimeout(fmt.Sprintf("copy buffer time out, limit [%v] seconds", timeoutSeconds))
			return 0, err
		default:
		}
	}
	return written, err
}
