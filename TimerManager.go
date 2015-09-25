package TimerManager

import "time"

func CreateTimer() time.Time {
    startTime, _ := time.Parse(time.RFC822, "01 Jan 15 00:00 UTC")

    return startTime
}
