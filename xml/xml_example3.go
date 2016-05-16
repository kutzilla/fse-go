package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Customers struct {
	XMLName   xml.Name   `xml:"customers"`
	Version   string     `xml:"version,attr"`
	Customers []Customer `xml:"customer"`
}

type Customer struct {
	Name         Name         `xml:"name"`
	Address      Address      `xml:"address"`
	Gender       string       `xml:"gender"`
	Birthday     string       `xml:"birthday"`
	Subscription Subscription `xml:"subscription"`
}

type Name struct {
	FirstName string `xml:"first,attr"`
	Middle    string `xml:"middle,attr"`
	Family    string `xml:"family,attr"`
}

type Address struct {
	Street  string `xml:"street,attr"`
	HouseNr string `xml:"houseNr,attr"`
	City    string `xml:"city,attr"`
	ZipCode string `xml:"zipCode,attr"`
	Region  string `xml:"region,attr"`
}

type Subscription struct {
	Imsi   uint   `xml:"imsi,attr"`
	Msisdn uint   `xml:"msisdn,attr"`
	Tariff string `xml:"tariff,attr"`
}

var allCustomers []Customer

func main() {

	file, err := os.Open("customers.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close() // Datei schließen nach der main function
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Customers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	allCustomers = v.Customers

	// Hole Liste mit allen Customers aus Nordrhein-Westfalen
	westfalen := getCustomersFromRegion("Nordrhein-Westfalen")
	printCustomers(westfalen)

	// Hole Liste mit allen Jaqulines aus Sachsen
	allJaquelinesFromSachsen := getCustomersWithNamesFromRegion("Sachsen", "Jaqueline")
	printCustomers(allJaquelinesFromSachsen)
}

func getCustomersFromRegion(region string) []Customer {
	var result []Customer
	for _, customer := range allCustomers {
		if customer.Address.Region == region {
			result = append(result, customer)
		}
	}
	return result
}

func getCustomersWithNamesFromRegion(region string, name string) []Customer {
	regionResults := getCustomersFromRegion(region)

	var result []Customer
	for _, customer := range regionResults {
		if customer.Name.FirstName == name {
			result = append(result, customer)
		}
	}
	return result
}

// Ausgabe der Customers
func printCustomers(customers []Customer) {
	for _, customer := range customers {
		fmt.Println(customer)
	}
}

func (c Customer) String() string {
	return fmt.Sprintf("%s %s aus %s in %s", c.Name.FirstName, c.Name.Family, c.Address.City, c.Address.Region)
}
