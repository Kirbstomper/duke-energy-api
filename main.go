package main

import (
	"encoding/xml"
	"log"
	"os"
)

func main() {

	//Lets read in our XML

	reportFile, err := os.ReadFile("report.xml")

	if err != nil {
		log.Fatal(err)
	}

	type interval struct {
		ServicePointID     string `xml:"servicePointId"`
		ServiceType        string `xml:"serviceType"`
		UnitOfMeasure      string `xml:"unitOfMeasure"`
		SecondsPerInterval string `xml:"secondsPerInterval"`
		Duration           string `xml:"duration"`
		Start              string `xml:"start"`
	}

	type IntervalReading struct {
		ReadingQuality string `xml:"readingQuality"`
		Value          string `xml:"value"`
		Start          string `xml:"timePeriod>start"`
	}
	var intervalBlock struct {
		Interval interval          `xml:"content>IntervalBlock>interval"`
		Readings []IntervalReading `xml:"content>IntervalBlock>IntervalReading"`
	}
	if err := xml.Unmarshal(reportFile, &intervalBlock); err != nil {
		log.Fatal(err)
	}

	//Print out the results
	println("Service Point ID: " + intervalBlock.Interval.ServicePointID)
	println("Service Type: " + intervalBlock.Interval.ServiceType)
	println("Unit of Measure: " + intervalBlock.Interval.UnitOfMeasure)
	println("Seconds Per Interval: " + intervalBlock.Interval.SecondsPerInterval)
	println("Duration: " + intervalBlock.Interval.Duration)
	println("Start: " + intervalBlock.Interval.Start)

	for i := 0; i < len(intervalBlock.Readings); i++ {
		println("Reading Quality: " + intervalBlock.Readings[i].ReadingQuality)
		println("Value: " + intervalBlock.Readings[i].Value)
		println("Time Period Start: " + intervalBlock.Readings[i].Start)
	}

}
