package lib

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

type Snowflake struct {
	Sm             sync.Mutex
	StartTimestamp int64 // 41位
	CurrentId      int64 // 12位  0 ~ 4095
	DataCenterId   int64
	WorkerId       int64
	LastTimestamp  int64
}

func (snowflake *Snowflake) ParseToTimestamp(times time.Time) int64 {
	return times.UnixNano() / int64(TIME_UNIT/time.Nanosecond)
}

func (snowflake *Snowflake) GenerateId() int64 {
	snowflake.Sm.Lock()
	currentTimestamp := snowflake.ParseToTimestamp(time.Now()) - snowflake.StartTimestamp
	if snowflake.LastTimestamp != currentTimestamp {
		snowflake.LastTimestamp = currentTimestamp
		snowflake.CurrentId = 0
	}
	if snowflake.CurrentId > MAX_CURRENT_ID {
		time.Sleep(TIME_UNIT)
		snowflake.Sm.Unlock()
		return snowflake.GenerateId()
	}
	Id := currentTimestamp<<(CURRENT_ID_MAX_BIT+WORKER_ID_MAX_BIT+DATA_ID_MAX_BIT) |
		snowflake.DataCenterId<<(CURRENT_ID_MAX_BIT+WORKER_ID_MAX_BIT) |
		snowflake.WorkerId<<CURRENT_ID_MAX_BIT |
		snowflake.CurrentId
	snowflake.CurrentId++
	snowflake.Sm.Unlock()
	return Id
}

func NewSnowflake() *Snowflake {
	snowflake := new(Snowflake)
	snowflake.StartTimestamp = snowflake.ParseToTimestamp(config.Config.SnowflakeStartTime)
	snowflake.DataCenterId = int64(config.Config.DataCenterId)
	snowflake.WorkerId = int64(config.Config.WorkerId)
	return snowflake
}
