package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jweir/csv"
)

type MachineInfo struct {
	Name      string `csv:"name"`
	Address   string `csv:"address (S)"`
	Latitude  string `csv:"latitude (N)"`
	Longitude string `csv:"longitude (N)"`
}

// type MachineInfoWithFloat struct {
// 	Latitude  float64
// 	Longitude float64
// }

type FinalMachineInfo struct {
	Name      string
	Address   string
	Latitude  float64
	Longitude float64
}

func UnmarshalCSV() ([]FinalMachineInfo, error) {
	finalMachineInfo := []FinalMachineInfo{}
	machineLocations := []MachineInfo{}

	sample := []byte(
		`name,address (S),latitude (N),longitude (N)
University of Illinois at Chicago - Student Center West,"Student Residence Hall, Chicago, IL 60612",41.871103,-87.6745017
600 W Chicago,"600 W Chicago Ave, Chicago, IL 60654",41.8975186,-87.6450724
Chase Tower,"Chase Tower, 111 E Wisconsin Ave, Milwaukee, WI 53202",43.0379231,-87.9093525
7-Eleven @ Kingsbury and Ontario ,"645 N Kingsbury St, Chicago, IL 60654",41.8933902,-87.6410962
Feinberg Pavilion - Northwestern Medicine,"251 E Huron St, Chicago, IL 60611",41.8946401,-87.6211275
Chicago Midway Airport - Ticketing Employee Lounge,"5700 S Cicero Ave, Chicago, IL",41.7883501,-87.741842
DeVry Chicago Campus (Students/Staff Only),"3300 N Campbell Ave, Chicago, IL 60618",41.942132,-87.691461
525 W Monroe,"525 W Monroe St, Chicago, IL 60661",41.8801934,-87.6401187
200 W Jackson (Tenants Only),"200 W Jackson Blvd, Chicago, IL  60606",41.878511,-87.634277
Allstate HQ (Tenants Only),"3075 Sanders Rd, Northbrook, IL  60062",42.09674,-87.870095
`)
	err := csv.Unmarshal(sample, &machineLocations)

	if err != nil {
		fmt.Println("Error unmarshalling data: ", err.Error())
		return finalMachineInfo, err
	}
	for i, _ := range machineLocations {
		newLat, err := strconv.ParseFloat(machineLocations[i].Latitude, 64)
		newLong, err := strconv.ParseFloat(machineLocations[i].Longitude, 64)
		if err != nil {
			fmt.Println("Error converting string to float for ", machineLocations[i].Name, err.Error())
		}
		machineInfo := FinalMachineInfo{
			Name:      machineLocations[i].Name,
			Address:   machineLocations[i].Address,
			Latitude:  newLat,
			Longitude: newLong,
		}
		finalMachineInfo = append(finalMachineInfo, machineInfo)
	}
	fmt.Printf("Final Output with %+v", finalMachineInfo)
	return finalMachineInfo, nil
}

func Distance(locations []FinalMachineInfo) float64 {
	// var final []FinalMachineInfo
	for _, points := range locations {
		for _, innerPoints := range locations {
			lat1 := points.Latitude
			lat2 := innerPoints.Latitude
			lon1 := points.Longitude
			lon2 := innerPoints.Longitude

			var R = float64(6371)
			var x1 = lat2 - lat1
			var dLat = x1 * math.Pi / 180
			var x2 = lon2 - lon1
			var dLon = x2 * math.Pi / 180
			var a = math.Sin(dLat/2)*math.Sin(dLat/2) +
				math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
					math.Sin(dLon/2)*math.Sin(dLon/2)
			var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
			var distance = R * c
			fmt.Printf("distance from %s to %s: %f", points.Name, innerPoints.Name, distance)
			// final = append(final, distance)
		}
		return 0.0
	}
	return 0.0
}

func main() {
	finalMachineInfo, err := UnmarshalCSV()
	if err != nil {
		fmt.Println("Error retrieving unmarshalled data in main", err.Error())
	}
	Distance(finalMachineInfo)

}
