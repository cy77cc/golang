package main

import (
	"encoding/json"
	"fmt"
)

const WEATHER = `[
        {
          "condition": "多云",
          "date": "2021-07-14",
          "high_temperature": "36",
          "low_temperature": "29",
          "weather_icon_id": "1",
          "wind_direction": "南风",
          "wind_level": "3-4"
        },
        {
          "condition": "多云",
          "date": "2021-07-15",
          "high_temperature": "36",
          "low_temperature": "28",
          "weather_icon_id": "1",
          "wind_direction": "南风",
          "wind_level": "2"
        },
        {
          "condition": "多云",
          "date": "2021-07-16",
          "high_temperature": "36",
          "low_temperature": "28",
          "weather_icon_id": "1",
          "wind_direction": "南风",
          "wind_level": "1"
        },
        {
          "condition": "多云",
          "date": "2021-07-17",
          "high_temperature": "35",
          "low_temperature": "28",
          "weather_icon_id": "1",
          "wind_direction": "南风",
          "wind_level": "2"
        },
        {
          "condition": "多云",
          "date": "2021-07-18",
          "high_temperature": "35",
          "low_temperature": "27",
          "weather_icon_id": "1",
          "wind_direction": "南风",
          "wind_level": "3-4"
        },
        {
          "condition": "雷阵雨转多云",
          "date": "2021-07-19",
          "high_temperature": "34",
          "low_temperature": "27",
          "weather_icon_id": "4",
          "wind_direction": "南风",
          "wind_level": "1"
        },
        {
          "condition": "雷阵雨转多云",
          "date": "2021-07-20",
          "high_temperature": "34",
          "low_temperature": "27",
          "weather_icon_id": "4",
          "wind_direction": "南风",
          "wind_level": "2"
        },
        {
          "condition": "阴转中雨",
          "date": "2021-07-21",
          "high_temperature": "37",
          "low_temperature": "26",
          "weather_icon_id": "2",
          "wind_direction": "微风",
          "wind_level": "1"
        },
        {
          "condition": "小雨转多云",
          "date": "2021-07-22",
          "high_temperature": "38",
          "low_temperature": "27",
          "weather_icon_id": "7",
          "wind_direction": "东风",
          "wind_level": "2"
        },
        {
          "condition": "多云转晴",
          "date": "2021-07-23",
          "high_temperature": "38",
          "low_temperature": "27",
          "weather_icon_id": "1",
          "wind_direction": "东南风",
          "wind_level": "2"
        },
        {
          "condition": "小雨转阴",
          "date": "2021-07-24",
          "high_temperature": "38",
          "low_temperature": "28",
          "weather_icon_id": "7",
          "wind_direction": "东南风",
          "wind_level": "1"
        },
        {
          "condition": "阴",
          "date": "2021-07-25",
          "high_temperature": "38",
          "low_temperature": "29",
          "weather_icon_id": "2",
          "wind_direction": "西南风",
          "wind_level": "1"
        },
        {
          "condition": "阴",
          "date": "2021-07-26",
          "high_temperature": "38",
          "low_temperature": "28",
          "weather_icon_id": "2",
          "wind_direction": "西南风",
          "wind_level": "2"
        },
        {
          "condition": "阴",
          "date": "2021-07-27",
          "high_temperature": "38",
          "low_temperature": "28",
          "weather_icon_id": "2",
          "wind_direction": "东北风",
          "wind_level": "1"
        },
        {
          "condition": "多云转阴",
          "date": "2021-07-28",
          "high_temperature": "38",
          "low_temperature": "29",
          "weather_icon_id": "1",
          "wind_direction": "东北风",
          "wind_level": "2"
        },
        {
          "condition": "多云",
          "date": "2021-07-29",
          "high_temperature": "38",
          "low_temperature": "29",
          "weather_icon_id": "1",
          "wind_direction": "东风",
          "wind_level": "1"
        }
      ]`

type Weather struct {
	Condition       string `json:"condition"`
	Date            string `json:"date"`
	HighTemperature string `json:"high_temperature"`
	LowTemperature  string `json:"low_temperature"`
	WeatherIconId   string `json:"weather_icon_id"`
	WindDirection   string `json:"wind_direction"`
	WindLevel       string `json:"wind_level"`
}

func main() {
	weathers := make([]Weather, 0)
	err := json.Unmarshal([]byte(WEATHER), &weathers)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range weathers {
		fmt.Println(v)
	}
	//response, err := http.Get("https://www.toutiao.com/stream/widget/local_weather/data/?_signature=_02B4Z6wo00d01OD8qxwAAIDDUA1H1yofG1jg2K-AAFjWh1ux7lZo4K6uY2Yx34z3W3hXp8fFUIAq.7YZNZL1-rIIuQ9lCCMugwYQ-.YyNWPUNNhV9JOFrc-U72btT30qjloEWmDjVFcptGggea")
	//defer func() {
	//	_ = response.Body.Close()
	//}()
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	all, _ := ioutil.ReadAll(response.Body)
	//	fmt.Println(string(all))
	//}
}
