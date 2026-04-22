package main

type Segment string

const (
	Premium Segment = "PREMIUM"
	Normal  Segment = "NORMAL"
)

type User struct {
	userID      string
	name        string
	emailID     string
	userSegment Segment
}
