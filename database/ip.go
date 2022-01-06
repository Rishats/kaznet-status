package database

import (
	"fmt"
	"kaznet-status/pkg/utils"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

func CheckIfIPFromKZ(geoip *geoip2.City) {
	if geoip.Country.IsoCode == "KZ" {

		type Info struct {
			CityName string
			GeoIp    *geoip2.City
		}

		templateData := Info{
			CityName: geoip.City.Names["en"],
			GeoIp:    geoip,
		}
		go utils.Notify("ip.gohtml", nil, templateData)

		fmt.Printf("FROM KAZNET! INTERNET AVAILABLE \n")
		fmt.Printf("Portuguese (BR) city name: %v\n", geoip.City.Names["en"])
		if len(geoip.Subdivisions) > 0 {
			fmt.Printf("English subdivision name: %v\n", geoip.Subdivisions[0].Names["en"])
		}
		fmt.Printf("Russian country name: %v\n", geoip.Country.Names["ru"])
		fmt.Printf("ISO country code: %v\n", geoip.Country.IsoCode)
		fmt.Printf("Time zone: %v\n", geoip.Location.TimeZone)
		fmt.Printf("Coordinates: %v, %v\n", geoip.Location.Latitude, geoip.Location.Longitude)
	} else {
		fmt.Printf("NO: ")
	}
}

func GetIpInfo(iip string) *geoip2.City {
	db, err := geoip2.Open("database/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(iip)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	return record
}

func CheckAllRegionsStatus() {
	ips := []string{
		//"8.8.4.4",
		"188.0.151.149", // Almaty STS
		"91.185.6.14",   // Almaty JST Telecom
		"2.133.19.254",
	}

	for _, ip := range ips {
		lastIpFromTraceroute := TracerouteLastHoopIP(ip)
		geoip := GetIpInfo(lastIpFromTraceroute)
		CheckIfIPFromKZ(geoip)
	}

}

func TracerouteLastHoopIP(destIP string) string {
	ipHops := trace(destIP)

	return ipHops[len(ipHops)-1].String()
}
