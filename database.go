package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Entry struct {
	Id       uint
	Long_Url string
	Suffix   string
}

var Entries []Entry
var count_of_entries uint = 0
var file *os.File

func load_database(file_name string) { // csv file for now
	// called from external program
	var err error
	file, err = os.OpenFile(file_name, os.O_CREATE|os.O_RDWR, 0644) //opened for r/w, should close the file at the end
	defer file.Close()
	if err != nil {
		fmt.Println("Err opening database:", err)
		return
	}
	csv_file := csv.NewReader(file) // open as csv
	content, err := csv_file.ReadAll()
	if err != nil {
		fmt.Println("Err opening csv file:", err)
	}
	Entries = arr2ent(&content) // load all data into a []struct
	count_of_entries = uint(len(Entries))

}
func add_entry(input []string) error { // called from external program input : {$Long_url, $Suffix}
	var entry = Entry{
		Id:       count_of_entries,
		Long_Url: input[0],
		Suffix:   input[1],
	}
	Entries = append(Entries, entry)
	err := write2csv()
	if err != nil {
		Entries = Entries[:len(Entries)-1] // entry is added to the array but not the csv file
	}
	return err
}

func write2csv() error {
	if file == nil {
		fmt.Println("File is not loaded yet")
		return errors.New("File is not loaded yet")
	}
	writer := csv.NewWriter(file)
	var err error
	err = writer.WriteAll(ent2arr(&Entries))
	if err != nil {
		fmt.Println("Error writing to database:", err)
	}
	return err
}

func arr2ent(content *[][]string) []Entry {
	var entries []Entry
	var id uint64
	for _, row := range *content {
		id, _ = strconv.ParseUint(row[0], 0, 0)
		entry := Entry{
			Id:       uint(id),
			Long_Url: row[1],
			Suffix:   row[2],
		}
		entries = append(entries, entry)
	}
	return entries
}
func ent2arr(entries *[]Entry) [][]string {
	var datas [][]string
	var data []string
	for _, row := range *entries {
		data = append(data, strconv.FormatUint(uint64(row.Id), 10))
		data = append(data, row.Long_Url)
		data = append(data, row.Suffix)
		datas = append(datas, data)
	}
	return datas
}
func search_in_database() {}
