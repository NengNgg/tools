package utils

import (
	"fmt"
	timeconv "github.com/Andrew-M-C/go.timeconv"
	"math"
	"strconv"
	"strings"
	"time"
)

func TimeHour(time2 time.Time) time.Time {
	return time.Date(time2.Year(), time2.Month(), time2.Day(), time2.Hour(), 0, 0, 0, time2.Location())
}

func TimeLastHour(time2 time.Time) time.Time {
	return time.Date(time2.Year(), time2.Month(), time2.Day(), time2.Hour()-1, 0, 0, 0, time2.Location())
}

func TimeToString(time2 time.Time) string {
	return time.Date(time2.Year(), time2.Month(), time2.Day(), time2.Hour(), time2.Minute(), time2.Second(), 0, time2.Location()).
		Format("2006-01-02 15:04:05")
}

// TimeDay 获取给定时间当天0点0分0秒的时间
func TimeDay(time2 time.Time) time.Time {
	return time.Date(time2.Year(), time2.Month(), time2.Day(), 0, 0, 0, 0, time2.Location())
}

// TimeDayFinalTime 获取给定时间当天23点59分59秒的时间
func TimeDayFinalTime(time2 time.Time) time.Time {
	return time.Date(time2.Year(), time2.Month(), time2.Day(), 23, 59, 59, 0, time2.Location())
}

func TimeStampDay(time2 time.Time) int64 {
	return time.Date(time2.Year(), time2.Month(), time2.Day(), 0, 0, 0, 0, time2.Location()).UnixMilli()
}

func TimeLastDay(time2 time.Time) time.Time {
	return time.Date(time2.Year(), time2.Month(), time2.Day()-1, 0, 0, 0, 0, time2.Location())
}

func TimeMonth(time2 time.Time) time.Time {
	return time.Date(time2.Year(), time2.Month(), 1, 0, 0, 0, 0, time2.Location())
}

func TimeLastMonth(time2 time.Time) time.Time {
	return time.Date(time2.Year(), time2.Month()-1, 1, 0, 0, 0, 0, time2.Location())
}

// TimeYearFirstDay 获取给定时间当年1月1号0点00分0秒的时间
func TimeYearFirstDay(time2 time.Time) time.Time {
	return time.Date(time2.Year(), 1, 1, 0, 0, 0, 0, time2.Location())
}

// TimeLastFewMonth monthNum 个月前
func TimeLastFewMonth(time2 time.Time, monthNum int) time.Time {
	time1 := timeconv.AddDate(time2, 0, -monthNum, 0)
	return time.Date(time1.Year(), time1.Month(), 1, 0, 0, 0, 0, time1.Location())
}

// TimeLastFewDay Num 个天前
func TimeLastFewDay(time2 time.Time, Num int) time.Time {
	time1 := time2.AddDate(0, 0, -Num)
	return time.Date(time1.Year(), time1.Month(), time1.Day(), 0, 0, 0, 0, time1.Location())
}

func TimeMonDay(time2 time.Time) string {
	return time2.Format("01-02")
}

func TimeFormatYear(time2 time.Time) string {
	return time2.Format("2006")
}

func TimeFormatSeason(time2 time.Time) string {
	switch time2.Month() {
	case 1, 2, 3:
		return time2.Format("2006") + "Q1"
	case 4, 5, 6:
		return time2.Format("2006") + "Q2"
	case 7, 8, 9:
		return time2.Format("2006") + "Q3"
	case 10, 11, 12:
		return time2.Format("2006") + "Q4"
	}
	return ""
}

func TimeFormatMonth(time2 time.Time) string {
	return time2.Format("2006-01")
}

func TimeFormatDay(time2 time.Time) string {
	return time2.Format("2006-01-02")
}

func TimeFormat(time2 time.Time) string {
	return time2.Format("2006-01-02 15:04:05")
}

func TimeParse(data string) time.Time {
	return TimeParseString(data, "2006-01-02T15:04:05Z")
}
func TimeParseCST(data string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", data, time.FixedZone("UTC+8", 8*60*60))
	if err != nil {
		return time.UnixMilli(0)
	}
	return t
}

func TimeParse1(data string) time.Time {
	return TimeParseString(data, "2006-01-02T15:04:05")
}
func TimeParse2(data string) time.Time {
	return TimeParseString(data, "2006-01-02 15:04:05")
}

