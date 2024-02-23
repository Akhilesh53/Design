# Backend for tracking Cricket matches and Sore card Application

Come-up with design for cricket website  
1. It can track tournaments
2. Tournament can contain multiple fixtures
3. Fixtures are to be played in multiple locations 
4. Matches will be played between teams (2)
5. Matches are played in stadiums
6. There are multiple stadiums but one stadium can host one match at the time of the day
7. Multiple teams can participate in a tournament
8. Each tournament has a start and end date
9. Teams may represent a country or a club
10. Tournaments can be played by clubs or countries
11. Teams are chosen for a tournament
12. A team contain up to 14 players
13. Team may contain more than one coach
13. Team may contain a doctor, physio and a manager
14. Team has a captain and vice captain
15. Team is composed of players and players are of the following types  
    A. BATSMAN  
    B. BOWLER  
    C. WICKETKEEPER
16. A match contian minimum of two innings
17. Match starts with a toss. Coin is tossed by one captain and asked by another
18. Toss has an outcome on which a team may choose to bowl / bat 
19. Match will involve the following people  
    A. Teams  
    B. Umpires  
    C. Referee  
    D. Scorers  
    E. Commentators    
20. Bowler bowls overs and Batsman scores runs
21. Bowled balls and runs should be tracked on scorecard


----------------------------------------------------------------
Classes
----------------------------------------------------------------
Tournament{
    List<Fixture> fixtures
    List<Team> particpatingTeams
    StartDate
    EndDate
}

Fixture (Matches){
    Team1
    Team2
    Stadium (Location)
    List<Umpire>
    Referee
    Commentators
}

Team{
    TeamName
    Country
    List<Player> players // max 14
    List<Coach> coaches, physio, doctor, manager,
}

Stadiums (Venue){
    OccupancyDetails map[datetimevalue]MatchId 
}

Coach, Doctor,Physio, manager, 
Captain, Vice Captain

PlayerType

Innings{
    playedbywhichteam
    totalRunsScored
    totalWicketsDown
    totalOversPlayed
    scorers // playeres who scored
    bowlers
}

Over{
    overnumber
    List<Ball>
}

ScoreCard{
    Innings1
    Innings2
}