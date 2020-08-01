package lock

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type redisLock struct {
	redisPool *redis.Pool

	key       string
	randVal   string
}

func NewRedisLock(redis *redis.Pool, redisKey string) (*redisLock, error) {
	if redis == nil {
		return nil, fmt.Errorf("redis不能为nil")
	}
	if redisKey == "" {
		return nil, fmt.Errorf("redisKey不能为空")
	}

	return &redisLock{
		redisPool: redis,
		key:       redisKey,
		randVal:   string(time.Now().UnixNano()),
	}, nil
}

func (r *redisLock) Set(expireSecond int32) (bool, error) {
	if expireSecond <= 0 {
		return false, fmt.Errorf("expireSecond 参数必须大于0")
	}

	conn := r.redisPool.Get()
	defer conn.Close()

	reply, err := conn.Do("SET", r.key, r.randVal, "NX", "PX", expireSecond*1000)
	if err != nil {
		return false, err
	}
	if reply == nil {
		return false, nil
	}

	return true, nil
}

func (r *redisLock) Release() error {
	conn := r.redisPool.Get()
	defer conn.Close()

	luaScript := `
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end;
	`
	script := redis.NewScript(1, luaScript)
	_, err := script.Do(conn, r.key, r.randVal)
	return err
}
