package lib

import (
	"time"
)

func GetRedisLock(key string, expiredIn time.Duration) (int64, error) {
	timestamp := time.Now().UnixNano()
	set, err := Client.SetNX(key, timestamp, expiredIn).Result()
	if err != nil {
		return 0, err
	}
	if set {
		return timestamp, nil
	}
	return 0, nil
}

func UnlockRedisLock(key string, value int64) (bool, error) {
	luaScript := `
		if (redis.call('get', KEYS[1]) == ARGV[1]) then
			redis.call('del', KEYS[1]);
			return 1
		else
			return 0
		end
	`
	result, err := Client.Eval(luaScript, []string{key}, value).Result()
	if err != nil {
		return false, err
	}
	return result.(int64) == 1, err
}
