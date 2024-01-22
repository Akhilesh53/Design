**Requirements**

1) A parking lot is an open space building
2) It might have several floors.
3) each floor has a particular number of spots.
4) each floor has several entry/ exit points.
5) either a spot can be occupied/ free.
6) the spots are divided into 3 types
   - a spot for heavy-weight vehicles
   - a spot for medium-weight vehicles
   - a spot for lightweight vehicles

7) The parking lot will charge some amount for the duration the vehicle is parked.
8) The base charge will be different for each type of vehicle. there might be different strategies to  calculate the charge.
   - heavyweight := 100 base charge up to 1 hr, then 100 extra for every hr
   - medium weight:= 50 base charge up to 1 hr, then 80 extra for every hr
   - lightweight := 30 base charge up to 1 hr, then 50 extra for every hr

9) A ticket/ receipt will be generated for each vehicle parked
10) payment should be done at the exit of the vehicle
11) we should know how many slots are available  on each floor.
12) Only that many vehicles can be parked until we have free slots.
13) Customer  can pay via both cards and cards
14) A ticket will be generated at the entry point and the bill will be generated at the exit point.
15) There can be different strategies to get free/ available slots

-------------------------------
**Additional use cases**

- there can be vehicle types like diesel/petrol, electric, bikes
- spots might have chargers for electric vehicles
- an additional charge will be there charging
- charging facility might be provided by an operator
- operator details need to be stored

----------------------------------
**Functionalities Implemented**

- Generate Ticket On Entry
- Generate Bill on Exit
