package main

Hotel {
    id string (PK)
    name string
    location string
    created_at timestamp
}

INDEX(location)

Room {
    id string (PK)
    hotel_id string (FK)

    room_number string
    room_type ENUM (SINGLE, DOUBLE, SUITE)

    capacity int
    base_price int

    created_at timestamp
}

INDEX(hotel_id)

Booking {
    id string (PK)

    user_id string
    hotel_id string

    total_amount int   // aggregated
    status ENUM (PENDING, CONFIRMED, CANCELLED)

    created_at timestamp
}

INDEX(user_id)
INDEX(hotel_id)

BookingRoom {
    id string (PK)

    booking_id string (FK)
    room_id string (FK)

    start_date timestamp
    end_date timestamp

    price int   // per-room price
}

INDEX(room_id, start_date, end_date)
INDEX(booking_id)

Payment {
    id string (PK)

    booking_id string (FK)

    amount int
    status ENUM (PENDING, SUCCESS, FAILED)

    payment_method string
    transaction_id string

    created_at timestamp
}

INDEX(booking_id)

SELECT * FROM BookingRoom
WHERE room_id = ?
AND NOT (end_date <= input_start OR start_date >= input_end)


RoomAvailability {
    room_id string
    date date

    available_count int   // or boolean for single room

    PRIMARY KEY (room_id, date)
}