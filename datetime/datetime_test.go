package datetime

import (
	"regexp"
	"testing"
)

func TestStrToDatetime(t *testing.T) {
	dt, err := StrToDatetime("2020-04-19 20:04:14", Ymd_HMS)
	if err != nil {
		t.Fatalf("Failed to parse datetime: %v", err)
	}

	if dt.Year() != 2020 {
		t.Errorf("Expected year 2020, got %d", dt.Year())
	}
	if dt.Month() != 4 {
		t.Errorf("Expected month 4, got %d", dt.Month())
	}
	if dt.Day() != 19 {
		t.Errorf("Expected day 19, got %d", dt.Day())
	}
	if dt.Hour() != 20 {
		t.Errorf("Expected hour 20, got %d", dt.Hour())
	}
	if dt.Minute() != 4 {
		t.Errorf("Expected minute 4, got %d", dt.Minute())
	}
	if dt.Second() != 14 {
		t.Errorf("Expected second 14, got %d", dt.Second())
	}
}

func TestDatetimeToStr(t *testing.T) {
	dtStr := "2020-04-19 20:04:14"
	dt, err := StrToDatetime(dtStr, Ymd_HMS)
	if err != nil {
		t.Fatalf("Failed to parse datetime: %v", err)
	}

	str := DatetimeToStr(dt, Ymd_HMS)
	if str != dtStr {
		t.Errorf("Expected datetime string %s, got %s", dtStr, str)
	}
}

func TestGetCurrentDatetime(t *testing.T) {
	formats := map[HappyDatetimeFormat]*regexp.Regexp{
		Ymd_HMS:     regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`),
		Y_m_d_H_M_S: regexp.MustCompile(`\d{4}_\d{2}_\d{2}_\d{2}_\d{2}_\d{2}`),
		YmdHMS:      regexp.MustCompile(`\d{4}\d{2}\d{2}\d{2}\d{2}\d{2}`),
		Ymd:         regexp.MustCompile(`\d{4}\d{2}\d{2}`),
		HMS:         regexp.MustCompile(`\d{2}\d{2}\d{2}`),
	}

	for format, re := range formats {
		value := GetCurrentDatetime(format)
		if !re.MatchString(value) {
			t.Errorf("Value %s does not match format %s", value, format)
		}
	}
}

func TestGetCurrentTimestamp(t *testing.T) {
	value := GetCurrentTimestamp()

	if value <= 0 {
		t.Errorf("Expected positive timestamp, got %d", value)
	}
}

func TestGetCurrentTimestampV2(t *testing.T) {
	value := GetCurrentTimestamp()

	if value <= 0 {
		t.Errorf("Expected positive timestamp, got %d", value)
	}
}
