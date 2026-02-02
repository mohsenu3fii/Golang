package main

import "time"

func CreateRateLimit(duration time.Duration) func(string) bool { // تغییر: bool به جای string
	type User struct {
		name      string
		vote_time time.Time
	}

	already_voted := make(map[string]*User)

	return func(username string) bool { // تغییر: bool به جای string
		user, exist := already_voted[username]
		now := time.Now()

		if !exist {
			already_voted[username] = &User{
				name:      username,
				vote_time: now,
			}
			return true
		}

		if now.Sub(user.vote_time) >= duration {
			user.vote_time = now
			return true
		}

		return false
	}
}
