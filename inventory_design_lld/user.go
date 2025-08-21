package main

type UserRole int

const (
	BUYER UserRole = iota
	SELLER
)

type User struct {
	UserID string
	Name   string
	Role   UserRole
}
