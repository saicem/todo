package db

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/saicem/todo/global"
)

func SearchSession(sessionId string) (int, bool) {
	// todo 不能整个redis全给存这个 需要优化存储策略
	r := global.Redis.Get()
	defer func(r redis.Conn) {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}(r)
	if userId, err := redis.Int(r.Do("GET", sessionKey(sessionId))); err == nil {
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
	if sessionId, err := redis.String(r.Do("GET", userKey(userId))); err == nil {
		_, _ = r.Do("DEl", sessionKey(sessionId))
	}
	if _, err := r.Do("SET", userKey(userId), sessionId); err != nil {
		panic(err)
	}
	if _, err := r.Do("SET", sessionKey(sessionId), userId, "EX", global.Config.Session.MaxAge); err != nil {
		panic(err)
	}
	return true
}

func DELSession(sessionId string) bool {
	r := global.Redis.Get()
	defer func(r redis.Conn) {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}(r)
	if _, err := r.Do("DEL", sessionKey(sessionId)); err == nil {
		return true
	}
	return false
}

func SetCaptcha(captcha string, email string) {
	r := global.Redis.Get()
	defer func(r redis.Conn) {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}(r)
	if _, err := r.Do("SET", captchaKey(email), captcha, "EX", global.Config.Redis.CaptchaDuration); err != nil {
		panic(err)
	}
}

func VerifyCaptcha(captcha string, email string) bool {
	r := global.Redis.Get()
	defer func(r redis.Conn) {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}(r)
	captchaGet, err := redis.String(r.Do("GET", captchaKey(email)))
	if err != nil {
		return false
	}
	return captchaGet == captcha
}

func sessionKey(key string) string {
	return "session:" + key
}

func userKey(key int) string {
	return "user:" + strconv.Itoa(key)
}

func captchaKey(key string) string {
	return "captcha:" + key
}
