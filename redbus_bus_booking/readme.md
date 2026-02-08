🚌 LLD Problem: Design a RedBus-like Bus Booking System
📌 Context (what interviewer says)

Design a simplified bus ticket booking system similar to RedBus.

You are expected to design core domain classes and logic, not UI or infra.

1️⃣ Functional Requirements (MUST)
1. Bus & Route Management

A Bus runs on a specific Route

A Route has:

Source city

Destination city

Intermediate stops (optional)

Each bus has:

Bus ID

Operator name

Bus type (Sleeper / Seater / Semi-Sleeper)

Total seats

Seat layout (row, column, seat type)

2. Seat & Availability

Seats can be:

Window / Aisle / Sleeper

Seats can be:

Available

Temporarily held

Booked

Seat availability is date-specific

3. Search

Users should be able to:

Search buses by:

Source

Destination

Date

View:

Bus details

Departure / Arrival time

Available seats

4. Booking Flow

User selects:

Bus

Date

Seats

Seats should be locked temporarily

Booking can be:

Confirmed

Expired (if not confirmed in time)

On confirmation:

Seats become booked

Ticket is generated

5. Ticket

Ticket should contain:

Ticket ID

User ID

Bus ID

Date

Seat numbers

Booking status (CONFIRMED / CANCELLED)

2️⃣ Core Rules & Constraints

One seat can be booked by only one user for a given date

Temporary seat lock expires after X minutes

Multiple users can search simultaneously

Booking should prevent double booking


3️⃣ Non-Functional Requirements (IMPORTANT)

Focus on:

Clean object modeling

Extensibility

Testability

In-memory implementation is sufficient

Thread safety:

Not required to implement

But must be discussed

No persistence / DB required

No payment integration (out of scope)







