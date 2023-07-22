package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

type UserCSV struct {
	ID   string
	Name string
	Age  string
	City string
}

func main() {

	users, err := ReadJSONUsers("resources/user.json")
	userCsvs, err := ReadCSVUsers("resources/user.csv")
	if err != nil {
		return
	}
	PrintJSON(*users)

	PrintCSV(userCsvs)
}

func ReadCSVUsers(filename string) (users []UserCSV, err error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		userCsv := UserCSV{
			ID:   record[0],
			Name: record[1],
			Age:  record[2],
			City: record[3],
		}
		users = append(users, userCsv)
	}
	return users, nil
}

func ReadJSONUsers(filename string) (users *[]User, err error) {
	jsonFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = json.Unmarshal(jsonFile, &users)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return users, nil
}

func PrintJSON(users []User) {
	for i := 0; i < len(users); i++ {
		fmt.Println("ID : ", users[i].ID)
		fmt.Println("Name : ", users[i].Name)
		fmt.Println("Age : ", users[i].Age)
		fmt.Println("City : ", users[i].City)
		fmt.Println("===============================")
	}
}

func PrintCSV(userCsvs []UserCSV) {
	for i := 0; i < len(userCsvs); i++ {
		fmt.Println("ID : ", userCsvs[i].ID)
		fmt.Println("Name : ", userCsvs[i].Name)
		fmt.Println("Age : ", userCsvs[i].Age)
		fmt.Println("City : ", userCsvs[i].City)
		fmt.Println("===============================")
	}
}
