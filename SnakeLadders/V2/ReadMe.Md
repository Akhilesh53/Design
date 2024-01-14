UseCase

1) In a Game, At max 4 players can play differentiated by color.
2) Once a color is assigned, it cannot be assigned to anyother player. Assignment of color will be random.
3) Need to maintain the winner for all played games.
4) Game has a board of 10x10 (100 cells) dimensions.
5) A block or cell consists of a number.
6) Consider snake and ladder as different objects in game.
7) Player rolls the dice to know how many steps he/she has to move.
8) If dice value comes out to be 6, the dice will be roled again. Total steps movew will be the sum of all values we get from rolling the dice.
9) If we get 6, three time continuously, curr player turn gets cancelled.
10) Game starts from block 1 and The player who reached 100th  block first, wins.
11) Dice value will be from 1-6.
12) A board has number in zig zag fashion like first 1-10 (from left to right) and then 11-20 (right to left) and so on.



User can give 2 inputs:
Play Game    // Launch game
Check Winner for gameId
GetTotalReward   // Get the total reward by entering user id

Note: There is some pointer problem, not able to figure out till now. Will update it once it get resolved