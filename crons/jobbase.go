package crons

import (
	"finance/app/public/models"
	"fmt"
	"github.com/chris1678/go-run/logger"
	"github.com/chris1678/go-run/utils"
	"time"

	"github.com/robfig/cron/v3"
)

var retryCount = 3

var jobList map[string]JobsExec

// var lock sync.Mutex

type JobCore struct {
	InvokeTarget   string
	Name           string
	JobId          int
	EntryId        int
	CronExpression string
	Args           string
}

// HttpJob 任务类型 http
type HttpJob struct {
	JobCore
}

// Run http 任务接口
func (h *HttpJob) Run() {

	startTime := time.Now()
	var count = 0
	var err error
	var str string
	/* 循环 */
LOOP:
	if count < retryCount {
		/* 跳过迭代 */
		str, err = utils.Get(h.InvokeTarget)
		if err != nil {
			// 如果失败暂停一段时间重试
			fmt.Println(time.Now().Format(utils.TimeFormat), " [ERROR] mission failed! ", err)
			fmt.Printf(time.Now().Format(utils.TimeFormat)+" [INFO] Retry after the task fails %d seconds! %s \n", (count+1)*5, str)
			time.Sleep(time.Duration(count+1) * 5 * time.Second)
			count = count + 1
			goto LOOP
		}
	}
	// 结束时间
	endTime := time.Now()

	// 执行时间
	latencyTime := endTime.Sub(startTime)
	//TODO: 待完善部分

	logger.LogHelper.Infof("[Job] JobCore %s exec success , spend :%v", h.Name, latencyTime)
	// return
}

func (h *HttpJob) addJob(c *cron.Cron) (int, error) {
	id, err := c.AddJob(h.CronExpression, h)
	if err != nil {
		fmt.Println(time.Now().Format(utils.TimeFormat), " [ERROR] JobCore AddJob error", err)
		return 0, err
	}
	EntryId := int(id)
	return EntryId, nil
}

type ExecJob struct {
	JobCore
}

func (j *ExecJob) Run() {
	startTime := time.Now()
	var obj = jobList[j.InvokeTarget]

	if obj == nil {
		logger.LogHelper.Warn("[Job] ExecJob Run job nil")
		return
	}
	err := CallExec(obj.(JobsExec), j.Args)
	if err != nil {
		// 如果失败暂停一段时间重试
		fmt.Println(time.Now().Format(utils.TimeFormat), " [ERROR] mission failed! ", err)
	}
	// 结束时间
	endTime := time.Now()

	// 执行时间
	latencyTime := endTime.Sub(startTime)
	//TODO: 待完善部分
	//str := time.Now().Format(TimeFormat) + " [INFO] JobCore " + string(j.EntryId) + "exec success , spend :" + latencyTime.String()
	//ws.SendAll(str)
	logger.LogHelper.Info("[Job] JobCore %s exec success , spend :%v", j.Name, latencyTime)
	return
}
func in(jobList []models.SysJob) {
	crontab := Engine
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

			AddJob(crontab, j)
		} else if jobList[i].JobType == 2 {
			j := &ExecJob{}
			j.InvokeTarget = jobList[i].InvokeTarget
			j.CronExpression = jobList[i].CronExpression
			j.JobId = jobList[i].JobId
			j.Name = jobList[i].JobName
			j.Args = jobList[i].Args
			AddJob(crontab, j)
		}
	}

	// 其中任务
	crontab.Start()
	logger.LogHelper.Info("JobCore start success.")
	// 关闭任务
	defer crontab.Stop()
	select {}
}
func (j *ExecJob) addJob(c *cron.Cron) (int, error) {
	id, err := c.AddJob(j.CronExpression, j)
	if err != nil {
		fmt.Println(time.Now().Format(utils.TimeFormat), " [ERROR] JobCore AddJob error", err)
		return 0, err
	}
	EntryId := int(id)
	return EntryId, nil
}

// AddJob 添加任务 AddJob(invokeTarget string, jobId int, jobName string, cronExpression string)
func AddJob(c *cron.Cron, job Job) (int, error) {
	if job == nil {
		fmt.Println("unknown")
		return 0, nil
	}
	return job.addJob(c)
}

// Remove 移除任务
func Remove(c *cron.Cron, entryID int) chan bool {
	ch := make(chan bool)
	go func() {
		c.Remove(cron.EntryID(entryID))
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
