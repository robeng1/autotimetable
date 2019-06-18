package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main()  {
	generateExtreme(1000, "source.csv", "xtreme.csv")
}

func generateExtreme(numberOfRooms int, source, dest string) bool{
	csvFile, _ := os.Open(source)
	in := csv.NewReader(bufio.NewReader(csvFile))
	records, err := in.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var acc int64
	for i, r := range records{
		if i == 0{
			continue
		}
		c, _ := strconv.ParseInt(r[5], 0, 0)
		acc+=c
	}
	total := 50*numberOfRooms
	var tmp [][]string
	if acc != int64(total){
		tmp = append(tmp, records...)
		maxIndex := len(records)
		minIndex := 0
		for i := acc; acc != int64(total); i++{
			randomIndex1 := rand.Intn(maxIndex -minIndex) + minIndex
			randomIndex2 := rand.Intn(maxIndex -minIndex) + minIndex
			newR1 := records[randomIndex1]
			newR2 := records[randomIndex2]
			c, _ := strconv.ParseInt(newR1[6], 0, 0)
			c = int64(rand.Intn((5)) +1)
			yr := rand.Intn(3)+1
			newRecord := []string{
				newR1[0],newR2[1],newR2[2],newR1[3],newR2[4], strconv.Itoa(int(c)), newR2[6], newR1[7],"", "", "", strconv.Itoa(yr), newR2[12], newR1[13],newR2[14],
			}
			tmp = append(tmp, newRecord)
			c, _ = strconv.ParseInt(newRecord[5], 0, 0)
			acc+=c
		}
	}
	if tmp != nil{
		writeAllCSV(dest, tmp)
		fmt.Println(acc)
		return true
	}
	return false
}

func writeAllCSV(fileName string, values [][]string){
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}
	
	csvWriter := csv.NewWriter(file)
	_ = csvWriter.WriteAll(values)
	csvWriter.Flush()
}
