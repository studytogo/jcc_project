package http_query

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type HttpQuery struct {
	Log            *logs.BeeLogger
	dingWebHookUrl string
	runMode        string
}

func NewInit(name string) *HttpQuery {
	var hq HttpQuery
	// 初始化日志
	hq.Log = logs.NewLogger()
	err := hq.Log.SetLogger(logs.AdapterMultiFile, fmt.Sprintf(`{
		"filename":"logs/%s/%s.log",
		"separate":[ "error", "info", "debug"],
		"level":7,
		"maxlines":100000,
		"daily":true,
		"maxdays":10,
		"color":true
	}`, name, name))
	fmt.Println("初始化日志", err)

	return &hq
}

/**
 * 配置日志
 * @Author: cs_shuai
 * @Date: 2019-09-27
 */
func (hq *HttpQuery) SetLogger(config string) error {
	err := hq.Log.SetLogger(logs.AdapterMultiFile, config)

	return err
}

/**
 * 配置钉钉地址
 * @Author: cs_shuai
 * @Date: 2019-09-27
 */
func (hq *HttpQuery) SetDingWebHookUrl(url string) *HttpQuery {
	hq.dingWebHookUrl = url
	return hq
}

/**
 * 配置钉钉地址
 * @Author: cs_shuai
 * @Date: 2019-09-27
 */
func (hq *HttpQuery) SetRunMode(runMode string) *HttpQuery {
	hq.runMode = runMode
	return hq
}

/**
 * 加盟商post请求
 * @Author: cs_shuai
 * @Date: 2019-09-27
 */
func (hq *HttpQuery) AgentPost(path string, params map[string]string) (result map[string]interface{}, err error) {
	params["token"] = "XIAOXIAOLUOSHIGEDASHAX"

	// 参数处理
	var paramsValue = make(url.Values)
	for key, value := range params {
		pValue := []string{value}
		paramsValue[key] = pValue
	}

	paramsJson, _ := json.Marshal(params)

	// 请求接口
	resp, err := http.PostForm(path, paramsValue)
	if err != nil {
		hq.Log.Error(fmt.Sprintf("请求地址: %s, 请求参数: %s , 错误提示: %s", path, string(paramsJson), fmt.Sprint(err)))
		return result, err
	}

	// 接收并记录
	rb, err := ioutil.ReadAll(resp.Body)
	hq.Log.Info(fmt.Sprintf("请求地址: %s, 请求参数: %s , 返回信息: %s", path, string(paramsJson), string(rb)))

	// 判断返回值
	result = make(map[string]interface{})
	err = json.Unmarshal(rb, &result)
	if err != nil {
		hq.Log.Error(fmt.Sprintf("请求地址: %s, 请求参数: %s , 错误提示: %s", path, string(paramsJson), fmt.Sprint(err)))
		return result, err
	}

	return result, err
}

/**
 * 加盟商Get请求
 * @Author: cs_shuai
 * @Date: 2019-09-27
 */
func (hq *HttpQuery) AgentGet(path string, params map[string]string) (result map[string]interface{}, err error) {

	// 处理请求地址
	u, _ := url.Parse(path)

	// 拼接参数
	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()
	paramsJson, _ := json.Marshal(params)

	// 请求接口
	resp, err := http.Get(u.String())
	if err != nil {
		hq.Log.Error(fmt.Sprintf("请求地址: %s, 请求参数: %s ,错误提示: %s", path, string(paramsJson), fmt.Sprint(err)))
		return result, err
	}

	// 接收并记录
	rb, err := ioutil.ReadAll(resp.Body)
	hq.Log.Info(fmt.Sprintf("请求地址: %s, 请求参数: %s ,返回信息: %s", path, string(paramsJson), rb))

	// 判断返回值
	result = make(map[string]interface{})
	err = json.Unmarshal(rb, &result)

	if err != nil {
		hq.Log.Error(fmt.Sprintf("请求地址: %s, 请求参数: %s ,错误提示: %s", path, string(paramsJson), fmt.Sprint(err)))
		return result, err
	}

	return result, err
}

/**
 * 发送钉钉消息
 * @Author: cs_shuai
 * @Date: 2019-09-27
 */
func (hq *HttpQuery) dingWebHook(content string) error {
	if hq.dingWebHookUrl == "" {
		return errors.New("钉钉机器人地址不存在")
	}

	hq.Log.Info("--------------发送钉钉消息---------------")
	jsonValue := []byte(content)
	resp, err := http.Post(hq.dingWebHookUrl, "application/json", bytes.NewBuffer(jsonValue))
	hq.Log.Info(fmt.Sprint(resp))

	return err
}

/**
 * markdown快捷发送钉钉消息
 * @Author: cs_shuai
 * @Date: 2019-09-27
 */
func (hq *HttpQuery) MarkdownDingWebHook(title string, content string) error {

	// 处理消息
	format := `{
		"msgtype": "markdown",
		"markdown": { 
			"title": "【%s】",
			"text" : " %s
					# ** %s \n
					- 发送时间：%s \n"
		}
	}`

	var runModeStr string
	if hq.runMode != "" {
		runModeStr = fmt.Sprintf(`# **【%s】环境, `, hq.runMode)
	}

	body := fmt.Sprintf(
		format,
		title,
		runModeStr,
		content,
		time.Now().Format("2006-01-02 15:04:05"))

	// 发送消息
	err := hq.dingWebHook(body)

	return err
}
