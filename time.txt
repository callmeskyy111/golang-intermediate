package main

import (
	"fmt"
	"time"
)

func main() {

	// Curr. local time
	fmt.Println(time.Now())

	// Specific time
	spicificTime:= time.Date(2025, time.August,30,12,0,0,0,time.UTC)
	fmt.Println("🌏 Specific Time:",spicificTime)

	// Parse Time
	parsedTime,err:=time.Parse("2006-01-02","2020-05-01") //2020-05-01 00:00:00 +0000 UTC
	parsedTime1,err:=time.Parse("06-01-02","20-05-01") // "
	parsedTime2,err:=time.Parse("06-1-2","20-5-1") // "
	parsedTime3,err:=time.Parse("06-1-2 15-04","20-5-1 18-03") //2020-05-01 18:03:00 +0000 UTC

	if err!=nil{
		fmt.Println("🔴ERR:",err)
	}
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)
	fmt.Println(parsedTime3)

	// Formatting time
	t := time.Now()
	fmt.Println("Formatted Time:",t.Format("06-01-02 04-15-05"))

	oneDayAdded:= t.Add(time.Hour * 24)
	fmt.Println(oneDayAdded)
	fmt.Println(oneDayAdded.Weekday())

	fmt.Println("Rounded Time:",t.Round(time.Hour))

	loc, _:=time.LoadLocation("Asia/Kolkata")
	t = time.Date(2025, time.August, 16, 8, 19,5,00, time.UTC )

	// Convert the UTC to specific time-zone
	tLocal:= t.In(loc)
	
	// Rounding
	roundedT:= t.Round(time.Hour)
	roundedTLocal:= roundedT.In(loc)

	fmt.Println("Original Time (UTC):",t)
	fmt.Println("Original Time (Local):",tLocal)
	fmt.Println("Rounded Time (UTC):",roundedT)
	fmt.Println("Rounded Time (Local):",roundedTLocal)

	fmt.Println("Truncated Time: ",t.Truncate(time.Hour))

	// Handling timezones
	loc1,_:= time.LoadLocation("America/New_York")

	// convert the time -> location
	tInNY:= time.Now().In(loc1)
	fmt.Println("New York Time:",tInNY) //New York Time: 2025-08-15 09:59:14.8998227 -0400 EDT

	t1:=time.Date(2025, time.August,4,12,0,0,0,time.UTC)
	t2:= time.Date(2023, time.September,4,18,0,0,0, time.UTC)
	duration:=t1.Sub(t2)
	fmt.Println("Duration",duration)

	// Compare time
	fmt.Println("t1 is before t2?",t1.Before(t2)) // false



}