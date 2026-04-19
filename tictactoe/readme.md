Requirements
    Functional
        NxN board (default 3x3) with two players
        Players alternate turns placing X and O
        Detect win (row, column, or diagonal) and draw conditions
        Support both human and AI players via pluggable move strategies
        Reject out-of-bounds and occupied cell moves
    Non-Functional
        O(1) win detection per move
        Game state transitions should be explicit and testable
        Board size should be configurable without code changes



[- - x - -]
[- x - - -]
[x - - - -]
[- - - - -]
[- - - - -]


