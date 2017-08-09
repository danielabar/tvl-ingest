package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// TtcURL is the endpoint for vehicleLocations command of TTC real-time API
const TtcURL = "http://webservices.nextbus.com/service/publicXMLFeed?command=vehicleLocations&a=ttc"

// TODO https://github.com/kellydunn/golang-geo for calc distance between points
func main() {
	route := "504"
	t := strconv.Itoa(0)
	url := TtcURL + "&r=" + route + "&t=" + t
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Print("Unable to get vehicleLocations")
		return
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Unable to reqad vehicleLocations response")
		return
	}

	body := new(Body)
	// https://golang.org/pkg/encoding/xml/#example_Unmarshal
	xml.Unmarshal(responseBody, &body)
	fmt.Printf("MyTime = %v, First Vehicle = %s, location = %f,%f\n", body.LastReported[0].MyEpoch, body.Vehicles[0].VehicleID, body.Vehicles[0].Lat, body.Vehicles[0].Lon)
	log.Printf("Whole thing = %v", body)
}

/*
Sample response from vehicleLocations command API
<?xml version="1.0" encoding="utf-8" ?>
<body copyright="All data copyright Toronto Transit Commission 2017.">
    <vehicle id="4229" routeTag="504" dirTag="504_1_504" lat="43.667183" lon="-79.353386" secsSinceReport="4" predictable="true" heading="168"/>
    <vehicle id="4210" routeTag="504" dirTag="504_0_504" lat="43.646549" lon="-79.390282" secsSinceReport="5" predictable="true" heading="74"/>
    <vehicle id="4222" routeTag="504" dirTag="504_1_504" lat="43.64135" lon="-79.415398" secsSinceReport="15" predictable="true" heading="255"/>
    <vehicle id="4200" routeTag="504" dirTag="504_0_504" lat="43.656849" lon="-79.453117" secsSinceReport="10" predictable="true" heading="257"/>
    <vehicle id="4242" routeTag="504" dirTag="504_0_504" lat="43.656216" lon="-79.357269" secsSinceReport="4" predictable="true" heading="51"/>
    <vehicle id="4208" routeTag="504" dirTag="504_1_504" lat="43.653049" lon="-79.362717" secsSinceReport="5" predictable="true" heading="229"/>
    <vehicle id="4226" routeTag="504" dirTag="504_1_504" lat="43.677067" lon="-79.358269" secsSinceReport="13" predictable="true" heading="67"/>
    <vehicle id="4202" routeTag="504" dirTag="504_0_504" lat="43.656765" lon="-79.453796" secsSinceReport="10" predictable="true" heading="263"/>
    <vehicle id="4223" routeTag="504" dirTag="504_1_504" lat="43.676949" lon="-79.358536" secsSinceReport="13" predictable="true" heading="64"/>
    <vehicle id="4238" routeTag="504" dirTag="504_0_504" lat="43.643768" lon="-79.403419" secsSinceReport="9" predictable="true" heading="75"/>
    <vehicle id="4246" routeTag="504" dirTag="504_0_504" lat="43.651051" lon="-79.369331" secsSinceReport="13" predictable="true" heading="73"/>
    <vehicle id="4240" routeTag="504" dirTag="504_1_504" lat="43.644535" lon="-79.399948" secsSinceReport="14" predictable="true" heading="255"/>
    <vehicle id="4209" routeTag="504" dirTag="504_1_504" lat="43.639618" lon="-79.446449" secsSinceReport="0" predictable="true" heading="344"/>
    <vehicle id="4224" routeTag="504" dirTag="504_0_504" lat="43.640667" lon="-79.446915" secsSinceReport="15" predictable="true" heading="158"/>
    <vehicle id="4237" routeTag="504" dirTag="504_0_504" lat="43.637665" lon="-79.433601" secsSinceReport="0" predictable="true" heading="74"/>
    <lastTime time="1502115823537"/>
</body>
*/

// LastReported mirrors the `lastTime` element in xml response from vehicleLocations command of TTC real-time API
type LastReported struct {
	MyEpoch int `xml:"time,attr"`
}

// Vehicle mirrors the `vehicle` element in xml response from vehicleLocations command of TTC real-time API
type Vehicle struct {
	VehicleID       string  `xml:"id,attr"`
	RouteTag        string  `xml:"routeTag,attr"`
	DirTag          string  `xml:"dirTag,attr"`
	Lat             float32 `xml:"lat,attr"`
	Lon             float32 `xml:"lon,attr"`
	SecsSinceReport int     `xml:"secsSinceReport,attr"`
	Predictable     bool    `xml:"predictable,attr"`
	Heading         int     `xml:"heading,attr"`
}

// Body mirrors the `body` element in xml response from vehicleLocations command of TTC real-time API
type Body struct {
	Vehicles     []Vehicle      `xml:"vehicle"`
	LastReported []LastReported `xml:"lastTime"` // There is only one lastTime element in response but xml.Unmarshal only reads the time attribute when this is declared as an array ¯\_(ツ)_/¯
}
