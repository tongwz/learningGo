package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

type ProcessInfo struct {
	Pid  int32
	Name string
	CPU  float64
	Mem  float64
}

func main() {
	for {
		printTopProcesses()
		time.Sleep(20 * time.Second)
	}
}

var feishuUrl = "https://open.feishu.cn/open-apis/bot/v2/hook/5116eba1-1e8b-4026-9fdf-e995e4911680"

func printTopProcesses() {
	var processes []ProcessInfo

	// 获取所有进程信息
	ps, err := process.Processes()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range ps {
		cpuPercent, _ := p.CPUPercent()
		memoryInfo, _ := p.MemoryInfo()

		if memoryInfo != nil {
			name, _ := p.Name()
			processes = append(processes, ProcessInfo{
				Pid:  p.Pid,
				Name: name,
				CPU:  cpuPercent,
				Mem:  float64(memoryInfo.RSS) / 1024 / 1024, // Convert to MB
			})
		}
	}

	// 按CPU和内存消耗排序
	sort.Slice(processes, func(i, j int) bool {
		if processes[i].CPU == processes[j].CPU {
			return processes[i].Mem > processes[j].Mem
		}
		return processes[i].CPU > processes[j].CPU
	})

	str := ""
	str += fmt.Sprintf("进程占用情况：Top 5 Processes by CPU and Memory Usage:\n")
	for i, p := range processes[:5] {
		str += fmt.Sprintf("%d. PID: %d, Name: %-20s, CPU: %.2f%%, Mem: %.2f MB\n", i+1, p.Pid, p.Name, p.CPU, p.Mem)
	}
	fmt.Println(str)
	FeiShuRobotMsgPost(str)
}

func FeiShuRobotMsgPost(content string) {
	rspInfo := CommonServerRsp{}
	robotMsg := FeiShuRobotMsg{
		Msgtype: "text",
		Content: FeiShuRobotContent{
			Text: content,
		},
	}
	_ = ApiPost(feishuUrl, &rspInfo, robotMsg, time.Second*3)
}

type CommonServerRsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type FeiShuRobotMsg struct {
	Msgtype string             `json:"msg_type"`
	Content FeiShuRobotContent `json:"content"`
}

type FeiShuRobotContent struct {
	Text string `json:"text"`
}

func ApiPost(url string, out interface{}, postBody interface{}, time time.Duration) error {
	// interface{}转化成[]byte
	postData, err := json.Marshal(postBody)
	// fmt.Printf("postBody------------------------------------:%s\n", postData)
	if err != nil {
		fmt.Println("转化成json出错，请求url：", url, "，请求参数：", postBody, err.Error())
		return err
	}

	// 以post方法请求
	newRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
	if err != nil {
		fmt.Println("请求url：", url, "请求参数：", string(postData), "，错误内容：", err.Error())
		return err
	}

	// 以json格式请求
	newRequest.Header.Set("Content-Type", "application/json;charset=utf-8")

	// 发送请求
	client := &http.Client{
		Timeout: time,
	}
	resp, err := client.Do(newRequest)
	if err != nil {
		fmt.Println("请求url：", url, "请求参数：", string(postData), "，错误内容：", err.Error())
		return err
	}

	defer resp.Body.Close()

	outByte, err := io.ReadAll(resp.Body)

	// byte[]转化成interface{}
	err = json.Unmarshal(outByte, out)
	if err != nil {
		fmt.Println("byte[]转化成interface{}出错，请求url：", url, "响应：", string(outByte), "，错误内容：", err.Error())
		return err
	}
	// fmt.Println("POST请求：", url, "，请求参数：", string(postData), ", 返回结果：", string(outByte))

	return nil
}
