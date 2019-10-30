package main

import (
	"time"

	"github.com/astaxie/beego/logs"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func parseWithLocation(name string, timeStr string) (time.Time, error) {
	locationName := name
	if l, err := time.LoadLocation(locationName); err != nil {
		logs.Debug(err.Error())
		return time.Time{}, err
	} else {
		lt, _ := time.ParseInLocation(TIME_LAYOUT, timeStr, l)
		name, offset := lt.Zone()
		logs.Debug(locationName, "\t", lt, "\t", name, "\t", offset/3600)
		return lt, nil
	}
}
func main() {
	logs.Debug("0. now: ", time.Now())
	str := "2018-09-10 00:00:00"
	// str := time.Now()
	logs.Debug("1. str: ", str)
	t, _ := time.Parse(TIME_LAYOUT, str)
	logs.Debug("2. Parse time: ", t)
	tStr := time.Now().Format(TIME_LAYOUT)
	logs.Debug("3. Format time str: ", tStr)
	name, offset := t.Zone()
	name2, offset2 := t.Local().Zone()
	logs.Debug("4. Zone name: %v, Zone offset: %v\n", name, offset)
	logs.Debug("5. Local Zone name: %v, Local Zone offset: %v\n", name2, offset2)
	tLocal := t.Local()
	tUTC := t.UTC()
	logs.Debug("6. t: %v, Local: %v, UTC: %v\n", t, tLocal, tUTC)
	logs.Debug("7. t: %v, Local: %v, UTC: %v\n", t.Format(TIME_LAYOUT), tLocal.Format(TIME_LAYOUT), tUTC.Format(TIME_LAYOUT))
	logs.Debug("8. Local.Unix: %v, UTC.Unix: %v\n", tLocal.Unix(), tUTC.Unix())
	str2 := "1969-12-31 23:59:59"
	t2, _ := time.Parse(TIME_LAYOUT, str2)
	logs.Debug("9. str2：%v，time: %v, Unix: %v\n", str2, t2, t2.Unix())
	logs.Debug("10. %v, %v\n", tLocal.Format(time.ANSIC), tUTC.Format(time.ANSIC))
	logs.Debug("11. %v, %v\n", tLocal.Format(time.RFC822), tUTC.Format(time.RFC822))
	logs.Debug("12. %v, %v\n", tLocal.Format(time.RFC822Z), tUTC.Format(time.RFC822Z))

	//指定时区
	parseWithLocation("America/Cordoba", tStr)
	parseWithLocation("Asia/Shanghai", tStr)
	parseWithLocation("Asia/Japan", tStr)

	tm := time.Now()
	logs.Debug(tm.String())
	logs.Debug(tm.Format(TIME_LAYOUT))

}
