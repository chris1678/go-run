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

func (h *HttpJob) addJob() (int, error) {
	id, err := CronEngine.AddJob(h.CronExpression, h)
	if err != nil {
		fmt.Println(time.Now().Format(utils.TimeFormat), " [ERROR] JobCore AddJob error", err)
		return 0, err
	}
	EntryId := int(id)
	return EntryId, nil
}
