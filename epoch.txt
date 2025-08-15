package main

import (
	"fmt"
	"time"
)

// EPOCH
// 00:00:00 UTC on Jan 1 1970

func main() {

	now := time.Now() // Curr. Time
	unixTime:=now.Unix() // seconds

	fmt.Println("Curr. Time:",now)
	fmt.Println("Unix Time/Value:",unixTime)

	// Human-readable form
	t:=time.Unix(unixTime,0)
	fmt.Println("⌛ Epoch time:",t) // 2025-08-15 19:48:49 +0530 IST
	fmt.Println("Time:",t.Format("2006-01-02")) // Time: 2025-08-15
}