package redis

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/models"

	"github.com/garyburd/redigo/redis"
)

type redisConfig struct {
	redisConn redis.Conn
}

func (r *redisConfig) IsClosed() bool {
	_, err := r.redisConn.Do("ping")
	if err != nil {
		return true
	}

	return false
}

var redisOps redisConfig

func newRedisConn() error {
	conn, err := redis.Dial("tcp", beego.AppConfig.String("redis.conn"))
	if err != nil {
		fmt.Println("连接redis服务器超时。。。。。。。。。。。。。。。。。。。。" + err.Error())
		helper.Log.ErrorString("连接redis服务器超时。。。。。。。。。。。。。。。。。。。。")
		return err
	}
	redisOps.redisConn = conn
	return err
}

func CreatConn() error {
	if redisOps.redisConn == nil || redisOps.IsClosed() {
		err := newRedisConn()
		if err != nil {
			return err
		}
	}

	return nil
}

func KeyExist(key string) string {
	err := CreatConn()
	if err != nil {
		return err.Error()
	}

	isExist, _ := redis.Bool(redisOps.redisConn.Do("EXISTS", key))

	if !isExist {
		//查询数据库是否存在
		appParam, _ := models.QueryAppParam(key)
		if appParam == nil {
			return "验证出现错误请修正后重试........."
		}
		//数据库中存在该值就存入redis缓存中
		err := SetOperation(appParam.AppKey, appParam.AppValue, 0)
		if err != nil {
			return err.Error()
		}

	}

	return ""

}

func SetOperation(key string, value interface{}, duration int) error {
	err := CreatConn()
	if err != nil {
		return err
	}
	if duration == 0 {
		_, err = redisOps.redisConn.Do("SET", key, value)
		return err
	} else {
		_, err = redisOps.redisConn.Do("SETEX", key, duration, value)
		return err
	}
	return nil
}

func GetOperation(key string) ([]byte, error) {
	err := CreatConn()
	if err != nil {
		return nil, err
	}

	value, err := redis.Bytes(redisOps.redisConn.Do("GET", key))
	return value, err
}

func SetJson(key string, message interface{}) error {
	err := CreatConn()
	if err != nil {
		return err
	}
	data, err := json.Marshal(message)

	if err != nil {
		return err
	}

	_, err = redisOps.redisConn.Do("SET", key, data)

	if err != nil {
		return err
	}
	return nil

}

func GetJson(key string) ([]byte, error) {
	err := CreatConn()
	if err != nil {
		return nil, err
	}

	result, err := redis.Bytes(redisOps.redisConn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetString(key string) (string, error) {
	err := CreatConn()
	if err != nil {
		return "", err
	}

	result, err := redis.String(redisOps.redisConn.Do("GET", key))

	if err != nil {
		return "", nil
	}

	return result, nil
}
