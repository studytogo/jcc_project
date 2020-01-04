package filter

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/helper/redis"
	"new_erp_agent_by_go/models/auth"
	"new_erp_agent_by_go/models/record"
	"strconv"
	"strings"
	"time"
)

func CheckAuth(ctx *context.Context) {
	output := new(controllers.Output)
	type Token struct {
		Token string
	}
	var token = Token{}
	token.Token = ctx.Input.Query("token")
	////获取uid,cid,openid
	openId, _ := auth.QueryOpenIdBytoken(token.Token)

	if openId == "" {
		//读取redis,如果redis没有就查询数据库
		keyPrefix := "role_keys" + token.Token
		r, _ := redis.GetOperation(keyPrefix)
		//如果读取的数据是空，则查询数据库，不是空直接判断
		if string(r) == "" {
			authPath, err := auth.AuthRequestAuth(token.Token)
			if err != nil {
				helper.Log.Error("读取数据库失败", err)
				ctx.Output.JSON(output.ErrorOutput("读取数据库失败....."), true, false)
			}
			authPathString := strings.Join(authPath, ",")
			index := strings.Index(authPathString, ctx.Input.URL())
			if index == -1 {
				helper.Log.ErrorString("权限不存在....." + token.Token)
				ctx.Output.JSON(output.ErrorOutput("权限不存在....."), true, false)
			}
		} else {
			verify := strings.Replace(string(r), `\`, "", -1)
			//将获取的验证参数也去掉下划线
			verify = strings.Replace(verify, `_`, "", -1)
			//将带下划线的路由的下划线去掉（php存值的时候不带下划线）
			formateUrl := strings.Replace(strings.ToLower(ctx.Input.URL()), `_`, "", -1)
			index := strings.Index(verify, formateUrl)
			if index == -1 {
				helper.Log.ErrorString("权限不存在....." + token.Token)
				ctx.Output.JSON(output.ErrorOutput("权限不存在....."), true, false)
			}

		}
	}
}

func CommonlyParam(ctx *context.Context) {
	output := new(controllers.Output)
	type Token struct {
		Token string
	}
	var token = Token{}
	token.Token = ctx.Input.Query("token")
	//获取uid,cid,openid
	openId, cid, uid, boss, err := auth.QueryAgentInfo(token.Token)

	if err != nil {
		helper.Log.Error("读取数据库失败", err)
		ctx.Output.JSON(output.ErrorOutput("获取加盟商常用参数失败....."), true, false)
	}

	//将常用参数设置到请求体中
	ctx.Request.Form.Set("OpenId", openId)
	ctx.Request.Form.Set("Companyid", strconv.Itoa(cid))
	ctx.Request.Form.Set("Uid", strconv.Itoa(uid))
	//将请求信息存入数据库
	agentInfo := new(auth.JccRecord)
	agentInfo.Boss = boss
	agentInfo.Api = ctx.Input.URL()
	agentInfo.CreatedAt = int(time.Now().Unix())
	agentInfo.Param = string(fmt.Sprintf("%+v", string(ctx.Input.RequestBody)))
	agentInfo.Ip = ctx.Input.IP()
	agentInfo.Path = ctx.Request.Header.Get("JCC-Path")

	agentInfo.AddJccRecord()
}

func OpenApiRecord(ctx *context.Context) {
	var requestParam = record.JccContextParam{}
	requestParam.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	requestParam.RequsetParam = string(ctx.Input.RequestBody)
	requestParam.Url = ctx.Input.URL()
	type RequestSource struct {
		Source string `json:"source" 请求来源`
	}

	var appParam RequestSource
	err := json.Unmarshal(ctx.Input.RequestBody, &appParam)
	output := new(controllers.Output)
	if err != nil {
		ctx.Output.JSON(output.ErrorOutput("非法身份来源....."), true, false)
		return
	}

	if appParam.Source == "" {
		ctx.Output.JSON(output.ErrorOutput("非法身份来源....."), true, false)
	}

	source := beego.AppConfig.String("source")
	if strings.Index(source, appParam.Source) == -1 {
		helper.Log.ErrorString("非法身份来源.....,IP地址是" + ctx.Input.IP())
		ctx.Output.JSON(output.ErrorOutput("非法身份来源....."), true, false)
		return
	}

	requestParam.AppName = appParam.Source
	_, err = requestParam.AddContext()
	if err != nil {
		helper.Log.Error("插入Context_Param数据库失败", err)
		return
	}
	ctx.Input.SetData("requestId", requestParam.Id)

}
