package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/streadway/amqp"
	"new_erp_agent_by_go/helper/dingdingAlert"
	"time"
)

var dingToken = ""
var mqurl = ""

func init() {
	dingToken = beego.AppConfig.String("dingdingToken")
	mqurl = beego.AppConfig.String("mqAddress")
}

func mqConnect() (*amqp.Channel, *amqp.Connection, error) {
	var err error
	conn, err := amqp.Dial(mqurl)
	if err != nil {
		return nil, nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	return channel, conn, nil
}

//连接rabbitmq server
func push(exchangeName, routerKey string, message []byte) error {

	channel, connection, err := mqConnect()
	defer channel.Close()
	defer connection.Close()

	if err != nil {
		return err
	}

	err = channel.Publish(exchangeName, routerKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        message,
	})

	if err != nil {
		return err
	}

	return nil
}

func SendRabbitMqMessage(exchangeName, routerKey string, message interface{}) error {

	marshal, err := json.Marshal(message)
	if err != nil {
		DingdingAlert(fmt.Sprint("%+v", message) + "----------------该信息转换错误，需要处理。。。")
		return err
	}
	for {
		err := push(exchangeName, routerKey, marshal)
		if err != nil {
			DingdingAlert(string(marshal) + "发送失败。。。。错误原因：" + err.Error())
			//报错重连
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	return nil
}

// 钉钉mq报警
func DingdingAlert(message string) {
	robot := dingdingAlert.NewRobot(dingToken)
	msg := dingdingAlert.DingMessage{
		Type: "text",
		Text: dingdingAlert.TextElement{
			Content: message,
		},
	}
	_ = robot.SendMessage(msg)
}
