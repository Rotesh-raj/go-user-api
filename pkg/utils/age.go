package utils

import "time"

func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}
	
	if age < 0 {
		return 0
	}
	return age
}
