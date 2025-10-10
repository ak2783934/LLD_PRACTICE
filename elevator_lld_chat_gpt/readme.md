🚪 Requirements

Your elevator should:
Start at floor 0 in an Idle state.
Handle requests like:
    CallElevator(floor int)
    SelectFloor(floor int)

Have distinct states that represent the elevator’s mode of operation:

Log transitions like:
    “Elevator moving up to floor 5” 
    “Elevator doors opening at floor 5”
    “Elevator idle at floor 5”

Should handle multiple floor requests in a queue.

Should be extensible — for example, adding maintenance mode or emergency stop later.

⚙️ Expectations
    Use interfaces to represent states (e.g., ElevatorState interface with methods like RequestFloor, Arrive, etc.)
    Each state should define its own behavior — no big switch on state type.
    Use composition, not inheritance.
    Keep it simple but clean — don’t simulate time delays.
    Implement in Golang (idiomatic design expected).
    No need for concurrency for now (we’ll add that later if you want).




Analysis of question and entites defination 
And member functions. 


Assumption:
- There is only one elevator that we are controlling. 
- And if we have multiple elevators, each will have its own controlling mechanism. 
- incase of multiple requests, we complete the one round up and then come one round down? 
    - for one elevator, I think this is the best way, only change the direction if you have not got any request in the coming directions. IMP. 
    - Two type of requests, first request that came after entering the lift, but the lift must continue in its own direction if there is some already loaded request? 


So basically the user can come in and request and if it has a direction right now, serve all request in that direction and then serve request in other direction. simple. We might need to use priority queue for that case. 





Entities: 
Request{
    floor int
}
Elevator{
    CurrentFloor int
    UpQueue []int --> sorted in ascending order
    DownQueue []int --> sorted in descending order 
    ElevatorState ElevatorState
}

ElevatorState {
    Request(floor int)
    MoveUp()
    MoveDown()
    OpenDoor()
    CloseDoor()
}







