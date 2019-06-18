package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)


type room struct {
	Name string
	Building string
	Capacity int
	Allowance int
}

func (r *room)String() string{
	return fmt.Sprintf("%s, %s, %d, %d", r.Name, r.Building, r.Capacity, r.Allowance)
}


var maxCapacity = 500
var minCapacity = 100
var maxAllowance = 50
var minAllowance = 5

//capacity will be randomly generated
type BlockGenerator struct {
	NumberingRange []int
	AllowanceRange []int
	BuildingName   string
	Prefix string
	//this specifies whether it should be alphabet or numeric
	RangeType int
}

func (b *BlockGenerator) generate() []room{
	var rooms []room
		for i:=0; i <= b.NumberingRange[1]; i ++{
			rooms = append(rooms,
				room{fmt.Sprintf("%s%d", b.Prefix, i), b.BuildingName,rand.Intn(maxCapacity - minCapacity) + minCapacity,rand.Intn(maxAllowance - minAllowance) + minAllowance},
			)
		}
	return rooms
}

func (b *BlockGenerator)writeAllCSV(fileName string){
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}

	rooms := b.generate()
	values := [][]string{
		{"Name", "Building", "Capacity", "Allowance"},
	}
	for _, i := range rooms{
		values= append(values, []string{i.Name, i.Building, fmt.Sprintf("%d", i.Capacity), fmt.Sprintf("%d", i.Allowance)})
	}

	csvWriter := csv.NewWriter(file)
	_ = csvWriter.WriteAll(values)
	csvWriter.Flush()
}

func main()  {

	LTgen := BlockGenerator{NumberingRange: []int{1, 1000}, BuildingName:"Petroleum Building", Prefix:"PB"}
	LTgen.writeAllCSV("rooms.csv")

}