func TimeParse4(data string) time.Time {
	data = data + " +0000 UTC"
	return TimeParseString(data, "2006-01-02 15:04:05 +0000 UTC")
}

func TimeParseP8(data string) time.Time {
	data += ":00:00Z"
	return TimeParse(data)
}

func TimeParseP10(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.UnixMilli(0)
	}
	return t
}

// UTCTransTimeStamp UTC时间转标准时间
func UTCTransTimeStamp(utcTime string) time.Time {
	t := TimeParseString(utcTime, "2006-01-02 15:04:05 +0000 UTC")
	return t
}

func TimeParseString(data string, format string) time.Time {
	t, err := time.Parse(format, data)
	if err != nil {
		return time.UnixMilli(0)
	}
	return t
}

func ParseStringToTime(timeStr string) time.Time {
	timeLay := "2006-01-02 15:04:05"
	timeParsed, err := time.ParseInLocation(timeLay, timeStr, time.Local)
	if err != nil {
		return time.UnixMilli(0)
	}
	return timeParsed
}

func Timestamp2Time(timestamp int64) time.Time {
	return time.UnixMilli(timestamp)
}

func Timestamp3Time(timestamp int64) string {
	return TimeFormat(time.UnixMilli(timestamp))
}

func WeekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}

func WeekRange(year, week int) (start, end time.Time) {
	start = WeekStart(year, week)
	end = start.AddDate(0, 0, 6)
	return
}

// GetMonthStartAndEnd 获取月份的第一天和最后一天
func GetMonthStartAndEnd(myYear string, myMonth string) (time.Time, time.Time) {
	// 数字月份必须前置补零
	if len(myMonth) == 1 {
		myMonth = "0" + myMonth
	}
	yInt, _ := strconv.Atoi(myYear)

	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, myYear+"-"+myMonth+"-01 00:00:00", loc)
	newMonth := theTime.Month()

	startTime := time.Date(yInt, newMonth, 1, 0, 0, 0, 0, time.Local)
	endTime := time.Date(yInt, newMonth+1, 0, 0, 0, 0, 0, time.Local)
	return startTime, endTime
}

// MonthToTime 补全时间,且把时区转为和入参time一致
func MonthToTime(timeStr string) time.Time {
	endTimeStr := timeStr + "-01 00:00:00"
	monthTime, err := time.Parse("2006-01-02 15:04:05", endTimeStr)
	if err != nil {
		return time.Now()
	}
	return monthTime
}

// MonthToTimeAndLocal 补全时间,且把时区转为和入参time一致  2012-12 --> 2012-12-01 00:00:00
func MonthToTimeAndLocal(timeStr string, timeLocal time.Time) time.Time {
	endTimeStr := timeStr + "-01 00:00:00"
	monthTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, timeLocal.Location())
	if err != nil {
		return time.Now()
	}
	return monthTime
}

// MonthToYearStartAndEnd 补全时间
func MonthToYearStartAndEnd(timeStr string) (time.Time, time.Time) {
	startTimeStr := timeStr + "-01-01 00:00:00"
	startTime, err := time.Parse("2006-01-02 15:04:05", startTimeStr)
	if err != nil {
		return time.Now(), time.Now()
	}

	endTimeStr := timeStr + "-12-31 23:59:59"
	endTime, err := time.Parse("2006-01-02 15:04:05", endTimeStr)
	if err != nil {
		return time.Now(), time.Now()
	}
	return startTime, endTime
}

// StringMatchTime 根据字符串长度匹配时间
func StringMatchTime(timeStr string, timeLocal time.Time) time.Time {
	endTimeStr := ""
	switch len(timeStr) {
	case 4:
		endTimeStr = timeStr + "-01-01 00:00:00"
	case 6:
		//Q1-Q4
		endStr := timeStr[4:6]
		switch endStr {
		case "Q1":
			endTimeStr = timeStr[0:4] + "-01-01 00:00:00"
		case "Q2":
			endTimeStr = timeStr[0:4] + "-04-01 00:00:00"
		case "Q3":
			endTimeStr = timeStr[0:4] + "-07-01 00:00:00"
		case "Q4":
			endTimeStr = timeStr[0:4] + "-10-01 00:00:00"
		}
	case 7:
		endTimeStr = timeStr + "-01 00:00:00"
	case 10:
		endTimeStr = timeStr + " 00:00:00"
	case 13:
		endTimeStr = timeStr + ":00:00"
	case 19:
		endTimeStr = timeStr
	}

	monthTime, err := time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, timeLocal.Location())
	if err != nil {
		return time.Now()
	}
	return monthTime
}

