package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/jweir/csv"
)

type MachineInfo struct {
	Name      string `csv:"name"`
	Address   string `csv:"address (S)"`
	Latitude  string `csv:"latitude (N)"`
	Longitude string `csv:"longitude (N)"`
}

type FinalMachineInfo struct {
	Name      string
	Address   string
	Latitude  float64
	Longitude float64
}

type Distances struct {
	From     string
	To       string
	Distance float64
}

type Node struct {
	Name    string
	Visited bool
}

func UnmarshalCSV() ([]FinalMachineInfo, error) {
	finalMachineInfo := []FinalMachineInfo{}
	machineLocations := []MachineInfo{}

	sample := []byte(
		`name,address (S),latitude (N),longitude (N)
Kitchen,"Lake and Racine Ave",41.8851024,-87.6618988
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
CVS @ 137 State,"137 S State St, Chicago, IL 60603",41.8796828,-87.6272829
West Suburban Medical Center,"1 Erie St, Oak Park, IL 60302",41.8913716,-87.7761143
Merchandise Mart,"222 W Merchandise Mart Pl, Chicago, IL 60654",41.888221,-87.635419
American Airlines Lounge (Employees Only),"10000 Ohare Ave, Chicago, IL 60666",41.978568,-87.9082054
Medical College of Wisconsin,"8701 W Watertown Plank Rd, Milwaukee, WI 53226",41.8969173,-87.6435474
100 E Wisconsin,"100 East Wisconsin, 100 E Wisconsin Ave, Milwaukee, WI 53202",43.0389389,-87.9095663
O'Hare 3H Left,"10000 Ohare Way, Chicago, IL 60666",41.9744812,-87.9094719
CME Center,"30 S Wacker Dr, Chicago, IL 60606",41.8813392,-87.6375575
Illinois Center,"111 E Upper Wacker Dr, Chicago, IL 60601",41.8876736,-87.6236084
Schlitz Park Rivercenter,"1555 N Rivercenter Dr, Milwaukee, WI 53212",43.050891,-87.910165
US Foods (Employees Only),"9377 W Higgins Rd, Rosemont, IL 60018",41.9886856,-87.8578175
North Park University,"3225 W Foster Ave, Chicago, IL 60625",41.8969173,-87.6435474
Epic Burger West Loop,"550 W Adams St, Chicago, IL 60661",41.8795147,-87.6417299
University of Wisconsin-Milwaukee EMS Building,"3200 N Cramer St, Milwaukee, WI 53211",43.0758306,-87.8857607
311 S Wacker (Tenants Only),"311 S Wacker Dr, Chicago, IL 60606",41.877458,-87.635963
MacNeal Hospital,"3249 S Oak Park Ave, Berwyn, IL 60402",41.84327,-87.793133
General Growth Properties HQ (Employees Only),"110 N Upper Wacker Dr, Chicago, IL 60606",41.883801,-87.6374336
University of Illinois at Chicago - College of Medicine ,"835 S Wolcott Ave, Chicago, IL 60612",41.8706072,-87.6734143
Allstate HQ (Tenants Only),"3075 Sanders Rd, Northbrook, IL  60062",42.09674,-87.870095
Prentice Womens Hospital - Northwestern Medicine,"250 E Superior St, Chicago, IL 60611",41.896332,-87.62095
DePaul University - Schmitt Academic Center,"2320 N Kenmore Ave, Chicago, IL 60614",41.924,-87.655134
University of Illinois at Chicago - Behavioral Sciences,"1007 W Harrison Street, Chicago, IL 60607",41.8735982,-87.6529995
CNA Center (Employees Only),"333 S Wabash Ave, Chicago, IL 60604",41.8773837,-87.6255692
University of Illinois at Chicago - Richard J Daley Library (Students/Staff Only),"801 S Morgan Avenue, Chicago, IL 60607",41.871871,-87.6504857
Chicago Midway Airport - Southwest Employee Lounge,"Midway International Airport, 5700 S Cicero Ave, Chicago, IL, USA",41.7883501,-87.741842
Schaumburg Towers,"1400 American Ln, Schaumburg, IL 60196",42.0453601,-88.0436918
Loyola Medical Center,"2160 S 1st Ave, Maywood, IL 60153",41.8605065,-87.8351127
7-Eleven @ Jackson and Desplaines,"627 W Jackson Boulevard, Chicago, IL 60661",41.8778687,-87.6436654
Walgreens Corporate Office (Employees Only),"1435 Lake Cook Rd, Deerfield, IL 60015",42.1513812,-87.8644641
Skokie Hospital / Northshore ,"9600 Gross Point Road, Skokie IL 60076",42.0563891,-87.7429858
DeVry Addison (Students/Staff Only)," 1221 N Swift Rd, Addison, IL 60101",41.94976,-88.038531
CVS @ 344 Hubbard,"344 W Hubbard St, Chicago, IL 60654",41.8901922,-87.6372236
100 E Wisconsin,"100 East Wisconsin, 100 E Wisconsin Ave, Milwaukee, WI 53202",43.0389389,-87.9095663
REI Building ,"1460 N Halsted Street, Chicago, IL 60642",41.907879,-87.648598
Peggy Notebaert Nature Museum,"2430 N Cannon Dr, Chicago, IL 60614",41.9267887,-87.6373734
MillerCoors HQ,"250 S Wacker Dr, Chicago, IL 60606",41.8783905,-87.6373008
Good Samaritan Hospital,"3815 Highland Ave, Downers Grove, IL 60515",41.8969173,-87.6435474
Moraine Valley Community College: Police Academy- Building B,"9000 College Pkwy, Palos Hills, IL 60465, USA",41.8969173,-87.6435474
O'Hare Terminal 2 - Gate F6,"10000 W Ohare Ave, Chicago, IL 60666",41.9742091,-87.906022
Good Shepherd Hospital,"450 IL-22, Barrington, IL 60010",42.196039,-88.1733328
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
	return finalMachineInfo, nil
}

func DistanceBetweenTwoPoints(locations []FinalMachineInfo) ([]Distances, map[float64]string) {
	var distances []Distances
	route := make(map[float64]string)
	for _, points := range locations {
		for j, _ := range locations {
			lat1 := points.Latitude
			lat2 := locations[j].Latitude
			lon1 := points.Longitude
			lon2 := locations[j].Longitude

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
			// fmt.Printf("\ndistance from %s to %s: %f \n", points.Name, locations[j].Name, distance)
			routeName := fmt.Sprintf("From %s to %s", points.Name, locations[j].Name)
			route[distance] = routeName
			newDistance := Distances{
				From:     points.Name,
				To:       locations[j].Name,
				Distance: distance,
			}
			distances = append(distances, newDistance)
			// routeList = append(routeList, route)
		}

	}
	return distances, route
}

func MapDistancesToKey(distancesArray []float64, routes map[float64]string) (map[float64]string, map[float64]string) {
	deDupedRoutesMap := make(map[float64]string)
	var x []float64
	startingPoint := make(map[float64]string)
	for i, _ := range distancesArray {
		deDupedRoutesMap[distancesArray[i]] = routes[distancesArray[i]]
		if strings.Contains(routes[distancesArray[i]], "Kitchen") {
			x = append(x, distancesArray[i])
		}
	}
	sort.Float64s(x)
	startingPoint[x[0]] = "From Kitchen to Behavioral Sciences"
	fmt.Println(startingPoint)
	return deDupedRoutesMap, startingPoint
}

func DeDupe(distancesArray []float64) []float64 {
	dupedDistances := map[float64]bool{}
	deDupedDistances := []float64{}

	for i := range distancesArray {
		if distancesArray[i] != 0 {
			if dupedDistances[distancesArray[i]] != true {
				dupedDistances[distancesArray[i]] = true
				deDupedDistances = append(deDupedDistances, distancesArray[i])
			}
		}
	}
	return deDupedDistances
}

func GrabAllDistances(nodeDistances []Distances) []float64 {
	var distances []float64
	for i, _ := range nodeDistances {
		distances = append(distances, nodeDistances[i].Distance)
	}
	deDuped := DeDupe(distances)
	return deDuped
}

// func FindShortestPath(nodes map[float64]string, startingPoint map[float64]string) [] {

// }

func main() {
	finalMachineInfo, err := UnmarshalCSV()
	if err != nil {
		fmt.Println("Error retrieving unmarshalled data in main", err.Error())
	}
	distances, routes := DistanceBetweenTwoPoints(finalMachineInfo)
	dedupedistance := GrabAllDistances(distances)
	graphOfNodes, _ := MapDistancesToKey(dedupedistance, routes)
	fmt.Println(graphOfNodes)
}
