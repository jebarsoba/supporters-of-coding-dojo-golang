package TimerManager

import ("time"; "testing")

func TestTimerManager(t *testing.T) {
    expectedStartTime, _ := time.Parse(time.RFC822, "01 Jan 15 00:00 UTC")

    startTime := CreateTimer();

    if (startTime != expectedStartTime) {
      t.Fatalf("Expected %d, actual %d.", expectedStartTime, startTime)
    }
}
