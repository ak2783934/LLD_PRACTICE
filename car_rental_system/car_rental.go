package main

import (
	"errors"
	"time"
)

type CarRentalSystem struct {
	Cars     map[int]*Car
	Bookings map[int]*Booking
	Users    map[int]*User
	Branches map[int]*Branch
	nextID   int
}

func NewCarRentalSystem() *CarRentalSystem {
	return &CarRentalSystem{
		Cars:     make(map[int]*Car),
		Bookings: make(map[int]*Booking),
		Users:    make(map[int]*User),
		Branches: make(map[int]*Branch),
		nextID:   0,
	}
}

/*
Functions
	AddBranch
	AddCar

	SearchAvailableCars
	IsCarAvailable(for one car for a given duration)
	BookCar
	CancelCar
	RegisterUser
*/

func (crs *CarRentalSystem) AddBranch(location, name string) int {
	crs.nextID++
	branch := &Branch{
		ID:       crs.nextID,
		Location: location,
		Name:     name,
	}

	crs.Branches[branch.ID] = branch
	return branch.ID
}

func (crs *CarRentalSystem) AddCar(regNo, brand, model, makingYear string, carType CarType, rate int, branchID int) int {
	crs.nextID++
	car := &Car{
		ID:         crs.nextID,
		RegNo:      regNo,
		Brand:      brand,
		Model:      model,
		CarType:    carType,
		MakingYear: makingYear,
		HourlyRate: rate,
		BranchID:   branchID,
	}

	crs.Cars[car.ID] = car
	return car.ID
}

func (crs *CarRentalSystem) RegisterUser(name, email, liscense string) int {
	crs.nextID++
	user := &User{
		ID:               crs.nextID,
		Name:             name,
		Email:            email,
		DiveringLiscense: liscense,
	}

	// can add validation for email check.
	crs.Users[user.ID] = user
	return user.ID
}

func (crs *CarRentalSystem) IsCarAvailable(carID int, startTime time.Time, endTime time.Time) bool {
	for _, booking := range crs.Bookings {
		if booking.CarID != carID {
			continue
		}
		if booking.IsCancelled {
			continue
		}

		// compare the time
		if booking.StartTime.Before(endTime) && booking.EndTime.After(startTime) {
			return false // overlapping booking exists
		}
	}
	return true
}

func (crs *CarRentalSystem) SearchAvailableCars(branchID int, cartype CarType, startTime time.Time, endTime time.Time) []*Car {
	// loop through all the cars under the given branch.
	var cars []*Car
	for _, car := range crs.Cars {
		if car.CarType != cartype {
			continue
		}

		if car.BranchID != branchID {
			continue
		}

		if crs.IsCarAvailable(car.ID, startTime, endTime) {
			cars = append(cars, car)
		}
	}
	return cars
}

func (crs *CarRentalSystem) BookCar(carID int, userID int, startTime, endTime time.Time) (int, error) {
	crs.nextID++
	if !crs.IsCarAvailable(carID, startTime, endTime) {
		return 0, errors.New("car is not available")
	}

	car, ok := crs.Cars[carID]
	if !ok {
		return 0, errors.New("car doesn't exist")
	}

	durationOfBookingInHours := endTime.Sub(startTime).Hours()
	cost := car.HourlyRate * int(durationOfBookingInHours)
	booking := &Booking{
		BookingID:   crs.nextID,
		UserID:      userID,
		CarID:       carID,
		StartTime:   startTime,
		EndTime:     endTime,
		TotalAmount: cost,
		IsCancelled: false,
	}
	crs.Bookings[booking.BookingID] = booking
	return booking.BookingID, nil
}

func (crs *CarRentalSystem) CancelBooking(bookingID int) error {
	booking, ok := crs.Bookings[bookingID]
	if !ok {
		return errors.New("booking doesn't exits")
	}

	booking.IsCancelled = true
	return nil
}
