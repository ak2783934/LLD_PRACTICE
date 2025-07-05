package main

import "time"

type Booking struct {
	BookingID   int
	UserID      int
	CarID       int
	StartTime   time.Time
	EndTime     time.Time
	IsCancelled bool
	TotalAmount int
}
