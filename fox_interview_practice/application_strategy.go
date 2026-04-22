package main

type ApplicationStrategy interface {
	isApplicable() bool
}