// GetThisMonthStartEnd 获取当前时间的这个月
func GetThisMonthStartEnd(now time.Time) (time.Time, time.Time) {
	thisMonthFirstDay := now.AddDate(0, 0, -now.Day()+1)
	thisMonthStart := time.Date(thisMonthFirstDay.Year(), thisMonthFirstDay.Month(), thisMonthFirstDay.Day(), 0, 0, 0, 0, now.Location())
	thisMonthEndDay := thisMonthFirstDay.AddDate(0, 1, -1)
	thisMonthEnd := time.Date(thisMonthEndDay.Year(), thisMonthEndDay.Month(), thisMonthEndDay.Day(), 23, 59, 59, 0, now.Location())
	return thisMonthStart, thisMonthEnd
}

// GetFirstDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return TimeDay(d)
}

// GetLastMonthStartEnd 获取当前时间的上个月
func GetLastMonthStartEnd(now time.Time) (time.Time, time.Time) {
	lastMonthFirstDay := now.AddDate(0, -1, -now.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, now.Location())
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, now.Location())
	return lastMonthStart, lastMonthEnd
}

// MonthToSeason 2022-06处理成季度
func MonthToSeason(yearMonth string) string {
	switch yearMonth[len(yearMonth)-1:] {
	case "1":
		return yearMonth[:len(yearMonth)-1] + "Q1"
	case "2":
		return yearMonth[:len(yearMonth)-1] + "Q2"
	case "3":
		return yearMonth[:len(yearMonth)-1] + "Q3"
	case "4":
		return yearMonth[:len(yearMonth)-1] + "Q4"
	}
	return ""
}

// GetYearStartAndEnd 获取某年第一天和最后一天时间
func GetYearStartAndEnd(myYear string) (time.Time, time.Time) {
	startTime, _ := GetMonthStartAndEnd(myYear, "1")
	_, endTime := GetMonthStartAndEnd(myYear, "12")
	return startTime, endTime
}

func GetDayStartAndEnd(timeStr string) (time.Time, time.Time) {
	startTimeStr := timeStr + " 00:00:00"
	endTimeStr := timeStr + " 23:59:59"
	startTime, err := time.Parse("2006-01-02 15:04:05", startTimeStr)
	if err != nil {
		return time.Now(), time.Now()
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", endTimeStr)
	if err != nil {
		return time.Now(), time.Now()
	}
	return startTime, endTime
}

func CompareMonth(cycle1, cycle2 string) int { // 2006-10
	s := strings.Split(cycle1, "-")
	year1, _ := strconv.Atoi(s[0])
	month1, _ := strconv.Atoi(s[1])

	s = strings.Split(cycle2, "-")
	year2, _ := strconv.Atoi(s[0])
	month2, _ := strconv.Atoi(s[1])

	if year1 > year2 {
		return 1
	} else if year1 < year2 {
		return -1
	}

	if month1 > month2 {
		return 1
	} else if month1 < month2 {
		return -1
	} else {
		return 0
	}
}

func AddMonth(cycle string) string {
	s := strings.Split(cycle, "-")
	year, _ := strconv.Atoi(s[0])
	month, _ := strconv.Atoi(s[1])

	if month+1 > 12 {
		return fmt.Sprintf("%d-01", year+1)
	} else {
		if month > 8 {
			return fmt.Sprintf("%d-%2d", year, month+1)
		} else {
			return fmt.Sprintf("%d-0%d", year, month+1)
		}

	}
}

func GetPrevMontyStEndTime(t time.Time) (time.Time, time.Time) {
	lastMonthFirstDay := t.AddDate(0, -1, -t.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, t.Location())
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, t.Location())
	return lastMonthStart, lastMonthEnd
}

// GetYearMonthToDay 查询指定年份指定月份有多少天
func GetYearMonthToDay(year int, m time.Month) int {
	// 有31天的月份
	month := int(m)
	day31 := map[int]struct{}{
		1:  {},
		3:  {},
		5:  {},
		7:  {},
		8:  {},
		10: {},
		12: {},
	}
	if _, ok := day31[month]; ok {
		return 31
	}
	// 有30天的月份
	day30 := map[int]struct{}{
		4:  {},
		6:  {},
		9:  {},
		11: {},
	}
	if _, ok := day30[month]; ok {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}

// GetTimeMonthDays 查询time的这个月有多少天
func GetTimeMonthDays(t time.Time) int {
	return GetYearMonthToDay(t.Year(), t.Month())
}

func GetWeekEnd(t time.Time) time.Time {
	now := t
	offset := int(7 - now.Weekday())
	if offset == 7 {
		offset = 0
	}

	weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Local).AddDate(0, 0, offset)
	return weekEnd
}

