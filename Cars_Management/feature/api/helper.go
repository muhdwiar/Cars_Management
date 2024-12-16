package api

import "time"

func Fail_Resp(msg string) map[string]interface{} {
	return map[string]interface{}{
		"Status":  "Failed",
		"Message": msg,
	}

}

func Succes_Resp() map[string]interface{} {
	return map[string]interface{}{
		"Status": "Success",
	}

}

func DateConvert(dateString string) (time.Time, error) {

	date, err := time.Parse("2006-01-02", dateString)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	newdate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, loc)
	return newdate, err
}
