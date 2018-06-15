package libs

import (
	"hangmango-web-api/config"
	"sync"
	"time"
)

const TIME_UNIT = 1 * time.Millisecond // 时间戳最小单位
const CURRENT_ID_MAX_BIT = 12
const DATA_ID_MAX_BIT = 5
const WORKER_ID_MAX_BIT = 5
const MAX_CURRENT_ID = 1<<CURRENT_ID_MAX_BIT - 1

type SnowFlake struct {
	Sm             sync.Mutex
	StartTimestamp int64 // 41位
	CurrentId      int64 // 12位  0 ~ 4095
	DataCenterId   int64
	WorkerId       int64
	LastTimestamp  int64
}

func (snowFlake *SnowFlake) ParseToTimestamp(times time.Time) int64 {
	return times.UnixNano() / int64(TIME_UNIT/time.Nanosecond)
}

func (snowFlake *SnowFlake) GenerateId() int64 {
	snowFlake.Sm.Lock()
	currentTimestamp := snowFlake.ParseToTimestamp(time.Now()) - snowFlake.StartTimestamp
	if snowFlake.LastTimestamp != currentTimestamp {
		snowFlake.LastTimestamp = currentTimestamp
		snowFlake.CurrentId = 0
	}
	if snowFlake.CurrentId > MAX_CURRENT_ID {
		time.Sleep(TIME_UNIT)
		snowFlake.Sm.Unlock()
		return snowFlake.GenerateId()
	}
	Id := currentTimestamp<<(CURRENT_ID_MAX_BIT+WORKER_ID_MAX_BIT+DATA_ID_MAX_BIT) |
		snowFlake.DataCenterId<<(CURRENT_ID_MAX_BIT+WORKER_ID_MAX_BIT) |
		snowFlake.WorkerId<<CURRENT_ID_MAX_BIT |
		snowFlake.CurrentId
	snowFlake.CurrentId++
	snowFlake.Sm.Unlock()
	return Id
}

func NewSnowflake() *SnowFlake {
	snowFlake := new(SnowFlake)
	snowFlake.StartTimestamp = snowFlake.ParseToTimestamp(config.Config.SnowflakeStartTime)
	snowFlake.DataCenterId = int64(config.Config.DataCenterId)
	snowFlake.WorkerId = int64(config.Config.WorkerId)
	return snowFlake
}
