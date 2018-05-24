package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//if no args given, use "s" as default
	todo := "s"
	if len(os.Args) > 1 {
		todo = os.Args[1]
	}
	//switch first args
	switch todo {
	//m = stopwatch prints every 100ms
	case "m":
		milliseconds()
	//s = stopwatch prints every 1s
	case "s":
		seconds()
	}
}

func seconds() {
	startTime := time.Now().UTC()
	for {
		time.Sleep(1 * time.Second)
		currentTime := time.Since(startTime)
		fmt.Println(formatDuration(currentTime))
	}
}

func milliseconds() {
	startTime := time.Now().UTC()
	for {
		time.Sleep(100 * time.Millisecond)
		currentTime := time.Since(startTime)
		fmt.Println(formatDuration(currentTime))
	}
}

//format duration to a beauty output
func formatDuration(d time.Duration) string {
	//calculate values
	milliseconds := int64(d/time.Millisecond) % 1000
	seconds := int64(d.Seconds()) % 60
	minutes := int64(d.Minutes()) % 60
	hours := int64(d.Hours()) % 24
	var str string = ""
	//only add values if they are not 0
	if hours != 0 {
		str = fmt.Sprintf("%s%d%s", str, hours, " h ")
	}
	if minutes != 0 {
		str = fmt.Sprintf("%s%d%s", str, minutes, " min ")
	}
	if seconds != 0 {
		str = fmt.Sprintf("%s%d%s", str, seconds, " sec ")
	}
	if milliseconds != 0 {
		str = fmt.Sprintf("%s%d%s", str, milliseconds, " ms ")
	}
	//var str string = fmt.Sprintf("%d%s%d%s%d%s%d%s", hours, " h ", minutes, " min ", seconds, " sec ", milliseconds, " ms")
	return str
}
