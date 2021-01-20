package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
)

func main() {

	//create a file
	fileName := "data.csv"
	file, err := os.Create(fileName)

	if err != nil {
		//displays message & exists program
		log.Fatalf("Couldn't create the file. Error: %q", err)
		return
	}
	//defer closes the file after we are done with it
	defer file.Close()

	//we need a writer to write to the csv file
	writer := csv.NewWriter(file)

	defer writer.Flush()

	log.Println("Scraping...")

	collector := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)

	collector.OnHTML(".internship_meta", func(element *colly.HTMLElement) {
		_ = writer.Write([]string{
			element.ChildText("a"),
			element.ChildText("span"),
		})
	})

	for i := 0; i < 283; i++ {
		fmt.Printf("We're visiting page: %d\n", i)
		_ = collector.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}

	log.Println("Scraping Complete")
}
