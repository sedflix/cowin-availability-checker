package main

import (
	"context"
	"flag"
	"fmt"
	cowin "github.com/sedflix/cowin-availability-checker/swagger"
	"strings"
)

var client *cowin.APIClient

func GetStateID(stateName string) (string, error) {
	states, _, err := client.MetadataAPIsApi.States(context.Background(), &cowin.MetadataAPIsApiStatesOpts{})
	if err != nil {
		return "", err
	}
	for _, state := range states.States {
		if strings.Contains(strings.ToLower(state.StateName), strings.ToLower(stateName)) {
			return fmt.Sprintf("%f", state.StateId), nil
		}
	}
	return "", fmt.Errorf("unbale to find state with name - %s", stateName)
}

func GetDistrictIDUsingStateID(stateID, districtName string) (string, error) {

	districts, _, err := client.MetadataAPIsApi.Districts(
		context.Background(),
		stateID,
		&cowin.MetadataAPIsApiDistrictsOpts{},
	)
	if err != nil {
		return "", err
	}

	for _, district := range districts.Districts {
		if strings.Contains(strings.ToLower(district.DistrictName), strings.ToLower(districtName)) {
			return fmt.Sprintf("%f", district.DistrictId), nil
		}
	}
	return "", fmt.Errorf("unbale to find district with name - %s", districtName)

}

func GetDistrictID(stateName, districtName string) (string, error) {
	stateID, err := GetStateID(stateName)
	if err != nil {
		return "", err
	}
	districtID, err := GetDistrictIDUsingStateID(stateID, districtName)
	if err != nil {
		return "", err
	}
	return districtID, nil
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
			return true
		}
	}
	return false
}
func ListAvailableCentres(districtId string) error {
	centres, _, err := client.AppointmentAvailabilityAPIsApi.CalendarByDistrict(
		context.Background(),
		districtId,
		"03-05-2021",
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

	err = ListAvailableCentres(districtId)
	if err != nil {
		fmt.Printf("%s", err)
		panic(1)
	}

}
