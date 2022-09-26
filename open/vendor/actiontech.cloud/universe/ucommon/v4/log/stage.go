package log

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"

	_trace "actiontech.cloud/universe/ucommon/v4/trace"
)

type Stage struct {
	stages  []string
	stagei  int
	traceID string
	mutex   sync.Mutex
}

const (
	StageIndexLimit = 100
	StageId         = "stage_id"
	StageName       = "stage_name"
)

func NewStage() *Stage {
	s := newEmptyStage()
	s.traceID = _trace.NewTraceID()
	return s
}

func newEmptyStage() *Stage {
	return &Stage{
		stages: make([]string, StageIndexLimit, StageIndexLimit),
		stagei: -1,
	}
}

func NewStageFromTraceID(traceID string) (st *Stage) {

	if traceID == "" {
		st = NewStage()
		UserWarn(st, "trace id is empty, new trace id is <%v>", st.GetTraceID())
	} else {
		st = newEmptyStage()
		st.traceID = traceID
	}

	return st
}

func genRandomThreadId() string {
	seq := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	l := len(seq)
	a := rand.Intn(l * l * l)
	return fmt.Sprintf("%c%c%c", seq[a%l], seq[(a/l)%l], seq[(a/l/l)%l])
}

func (s *Stage) RefreshThreadId() *Stage {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.traceID = _trace.NewTraceID()
	return s
}

func (s *Stage) Enter(desc string) *Stage {
	return s.Enterf(desc)
}

func (s *Stage) Enterf(desc string, args ...interface{}) *Stage {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.stagei >= StageIndexLimit {
		panic(fmt.Sprintf("log.Stage exceed StageIndexLimit: <%v>", StageIndexLimit))
	}
	s.stagei++
	if len(args) == 0 {
		s.stages[s.stagei] = desc
	} else {
		s.stages[s.stagei] = fmt.Sprintf(desc, args...)
	}

	return s
}

func (s *Stage) Exit() *Stage {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.stagei < 0 {
		panic("log.Stage exceed 0")
	}
	s.stagei--
	return s
}

func (s *Stage) Go() *Stage {
	ret := NewStage()
	s.mutex.Lock()
	defer s.mutex.Unlock()
	copy(ret.stages, s.stages)
	ret.stagei = s.stagei
	return ret
}

func (s *Stage) GoWithTraceID() (st *Stage) {
	st = newEmptyStage()
	s.mutex.Lock()
	copy(st.stages, s.stages)
	st.traceID = s.traceID
	st.stagei = s.stagei
	s.mutex.Unlock()
	return st
}

// func (s *Stage) GoNew() *Stage {
// 	return NewStage()
// }

func (s *Stage) ToPrefix() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.stagei < 0 {
		return ""
	}
	return " " + "|" + s.traceID + "| <" + strings.Join(s.stages[0:s.stagei+1], ".") + ">"
}
func (s *Stage) String() string {
	return s.ToPrefix()
}

// Deprecated: use GetTraceID() instead
func (s *Stage) GetStageId() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.traceID
}

func (s *Stage) SetTraceID(traceID string) {
	s.mutex.Lock()
	s.traceID = traceID
	s.mutex.Unlock()
}

func (s *Stage) GetTraceID() (traceID string) {
	s.mutex.Lock()
	traceID = s.traceID
	s.mutex.Unlock()
	return
}

// func (s *Stage) GetStagesCount() int {
// 	s.mutex.Lock()
// 	defer s.mutex.Unlock()
// 	return s.stagei + 1
// }

func (s *Stage) GetFirstStageName() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.stagei < 0 {
		return ""
	}
	return s.stages[0]
}

func (s *Stage) GetLastStageName() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.stagei < 0 {
		return ""
	}
	return s.stages[s.stagei]
}

func (s *Stage) GetAllStageNames() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.stagei < 0 {
		return ""
	}
	return strings.Join(s.stages[0:s.stagei+1], ".")
}

// func (s *Stage) GetFirstAndLastStageNames() string {
// 	count := s.GetStagesCount()
// 	if count == 0 {
// 		return ""
// 	} else if count == 1 {
// 		return s.GetFirstStageName()
// 	} else if count == 2 {
// 		return fmt.Sprintf("%v.%v", s.GetFirstStageName(), s.GetLastStageName())
// 	} else {
// 		return fmt.Sprintf("%v...%v", s.GetFirstStageName(), s.GetLastStageName())
// 	}
// }
