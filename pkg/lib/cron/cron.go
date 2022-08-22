package cron

import (
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

type (
	RobfigCron struct {
		handler *cron.Cron
		jobs    *sync.Map
	}
	jobItem struct {
		ID         string    // 任务标识
		CreateTime time.Time // 任务创建时间
		Count      uint64    // 已经执行次数

		fn   func() // 任务待执行的回调函数
		spec string
		eid  cron.EntryID
	}
)

const (
	SpecOneSecond  = "*/1 * * * * *"
	SpecSymbolInfo = "0 0 2 * * *" // 每天凌晨两点更新数据

	SpecOneMinute  = "0 */1 * * * *"  // 每分钟更新一次
	SpecFiveMinute = "0 */5 * * * *"  // 每5分钟更新一次
	SpecTenMinute  = "0 */10 * * * *" // 每10分钟更新一次

	SpecOneHour = "0 0 */1 * * *" // 每小时更新一次
)

var (
	corns   sync.Map // 定时任务集合
	handler *cron.Cron
)

// 初始化
func New() *RobfigCron {
	// corn对象存在不需要重复创建
	if handler == nil {
		// 设置支持秒级控制
		secondParser := cron.NewParser(cron.Second | cron.Minute |
			cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
		handler = cron.New(cron.WithParser(secondParser), cron.WithChain())
	}
	return &RobfigCron{handler, &corns}
}

// 添加一个定时任务 回调函数
func (c *RobfigCron) Add(id, spec string, fn func()) error {
	eid, err := c.handler.AddFunc(spec, fn)
	if err != nil {
		return err
	}
	job := jobItem{
		ID:         id,
		CreateTime: time.Now(),
		fn:         fn,
		spec:       spec,
		eid:        eid,
	}
	_ = c.Del(id)
	c.jobs.Store(id, &job)
	return nil
}

// 删除一个定时任务
func (c *RobfigCron) Del(id string) error {
	job, ok := c.jobs.Load(id)
	if !ok {
		return nil
	}
	c.handler.Remove(job.(*jobItem).eid)
	c.jobs.Delete(id)
	return nil
}

// 定时任务启动
func (c *RobfigCron) Start() {
	c.handler.Start()
}

// 定时任务暂停
func (c *RobfigCron) Stop() {
	c.handler.Stop()
}