// TampDifferenceToString 把时间戳转为天数，结果格式为 1天
func TampDifferenceToString(t int64) string {
	return strconv.FormatInt(t/(time.Hour.Milliseconds()*24), 10) + "天"
}

func TimeZeroPoint(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func TimeZeroNsecPoint(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}

func TampZeroNsecPoint(t int64) (res int64) {
	resT := TimeZeroNsecPoint(time.UnixMilli(t))
	return resT.UnixMilli()
}

func TimeZeroHourPoint(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), 0, 0, 0, t.Location())
}

func TimeToSec(t time.Time) string {
	layout := "2006-01-02 15:04:05" // 添加日期的格式
	return t.Format(layout)
}

func TampZeroHourPoint(tamp int64) int64 {
	t := time.UnixMilli(tamp)
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), 0, 0, 0, t.Location()).UnixMilli()
}

func TampToTimeZeroHourPoint(tamp int64) time.Time {
	t := time.UnixMilli(tamp)
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), 0, 0, 0, t.Location())
}

func MilSecToDurationString(t int64) string {
	sec := t / 1000 % 60
	min := t / 1000 / 60 % 60
	hour := t / 1000 / 60 / 60 % 24
	day := t / 1000 / 60 / 60 / 24

	s := ""
	if day > 0 {
		s += fmt.Sprintf("%dd", day)
	}
	if hour > 0 {
		s += fmt.Sprintf("%dh", hour)
	}
	if min > 0 {
		s += fmt.Sprintf("%dm", min)
	}
	if sec > 0 {
		s += fmt.Sprintf("%ds", sec)
	}

	return s
}

func GetDayLastTime(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

func TimestampToDay(timestamp int64) int64 {
	dateTime := time.UnixMilli(timestamp)
	return time.Date(
		dateTime.Year(),
		dateTime.Month(),
		dateTime.Day(),
		0, 0, 0, 0, dateTime.Location(),
	).UnixMilli()
}

func TimestampToMonth(timestamp int64) int64 {
	dateTime := time.UnixMilli(timestamp)
	return time.Date(
		dateTime.Year(),
		dateTime.Month(),
		1,
		0, 0, 0, 0, dateTime.Location(),
	).UnixMilli()
}

func TimestampToSeason(timestamp int64) int64 {
	dateTime := time.UnixMilli(timestamp)
	month := time.Month(1)
	switch dateTime.Month() {
	case 1, 2, 3:
		month = 1
	case 4, 5, 6:
		month = 4
	case 7, 8, 9:
		month = 7
	case 10, 11, 12:
		month = 10
	}
	return time.Date(
		dateTime.Year(),
		month,
		1,
		0, 0, 0, 0, dateTime.Location(),
	).UnixMilli()
}

func TimestampToYear(timestamp int64) int64 {
	dateTime := time.UnixMilli(timestamp)
	return time.Date(
		dateTime.Year(),
		1,
		1,
		0, 0, 0, 0, dateTime.Location(),
	).UnixMilli()
}

func TimeDifferenceIsOneSecond(t1, t2 int64) bool {
	return math.Abs(float64(t1-t2)) <= float64(time.Second.Milliseconds())
}

func StartEndTampToCycles(start, end int64) []string {
	if end < start {
		return nil
	}

	if end == start {
		return []string{TimeFormatMonth(time.UnixMilli(start))}
	}

	startT, endT := TimeMonth(time.UnixMilli(start)), TimeMonth(time.UnixMilli(end))
	var res []string
	for startT.Before(endT) {
		res = append(res, TimeFormatMonth(startT))
		startT = startT.AddDate(0, 1, 0)

		if startT == endT {
			res = append(res, TimeFormatMonth(startT))
		}
	}
	return res
}

func SplitMonthToTime(Month string) time.Time {
	s := strings.Split(Month, "-")
	if len(s) == 2 {
		year1, _ := strconv.Atoi(s[0])
		month1, _ := strconv.Atoi(s[1])
		return time.Date(year1, time.Month(month1), 1, 0, 0, 0, 0, time.Local)
	}
	return time.Time{}
}
