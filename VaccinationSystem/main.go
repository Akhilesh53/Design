package main

import (
	"fmt"
	"pattern/VaccinationSystem/entities"
	"pattern/VaccinationSystem/services"

	"google.golang.org/genproto/googleapis/type/date"
)

func main() {

	userService := services.NewUserService()
	vcService := services.NewVCService()
	bookingService := services.NewBookingService()

	entryDate := &date.Date{
		Year:  2023,
		Month: 11,
		Day:   1,
	}

	// Add VC
	vc1 := entities.NewVaccinationCentre(101, "Kolkata", "Kolkata district VC address")
	vc2 := entities.NewVaccinationCentre(102, "Noida", "Noida district VC address")

	vc1.SetCapacityForADay(entryDate.String(), 100)
	vc2.SetCapacityForADay(entryDate.String(), 100)

	vcService.AddVC(vc2)
	vcService.AddVC(vc1)

	const printStr = `
	   1. Register User
	   2. See VC 
	   3. Book Slot
	   4. Exit
	`

	for {

		fmt.Println(printStr)

		var inputVal int
		fmt.Scanln(&inputVal)

		switch inputVal {
		case 1:

			fmt.Println("Register User")

			var userId int
			var name string
			var year, month, day int32
			var isVaccinated bool

			fmt.Println("Enter User Id")
			fmt.Scanln(&userId)

			fmt.Println("Enter Name")
			fmt.Scanln(&name)

			fmt.Println("Enter DOB Year")
			fmt.Scanln(&year)

			fmt.Println("Enter DOB Month")
			fmt.Scanln(&month)

			fmt.Println("Enter DOB Day")
			fmt.Scanln(&day)

			fmt.Println("Enter isVaccinated")
			fmt.Scanln(&isVaccinated)

			dob := &date.Date{
				Year:  year,
				Month: month,
				Day:   day,
			}
			userService.AddUser(userId, name, dob, isVaccinated)

			fmt.Println("User Added with id", userId, "name", name, "dob", dob.String(), "isVaccinated", isVaccinated)

		case 2:
			fmt.Println("See VC")

			// see VC based on district entered and date entered. show only VC which has available slots
			var district string
			var year, month, day int32

			fmt.Println("Enter District")
			fmt.Scanln(&district)

			fmt.Println("Enter BookingDate Year")
			fmt.Scanln(&year)

			fmt.Println("Enter BookingDate  Month")
			fmt.Scanln(&month)

			fmt.Println("Enter BookingDate Day")
			fmt.Scanln(&day)

			bookingDate := &date.Date{
				Year:  year,
				Month: month,
				Day:   day,
			}

			allVC := vcService.GetVC(district)

			fmt.Println(len(allVC))
			for _, vc := range allVC {
				if vc.GetCapacityForADay(bookingDate.String()) > 0 {
					fmt.Println("Vaccination Centre Available ::", vc.GetId(), vc.GetDistrict()," having slots ", vc.GetCapacityForADay(bookingDate.String()))
				}
			}

		case 3:
			fmt.Println("Book Slot")
			// book a slot based on input date and district and vc id
			var district string
			var year, month, day int32
			var vcId int
			var userId int

			fmt.Println("Enter User Id")
			fmt.Scanln(&userId)

			fmt.Println("Enter Vaccination Centre District")
			fmt.Scanln(&district)

			fmt.Println("Enter BookingDate Year")
			fmt.Scanln(&year)

			fmt.Println("Enter BookingDate Month")
			fmt.Scanln(&month)

			fmt.Println("Enter BookingDate Day")
			fmt.Scanln(&day)

			fmt.Println("Enter Vaccination Centre Id")
			fmt.Scanln(&vcId)

			bookingDate := &date.Date{
				Year:  year,
				Month: month,
				Day:   day,
			}

			// if user is less than 18 or user has some booking alread or user is already vaccinatedy, dont book
			user := userService.GetUser(userId)

			if user.GetIsVaccinated() {
				fmt.Println("User already vaccinated")
				continue
			}

			if user.GetDob().GetYear() > year-18 {
				fmt.Println("User is less than 18")
				continue
			}

			if len(user.GetBookings()) > 0 {
				fmt.Println("User already has booking")
				continue
			}

			// if vc has slot availle or not
			allVC := vcService.GetVC(district)

			for _, vc := range allVC {
				if vc.GetId() == vcId {
					if vc.GetCapacityForADay(bookingDate.String()) > 0 {
						booking := entities.NewBooking(1, user, vc, bookingDate)
						bookingService.AddBooking(booking.GetId(), booking.GetDate(), booking.GetIUser(), booking.GetIVC())
						user.AddBoookingForADay(bookingDate.String(), booking)
						vc.AddBookingForADay(bookingDate.String(), user)
						fmt.Println("Slot booked for user: ", user.GetName(), " for date: ", bookingDate.String())
					} else {
						fmt.Println("No slot available")
						continue
					}
				}
			}
		case 4:
			fmt.Println("Exit")
			break
		}

	}
}
