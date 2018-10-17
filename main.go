package main

import (
	"fmt"

	"github.com/jweir/csv"
)

type MachineInfo struct {
	Name      string `csv:"name"`
	Address   string `csv:"address (S)"`
	Latitude  string `csv:"latitude (N)"`
	Longitude string `csv:"longitude (N)"`
}

// func (m *MachineInfo) UnmarshalCSV(val string, row *csv.Row) error {
// 	name, _ := row.Named("Name")
// 	address, _ := row.Named("Address")
// 	latitude, _ := row.Named("Latitude")
// 	longitude, _ := row.Named("Longitude")

// 	m.Name = name
// 	m.Address = address
// 	m.Latitude = latitude
// 	m.Longitude = longitude

// 	return nil
// }

func UnmarshalCSV() error {
	machineLocations := []MachineInfo{}

	sample := []byte(
		`name,address (S),latitude (N),longitude (N)
University of Illinois at Chicago - Student Center West,"Student Residence Hall, Chicago, IL 60612",41.871103,-87.6745017
600 W Chicago,"600 W Chicago Ave, Chicago, IL 60654",41.8975186,-87.6450724
Chase Tower,"Chase Tower, 111 E Wisconsin Ave, Milwaukee, WI 53202",43.0379231,-87.9093525"
`)

	err := csv.Unmarshal(sample, &machineLocations)

	if err != nil {
		fmt.Println("Error unmarshalling data: ", err.Error())
		return err
	}

	fmt.Printf("%+v", machineLocations)
	return nil
}

func main() {
	UnmarshalCSV()
}
