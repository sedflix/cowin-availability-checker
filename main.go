package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gen2brain/beeep"
	cowin "github.com/sedflix/cowin-availability-checker/swagger"
	"time"
)

var client *cowin.APIClient

func Notify(message string) {
	err := beeep.Alert("Vaccination Appointment Available", message, "./assets/icon.jpg")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func IsSessionAvailable(session cowin.SessionCalendarEntrySchemaSessions) bool {
	if session.MinAgeLimit < 45 {
		if session.AvailableCapacity > 0 {
			return true
		}
	}
	return false
}
func IsCentreAvailable(centre cowin.SessionCalendarEntrySchema) bool {
	for _, session := range centre.Sessions {
		if IsSessionAvailable(session) {
			Notify(fmt.Sprintf("Availble at %s on %s. Capacity: %d. Minimum Age: %d.", centre.Name, session.Date, int(session.AvailableCapacity), int(session.MinAgeLimit)))
			return true
		}
	}
	return false
}

func ListAvailableCentres(districtId string) error {
	centres, _, err := client.AppointmentAvailabilityAPIsApi.CalendarByDistrict(
		context.Background(),
		districtId,
		time.Now().Local().Format("02-01-2006"),
		&cowin.AppointmentAvailabilityAPIsApiCalendarByDistrictOpts{},
	)
	if err != nil {
		return err
	}
	for _, centre := range *centres.Centers {
		if IsCentreAvailable(centre) {
			fmt.Printf("%s available \n", centre.Name)
		}
	}
	return nil
}

func main() {
	var state = flag.String("s", "Uttar Pradesh", "Name of the Sate")
	var district = flag.String("d", "Lucknow", "Name of the District")
	flag.Parse()

	client = cowin.NewAPIClient(cowin.NewConfiguration())

	districtId, err := GetDistrictID(*state, *district)
	if err != nil {
		fmt.Printf("%s", err)
		panic(1)
	}

	fmt.Printf("Listing availble centres in %s, %s\n", *district, *state)
	err = ListAvailableCentres(districtId)
	if err != nil {
		fmt.Printf("%s", err)
		panic(1)
	}
	fmt.Printf("End\n")

}
