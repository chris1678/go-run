package crons

import (
	"fmt"
	"github.com/chris1678/go-run/logger"
	"github.com/chris1678/go-run/utils"
	"time"

	"github.com/robfig/cron/v3"
)

var retryCount = 3

//var JobList map[string]JobsExec

// var lock sync.Mutex
var CronEngine *cron.Cron

type JobCore struct {
	InvokeTarget   string
	Name           string
	JobId          int
	EntryId        int
	CronExpression string
	Args           string
}

func Initialize(jobList []JobModel) {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	CronEngine = cron.New(cron.WithParser(secondParser), cron.WithChain())

	if len(jobList) == 0 {
		logger.LogHelper.Info("JobCore total:0")
	}
	for i := 0; i < len(jobList); i++ {
		if jobList[i].JobType == 1 {
			j := &HttpJob{}
			j.InvokeTarget = jobList[i].InvokeTarget
			j.CronExpression = jobList[i].CronExpression
			j.JobId = jobList[i].JobId
			j.Name = jobList[i].JobName

			j.addJob()
		} else if jobList[i].JobType == 2 {
			j := &ExecJob{}
			j.InvokeTarget = jobList[i].InvokeTarget
			j.CronExpression = jobList[i].CronExpression
			j.JobId = jobList[i].JobId
			j.Name = jobList[i].JobName
			j.Args = jobList[i].Args
			j.addJob()
		}
	}

	// 其中任务
	CronEngine.Start()
	logger.LogHelper.Info("JobCore start success.")
	// 关闭任务
	defer CronEngine.Stop()
}

// AddJob 添加任务 AddJob(invokeTarget string, jobId int, jobName string, cronExpression string)
//func AddJob(job Job) (int, error) {
//	if job == nil {
//		fmt.Println("unknown")
//		return 0, nil
//	}
//	return job.addJob()
//}

// Remove 移除任务
func Remove(entryID int) chan bool {
	ch := make(chan bool)
	go func() {
		CronEngine.Remove(cron.EntryID(entryID))
		fmt.Println(time.Now().Format(utils.TimeFormat), " [INFO] JobCore Remove success ,info entryID :", entryID)
		ch <- true
	}()
	return ch
}

// 任务停止
//func Stop() chan bool {
//	ch := make(chan bool)
//	go func() {
//		global.GADMCron.Stop()
//		ch <- true
//	}()
//	return ch
//}
