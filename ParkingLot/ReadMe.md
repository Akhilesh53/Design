Requirements
======================

1) A parking lot is an open space building
2) It might have number of floors.
3) each floor has particular number of spots.
4) each floor has number of entry/ exit points.
5) either a spot can be occupied/ free.
6) the spot are divided into 3 types
   - spot for heavy vehicles
   - spot for medium weight
   - spot for light weight vechicles

7) parking lot will charge some amount for the duration vehicle is parked.
8) Base charge will be different for each type of vehicle. there might be different strategies to  calculate the charge.
   - heavy weight := 100 base charge upto 1 hr, then 100 extra for every hr
   - medium weight := 50 base charge upto 1 hr, then 80 extra for every hr
   - light weight := 30 base charge upto 1 hr, then 50 extra for every hr

9) A ticket/ receipt will be generated for each vechile parked
10) payment should be done at the exit of vehicle
11) we should know how many slots are available  at each floor.
12) Only that much vechicles can be parked, untill we have free slots.
13) Customer  can pay via both cards and cards
14) A ticket will be genreated at the entry point and bill is generated at exit point.
15) There can be different strategies to get free/ available slots

================================================================
Additional use cases
================================================================
- there can be vechile types like diesel/petrol, electric, bikes
- spots might have chargeres for electic vehicles
- additional charge will be there charging
- charging facility might be provided by an operator
- operator details need to be stored


================================================================
Functionalities Implemented
================================================================
- Generate Ticket On Entry
- Generate Bill on Exit