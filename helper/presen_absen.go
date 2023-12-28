package helper

import "time"

func TimeNow() time.Time {
	return time.Now()
}

func TimeHM() (int, int) {
	time := TimeNow()
	h, m := time.Hour(), time.Minute()

	return h, m
}

func PresenceTime() bool {
	h, m := TimeHM()
	if h <= 9 && m == 0 {
		return false
	} else {
		return true
	}
}

func AbsenceTime() bool {
	h, m := TimeHM()
	if h >= 17 && m >= 0 {
		return false
	} else {
		return true
	}
}
