//Ninja Level 4
//Hands-on Exercise 6
//Create a slice to store the names of all of the province in the Indonesia.
// What is the length of your slice?
//What is Capacity?
//Print out all of the values,
// along with their index position in the slice,
// without using range clause. Here is a list of these province
//Aceh
//North Sumatra
//West Sumatra
//Riau
//Riau Islands
//Jambi
//Bengkulu
//South Sumatra
//Bangka Belitung
//Lampung
//Banten
//Jakarta SCR
//West Java
//Central Java
//Yogyakarta SR
//East Java
//Bali
//West Nusa Tenggara
//East Nusa Tenggara
//West Borneo
//Central Borneo
//South Borneo
//East Borneo
//North Borneo
//North Sulawesi
//Gorontalo
//Central Sulawesi
//West Sulawesi
//South Sulawesi
//Southeast Sulawesi
//North Maluku
//Maluku
//West Papua
//Papua
package main

import (
	"fmt"
)

func main() {
	x := make([]string,50,50)
	x=[]string{"Aceh",
		"North Sumatra",
		"West Sumatra",
		"Riau",
		"Riau Islands",
		"Jambi",
		"Bengkulu",
		"South Sumatra",
		"Bangka Belitung",
		"Lampung",
		"Banten",
		"Jakarta SCR",
		"West Java",
		"Central Java",
		"Yogyakarta SR",
		"East Java",
		"Bali",
		"West Nusa Tenggara",
		"East Nusa Tenggara",
		"West Borneo",
		"Central Borneo",
		"South Borneo",
		"East Borneo",
		"North Borneo",
		"North Sulawesi",
		"Gorontalo",
		"Central Sulawesi",
		"West Sulawesi",
		"South Sulawesi",
		"Southeast Sulawesi",
		"North Maluku",
		"Maluku",
		"West Papua",
		"Papua",
	}
	fmt.Println(x)
	fmt.Println(len(x)) //lebar yang dibuat
	fmt.Println(cap(x)) //kapasitas yang disediakan
	for i:= 0; i<len(x); i++{
		fmt.Println(x[i])
	}
}

