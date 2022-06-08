/*
@Date : 2022/6/8 15:34
@Description go-run
@Author : github.com/chris1678
*/
package crons

import (
	"fmt"
	"github.com/chris1678/go-run/logger"
	"github.com/chris1678/go-run/utils"
	"time"
)

var JobList map[string]Job

type ExecJob struct {
	JobCore
}

func (j *ExecJob) Run() {
	startTime := time.Now()
	var obj = JobList[j.InvokeTarget]

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
func (j *ExecJob) addJob() (int, error) {
	id, err := CronEngine.AddJob(j.CronExpression, j)
	if err != nil {
		fmt.Println(time.Now().Format(utils.TimeFormat), " [ERROR] JobCore AddJob error", err)
		return 0, err
	}
	EntryId := int(id)
	return EntryId, nil
}
