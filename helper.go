package main

import (
	"context"
	"fmt"
	cowin "github.com/sedflix/cowin-availability-checker/swagger"
	"strings"
)

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
