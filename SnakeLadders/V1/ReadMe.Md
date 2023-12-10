Snake and Ladders
================================================
Step 1: Requirements
================================================

1) There is a game.
2) Any number of players can play the game.
3) Game has a board of 10x10 (100 cells) dimensions.
4) A block or cell consistes with a number.
5) A number with a snake has a destination number which is less than curr position.
6) A number with a ladder has a destination number which is greater than curr position.
7) A dice is rolled out to know how many steps player has to move. A move is deined as the number of number of steps player will move.
8) Game starts from block 1 and The player who reached 100th  block first, wins.
9) A random number is displayed on dice from 1-6.
10) A board has number in zig zag fashion like first 1-10 (from left to right) and then 11-20 (right to left) and so on.

================================================================
Step 2: Find entities
================================================================
Game, Player,  Board, Cell, Snake, Ladders, Dice

1) There is a game.
Game{
    List<Player>
}

2) Any number of players can play the game.
Relation b/w game and players.
GamePlayers{
    gameid
    playerid
}

3) Game has a board of 10x10 (100 blocks) dimensions.
Relation b/w game and board.

Game{
    List<Player>
    Board
}

Board{
    List<List<Cell>>
}

GameBoard{
    gameid
    boardid
}


4) A block or cell consistes with a number.
Cell{
    number
}

5) A number with a snake has a destination number which is less than curr position.
6) A number with a ladder has a destination number which is greater than curr position.

Thinking to handle this by storing start and end in a map.

7) A dice is rolled out to know how many steps player has to move. A move is deined as the number of number of steps player will move.

Generate a random number at run time

- Dice{
    Roll() int
}

================================================================
Step 3: Draw Class Diagram
1) Consider only those entites which has attributes defined.
2) Only those relation will be considered, who has some attributes defined.
================================================================

Game{
    List<Player> players
    Board board
    Dice dice
}

Player{
    name
    colour
}

Board{
    [][]Cell board
}




