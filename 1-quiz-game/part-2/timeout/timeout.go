package timeout

import "time"

func Handler(timeout int, endChannel chan bool) {
	timer := time.NewTicker(time.Duration(timeout) * time.Second)
	<- timer.C
	endChannel <- true
}