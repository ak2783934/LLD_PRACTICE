# Parking Lot System - Machine Coding Problem

## 🎯 Objective

Design and implement a Parking Lot System that can manage the entry and exit of vehicles in a multi-floor parking facility.

---

## 🏢 Parking Lot Structure

* The parking lot has **multiple floors**.
* Each floor contains a fixed number of **parking slots**.
* Each parking slot is designated for a **specific vehicle type**:

  * `TwoWheeler`
  * `FourWheeler`
  * `HeavyVehicle`

---

## ✅ Requirements

### 1. System Operations

You need to support the following commands:

| Command                                               | Description                                                                 |
| ----------------------------------------------------- | --------------------------------------------------------------------------- |
| `create_parking_lot numFloors slotsPerFloor`          | Initializes the parking lot with specified floors and slots on each floor.  |
| `park_vehicle vehicle_type registration_number color` | Parks a vehicle in the nearest available slot and returns a parking ticket. |
| `unpark_vehicle ticket_id`                            | Removes the vehicle from the slot and frees it.                             |
| `display_free_slots vehicle_type`                     | Shows the number of free slots for the given vehicle type on each floor.    |
| `display_occupied_slots vehicle_type`                 | Shows occupied slot details per floor.                                      |

---

## 🎟 Parking Ticket Format

Ticket ID should follow the format:

```
<parking_lot_id>_<floor_number>_<slot_number>
```

Ticket contains:

* Ticket ID
* Vehicle Registration Number
* Vehicle Color

---

## 📌 Parking Rules

* Vehicles should be parked in the **nearest available slot** – floor-wise, then slot-wise.
* One slot can hold only one vehicle at a time.
* Vehicle type must match the slot type.

---

## 🚗 Vehicle Types

| Vehicle Type | Examples        |
| ------------ | --------------- |
| TwoWheeler   | Bikes, Scooters |
| FourWheeler  | Cars, SUVs      |
| HeavyVehicle | Trucks, Buses   |

---

## 🧩 Assumptions (You Can Refine)

* No payment module initially (can be added later).
* System is in-memory (no database).
* A floor has fixed slots in order, like Slot 1, Slot 2, Slot 3...

---

## 🎤 What You Need To Do First

Before coding, you must:

1. Explain your understanding of the problem.
2. List your assumptions.
3. Ask any clarifying questions.

**Do not start coding immediately. This is part of the evaluation.**
