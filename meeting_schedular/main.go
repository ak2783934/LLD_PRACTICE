package main

import (
	"errors"
	"time"
)

type MeetingRoom struct {
	MeetingRoomID string
	Name          string
	Bookings      []Booking
	Capacity      int
}

type BookingStatus int

const (
	BOOKED BookingStatus = iota
	CANCELLED
)

type Booking struct {
	BookingID            string
	UserID               string
	NumberOfParticipants int
	MeetingRoomID        string
	StartTime            time.Time
	EndTime              time.Time
	Status               string
}

type User struct {
	Name   string
	Email  string
	UserID string
}

type BookingManager struct {
	MeetingRooms map[string]MeetingRoom
	Bookings     map[string]Booking
}

// BookMeeting(meetingRoomID, startTime, endTime, numberOfParticipants)
// Cancel(BookingID)

func (b *BookingManager) BookMeeting(userID string, meetingRoomID string, startTime time.Time, endTime time.Time, numberOfParticipants int) (string, error) {
	meetingRoom, ok := b.MeetingRooms[meetingRoomID]
	if !ok {
		return "", errors.New("meeting room invalid")
	}

	if meetingRoom.Capacity < numberOfParticipants {
		return "", errors.New("meeting room doesn't have requested capcaity")
	}

	ExistingBookings := meetingRoom.Bookings

	for booking := range ExistingBookings {
		if booking.Status == "CANCELLED" {
			continue
		}

		if startTime.After(booking.StartTime) && endTime.Before(booking.StartTime) || startTime.After(booking.EndTime) && endTime.Before(booking.EndTime) {
			return "", errors.New("meeting room already booked")
		}
	}

	// Create a booking.
	booking := Booking{
		BookingID:            uuid(),
		UserID:               userID,
		MeetingRoomID:        meetingRoomID,
		NumberOfParticipants: numberOfParticipants,
		StartTime:            startTime,
		EndTime:              endTime,
		Status:               "BOOKED",
	}

	b.Bookings[booking.BookingID] = booking
	meetingRoom.Bookings = append(meetingRoom.Bookings, booking)

	return booking.BookingID, nil
}

// [2:30, 3:00]
// [12:00, 1:00]
// [1:00, 2:00]
// [5:00, 6:00]

func (b *BookingManager) Cancel(bookingID string) bool {
	booking, ok := b.Bookings[bookingID]
	if !ok {
		return false
	}

	booking.Status = "CANCELLED"
	b.Bookings[bookingID] = booking
	return true
}
