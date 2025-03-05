package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

/*const timeLayoutDate = "2006-01-02"
const timeLayoutTime = "15:04:05"*/

const timezone = "Asia/Shanghai"

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeLayout)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeLayout)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	now, err := time.ParseInLocation(`"`+timeLayout+`"`, string(data), time.Local)
	*t = Time(now)
	return err
}

func (t Time) String() string {
	return time.Time(t).Format(timeLayout)
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil

}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
