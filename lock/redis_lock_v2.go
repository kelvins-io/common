package lock

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisLockInterface interface {
	Set(*redis.Pool, string, uint32) (bool, string, error)
	Release(*redis.Pool, string, string) error
}

type RedisLock struct{}

func (r *RedisLock) Set(redisPool *redis.Pool, key string, expireSecond uint32) (bool, string, error) {
	if expireSecond == 0 {
		return false, "", fmt.Errorf("expireSecond参数必须大于0")
	}

	conn := redisPool.Get()
	defer conn.Close()

	randVal := time.Now().Format("2006-01-02 15:04:05.000")
	reply, err := conn.Do("SET", key, randVal, "NX", "PX", expireSecond*1000)
	if err != nil {
		return false, "", err
	}
	if reply == nil {
		return false, "", nil
	}

	return true, randVal, nil
}

func (r *RedisLock) Release(redisPool *redis.Pool, key string, randVal string) error {
	conn := redisPool.Get()
	defer conn.Close()

	luaScript := `
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end;
	`
	script := redis.NewScript(1, luaScript)
	_, err := script.Do(conn, key, randVal)

	return err
}
