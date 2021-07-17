package db

import (
	"github.com/gomodule/redigo/redis"
	"github.com/saicem/todo/global"
)

func SearchSession(sessionId string) (int, bool) {
	// todo 不能整个redis全给存这个 需要优化存储策略
	r := global.Redis.Get()
	defer func(r redis.Conn) {
		err := r.Close()
		if err != nil {
			panic("关不掉？？")
		}
	}(r)
	if userId, err := redis.Int(r.Do("GET", sessionId)); err == nil {
		return userId, true
	}
	return 0, false
}

func SetSession(sessionId string, userId int) bool {
	r := global.Redis.Get()
	defer func(r redis.Conn) {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}(r)
	if sessionId, err := redis.String(r.Do("GET", userId)); err == nil {
		_, _ = r.Do("DEl", sessionId)
	}
	if _, err := r.Do("SET", userId, sessionId); err != nil {
		panic(err)
	}
	if _, err := r.Do("SET", sessionId, userId, "EX", global.Config.Session.MaxAge); err != nil {
		panic(err)
	}
	return true
}

func DELSession(sessionId string) bool {
	r := global.Redis.Get()
	defer func(r redis.Conn) {
		err := r.Close()
		if err != nil {
			panic("关不掉？？")
		}
	}(r)
	if _, err := r.Do("DEL", sessionId); err == nil {
		return true
	}
	return false
}
