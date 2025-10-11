# 🧩 LLD Question — Design Snake and Ladder Game

## **Problem Statement**

Design and implement a **Snake and Ladder** game that can be played by multiple players. The game should simulate all the core features of the traditional board game.

Your design should be **object-oriented**, modular, and extensible — so that new features like custom boards, different dice types, or new rules can be added in the future with minimal changes.

---

## **Functional Requirements**

1. The game should have:

   * A **board** of configurable size (default: 100 cells).
   * A set of **snakes** — each defined by a head and a tail position.
   * A set of **ladders** — each defined by a start and end position.
   * A **dice** — rolls a random number between 1 and 6.
   * **Players** — each with a name and current position.

2. The game starts with all players at position 0.

3. On each player’s turn:

   * The player rolls the dice.
   * Moves forward by the number shown on the dice.
   * If the player lands at the **start of a ladder**, they climb up to the ladder’s end.
   * If the player lands at the **head of a snake**, they slide down to the snake’s tail.

4. The first player to reach the last cell (exactly) wins the game.

5. If the dice roll would move the player **beyond the last cell**, their position remains unchanged for that turn.

6. The system should display each move — player name, dice value, and new position.

---

## **Non-Functional Requirements**

* The design should be **clean, modular, and extensible**.
* The system should be **thread-safe** if multiple players are playing concurrently.
* Code should be easy to test and debug.
* Prefer clean abstractions for:

  * Player
  * Dice
  * Board
  * Snake
  * Ladder
  * Game engine / controller

---

## **Expected Output (Sample Simulation)**

```
Game started with 2 players: Avinash, Rahul

Avinash rolled a 4 → moved from 0 to 4
Rahul rolled a 6 → moved from 0 to 6
Avinash rolled a 3 → moved from 4 to 7
Rahul rolled a 2 → moved from 6 to 8
...
Avinash rolled a 2 → climbed ladder from 14 to 28
...
Rahul rolled a 5 → bitten by snake from 97 to 78
...
Avinash wins the game!
```

---

## **Follow-Up Questions (Optional)**

If time allows, the interviewer may extend the scope with:

1. Support for multiple dice (e.g., 2 dice instead of 1).
2. Add new types of cells like **Boost** (move forward 3) or **Trap** (skip next turn).
3. Allow the board configuration (snakes/ladders) to be read from a file.
4. Add persistence or replay capability.



Lets implement the base case first. 

Number of player and name must be entered before starting the game. 

It should be running automatically once we start the service. 

first few functions to just create the players and then it would keep on loggin till any of the player is reaching the end. 



Entities

Player {
   name string 
   position int
}

Board{
   size int 
   ladders map[int]int
   snakes map[int]int
}

GameManager{
   Players []*Player
   GameBoard *Board
   CurrentPlayer int 
}

AddPlayers
CreateBoard(n, ladders, snakes)

Start() // minimum two players required. 



