package datetime

import (
	"time"
)

type HappyDatetimeFormat string

const (
	Ymd_HMS     HappyDatetimeFormat = "2006-01-02 15:04:05"
	YmdHMS      HappyDatetimeFormat = "20060102150405"
	Ymd         HappyDatetimeFormat = "20060102"
	HMS         HappyDatetimeFormat = "150405"
	Y_m_d_H_M_S HappyDatetimeFormat = "2006_01_02_15_04_05"
)

func StrToDatetime(dateString string, customFormat HappyDatetimeFormat) (time.Time, error) {
	// 将时间字符串转换为datetime对象
	// dateString: 时间字符串，如2018-08-27 02:09:34
	// custom_format: 日期格式
	return time.Parse(string(customFormat), dateString)
}

func DatetimeToStr(datetimeObj time.Time, customFormat HappyDatetimeFormat) string {
	// 将datetime对象转换为指定格式的字符串
	// datetime_obj: 时间对象
	// custom_format: 日期格式
	return datetimeObj.Format(string(customFormat))
}

func GetCurrentDatetime(customFormat HappyDatetimeFormat) string {
	// 获取当前时间字符串
	// custom_format: 日期格式，默认 2018-09-03 13:36:02
	now := time.Now()
	return DatetimeToStr(now, customFormat)
}

func GetCurrentTimestamp() int64 {
	// 获取当前时间戳
	return time.Now().Unix()
}
