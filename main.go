package main

import (
    "encoding/json"
    "fmt"
    "log"
    "github.com/gocolly/colly"
)

type Flight struct {
    FlightNumber string `json:"flight_number"`
    Destination  string `json:"destination"`
    Time         string `json:"time"`
    Status       string `json:"status"`
}

func main() {
    c := colly.NewCollector(
        colly.AllowedDomains("www.sabihagokcen.aero"),
    )

    var arrivingFlights []Flight
    var departingFlights []Flight

    // Selectorları belirleyin
    c.OnHTML(".arriving-flights .flight-row", func(e *colly.HTMLElement) {
        flight := Flight{
            FlightNumber: e.ChildText(".flight-number"),
            Destination:  e.ChildText(".flight-destination"),
            Time:         e.ChildText(".flight-time"),
            Status:       e.ChildText(".flight-status"),
        }
        arrivingFlights = append(arrivingFlights, flight)
    })

    c.OnHTML(".departing-flights .flight-row", func(e *colly.HTMLElement) {
        flight := Flight{
            FlightNumber: e.ChildText(".flight-number"),
            Destination:  e.ChildText(".flight-destination"),
            Time:         e.ChildText(".flight-time"),
            Status:       e.ChildText(".flight-status"),
        }
        departingFlights = append(departingFlights, flight)
    })

    // Uçuş bilgileri sayfasını ziyaret et
    c.Visit("https://www.sabihagokcen.aero/yolcu-ve-ziyaretciler/yolcu-rehberi/ucus-bilgi-ekrani")

    // Gelen uçuşları JSON formatında yazdır
    arrivingFlightsJSON, err := json.MarshalIndent(arrivingFlights, "", "  ")
    if err != nil {
        log.Fatalf("Error marshalling arriving flights: %v", err)
    }
    fmt.Println("Arriving Flights:")
    fmt.Println(string(arrivingFlightsJSON))

    // Giden uçuşları JSON formatında yazdır
    departingFlightsJSON, err := json.MarshalIndent(departingFlights, "", "  ")
    if err != nil {
        log.Fatalf("Error marshalling departing flights: %v", err)
    }
    fmt.Println("Departing Flights:")
    fmt.Println(string(departingFlightsJSON))
}
