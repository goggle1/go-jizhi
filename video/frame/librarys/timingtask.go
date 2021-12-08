package librarys

import (
	"fmt"
	"sync"
	"time"
)

type TimingTaskType int

const (
	TaskTypeTiming   TimingTaskType = 1 //定时
	TaskTypeInterval TimingTaskType = 2 //间隔
)

type taskInfo struct {
	name     string
	taskType TimingTaskType
	second   int32     //间隔秒数
	exCount  int       //执行次数
	exTime   time.Time //执行时间
	fun      func()
}

//定时任务开启
func (t *taskInfo) startTask() {
	for {

		now := time.Now()
		t1 := now.UTC()
		t2 := t.exTime.UTC()
		dur := t2.Sub(t1)
		if dur.Seconds() <= 0 {
			t.timeReset()
			t.fun()
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func (t *taskInfo) timeReset() {

	dd, _ := time.ParseDuration("24h")
	t.exTime = t.exTime.Add(dd)
}

//间隔任务
func (t *taskInfo) intervalTask() {
	for {
		time.Sleep(time.Duration(t.second) * time.Second)
		if t.exCount <= 0 {
			return
		}
		t.fun()
		t.exCount--
	}
}

type timingTask struct {
	sync.Mutex
	taskmap map[string]*taskInfo
}

func NewTimingTask() *timingTask {
	t := &timingTask{
		taskmap: make(map[string]*taskInfo),
	}
	t.load()
	return t
}

func (t *timingTask) load() {
	go t.checkTask()
}

//间隔任务直接执行
func (t *timingTask) AddIntervalTask(second int32, count int, f func()) {
	t.Lock()
	defer t.Unlock()
	task := &taskInfo{
		taskType: TaskTypeInterval,
		exCount:  count,
		second:   second,
		fun:      f,
	}
	go task.intervalTask()
}

func (t *timingTask) AddTimingTask(hour, minute, second int, f func()) {
	t.Lock()
	defer t.Unlock()
	now := time.Now()
	time := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, second, 0, time.Local)
	len := len(t.taskmap)
	task := &taskInfo{
		name:     fmt.Sprintf("timing-%d", len),
		taskType: TaskTypeTiming,
		exCount:  0,
		exTime:   time,
		fun:      f,
	}
	t.taskmap[task.name] = task
}

func (t *timingTask) checkTask() {
	for {
		now := time.Now()
		for _, v := range t.taskmap {
			t1 := now.UTC()
			t2 := v.exTime.UTC()
			dur := t2.Sub(t1)
			if dur.Seconds() < 0 {
				t.Lock()

				v.timeReset()

				t.Unlock()
			} else if dur.Minutes() <= 1 {
				go v.startTask()
			}

		}
		time.Sleep(1 * time.Minute)
	}
}
