package main

/*
🔴 Exercise 10: Pipeline pattern

Problem
Create pipeline:

Generator → Doubler → Printer


Each stage:

Runs in its own goroutine

Communicates via channel

Closes output channel properly

This tests:

Channel ownership

Close responsibility

Goroutine lifecycle

*/
