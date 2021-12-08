package cmdhandler

import (
	"encoding/json"
	"sync"
	"time"

	"git.tvblack.com/video/frame/utils"
)

type StaticsWatch struct {
	ID         int32
	Server     string
	Service    string
	Desc       string
	MaxTiem    float64
	AvgTime    float64
	Count      int
	CreateTime time.Time
	FailCount  int
}

type staticJson struct {
	StartTime string
	Service   []*cmdStaticJson `json:"cmds"`
}
type cmdStaticJson struct {
	Name      string  `json:"name"`
	Count     int     `json:"count"`
	AvgTime   float64 `json:"avgTime"`
	MaxTiem   float64 `json:"maxTiem"`
	FailCount int     `json:"failCount"`
}

var ServiceStatics = newStatics()

type serviceStatics struct {
	sync.RWMutex
	watchs    map[string]*StaticsWatch
	startTime string
}

func newStatics() *serviceStatics {
	stat := &serviceStatics{
		watchs:    make(map[string]*StaticsWatch),
		startTime: utils.TimeToStr(time.Now()),
	}
	return stat
}

func (s *serviceStatics) Statics(cmd string, m float64) {
	s.Lock()
	defer s.Unlock()
	watch, ok := s.watchs[cmd]
	if !ok {
		watch = &StaticsWatch{
			CreateTime: time.Now(),
		}
		s.watchs[cmd] = watch
	}

	if m > watch.MaxTiem {
		watch.MaxTiem = m
	}
	f := watch.AvgTime * float64(watch.Count)
	watch.AvgTime = (f + m) / float64(watch.Count+1)
	watch.Count++
}

func (s *serviceStatics) GetStatics() string {
	cmds := &staticJson{
		StartTime: s.startTime,
	}
	for k, v := range s.watchs {
		s := &cmdStaticJson{
			Name:      k,
			Count:     v.Count,
			AvgTime:   v.AvgTime,
			MaxTiem:   v.MaxTiem,
			FailCount: v.FailCount,
		}
		cmds.Service = append(cmds.Service, s)
	}
	b, err := json.Marshal(cmds)
	if err != nil {
		return "error"
	}
	return string(b)
}
