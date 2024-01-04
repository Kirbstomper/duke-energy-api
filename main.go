package main

import (
	"encoding/xml"
	"log"
)

func main() {

	//Lets read in our XML

	blob := `
	<espi:IntervalBlock>

	<espi:interval>
	<espi:servicePointId>6002421765</espi:servicePointId>
	<espi:serviceType>ELECTRIC</espi:serviceType>
	<espi:unitOfMeasure>kWH</espi:unitOfMeasure>
	<espi:secondsPerInterval>1800</espi:secondsPerInterval>
	<espi:duration>65840400</espi:duration>
	<espi:start>1659243600</espi:start>
	</espi:interval>
	<espi:IntervalReading>
	<espi:timePeriod>
		<espi:start>1659243600</espi:start>
	</espi:timePeriod>
	<espi:readingQuality>ACTUAL</espi:readingQuality>
	<espi:value>0.12</espi:value>
</espi:IntervalReading>
<espi:IntervalReading>
	<espi:timePeriod>
		<espi:start>1659245400</espi:start>
	</espi:timePeriod>
	<espi:readingQuality>ACTUAL</espi:readingQuality>
	<espi:value>0.13</espi:value>
</espi:IntervalReading>
<espi:IntervalReading>
	<espi:timePeriod>
		<espi:start>1659247200</espi:start>
	</espi:timePeriod>
	<espi:readingQuality>ACTUAL</espi:readingQuality>
	<espi:value>0.1</espi:value>
</espi:IntervalReading>
<espi:IntervalReading>
	<espi:timePeriod>
		<espi:start>1659249000</espi:start>
	</espi:timePeriod>
	<espi:readingQuality>ACTUAL</espi:readingQuality>
	<espi:value>0.28</espi:value>
</espi:IntervalReading>
<espi:IntervalReading>
	<espi:timePeriod>
		<espi:start>1659250800</espi:start>
	</espi:timePeriod>
	<espi:readingQuality>ACTUAL</espi:readingQuality>
	<espi:value>0.39</espi:value>
</espi:IntervalReading>
<espi:IntervalReading>
	<espi:timePeriod>
		<espi:start>1659252600</espi:start>
	</espi:timePeriod>
	<espi:readingQuality>ACTUAL</espi:readingQuality>
	<espi:value>0.1</espi:value>
</espi:IntervalReading>
</espi:IntervalBlock>

`

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
		Interval interval          `xml:"interval"`
		Readings []IntervalReading `xml:"IntervalReading"`
	}
	if err := xml.Unmarshal([]byte(blob), &intervalBlock); err != nil {
		log.Fatal(err)
	}

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
