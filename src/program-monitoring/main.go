package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// 自己写的检测 不写配置文件，提醒也写死
var NoticeUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=6aacc64d-9fc8-452f-b2d1-aeb94d291a5e"

var supervisorPrograms = []string{
	"integration:integrationServiceApi",
	"integration:integrationServiceCrontab",
	"integration:integrationServiceQueue",
}

func main() {
	for _, program := range supervisorPrograms {
		cmd := exec.Command("bash", "-c", "supervisorctl status | grep "+program+"| grep -a RUNNING | wc -l")
		// cmd := exec.Command("bash", "-c", "ls -l|grep -a main.go1 |wc -l")

		res, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("执行指令失败：%s , %s \n", err.Error(), cmd.String())
			continue
		}
		// 去掉换行符
		newStr := strings.Replace(string(res), "\n", "", 10)
		// 将[]byte转化成数字
		runningCount, err := strconv.Atoi(newStr)
		if err != nil {
			fmt.Printf("字符串转化成整数失败：%s  \n", err.Error())
			continue
		}
		// 如果数量是0 直接进行提醒并进行重启服务
		if runningCount == 0 {
			fmt.Printf("查询running异常：%s, 直接重启 \n", program)
			cmd = exec.Command("bash", "-c", "supervisorctl start "+program)
			res, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("执行指令失败：%s , %s \n", err.Error(), cmd.String())
				continue
			}
			resStr := string(res)
			resStr = strings.Replace(resStr, "\n", "", 10)
			fmt.Printf("启动结果：%s \n", resStr)
			if strings.Contains(resStr, "ERROR") && !strings.Contains(resStr, "already started") {
				fmt.Printf("项目启动失败：%s , %s \n", resStr, cmd.String())
			}
			_ = httpSendRobot(fmt.Sprintf("查询running异常：%s, 直接重启, 重启结果：%s \n", program, resStr))
		}
	}

}

// http请求
func httpSendRobot(content string) error {
	// 设置请求的 Body 数据，这里以 JSON 格式为例
	var reqInfo = map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": content,
		},
	}
	reqJson, _ := json.Marshal(reqInfo)

	// 创建一个 HTTP 请求，并设置请求方法、URL、Body 等信息
	request, err := http.NewRequest("POST", NoticeUrl, bytes.NewBuffer(reqJson))
	if err != nil {
		fmt.Printf("机器人创建请求失败:%s \n", err.Error())
		return err
	}
	// 设置请求的 Header 信息，根据需要设置相应的 Header
	request.Header.Set("Content-Type", "application/json")
	// 创建一个 HTTP Client，并发送请求
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("机器人发送请求失败:%s \n", err.Error())
		return err
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Printf("机器人请求响应信息:%s \n", string(bodyBytes))
	return nil
}
