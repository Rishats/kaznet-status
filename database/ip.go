package database

import (
	"github.com/hashicorp/go-memdb"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

func CheckIfIPFromKZ(ip string, lastIpFromTraceroute string, geoip *geoip2.City, memDb *memdb.MemDB) {
	var citySubdivision string
	if len(geoip.Subdivisions) > 0 {
		citySubdivision = geoip.Subdivisions[0].Names["en"]
	} else {
		citySubdivision = "not set"
	}

	var cityName string
	if len(geoip.City.Names["en"]) != 0 {
		cityName = geoip.City.Names["en"]
	} else {
		cityName = "not set"
	}

	if geoip.Country.IsoCode == "KZ" {
		// Create a write transaction
		txn := memDb.Txn(true)
		// Insert some ips in to memDb
		if err := txn.Insert("ip", &IP{ip, lastIpFromTraceroute, cityName, citySubdivision, geoip.Location.Longitude, geoip.Location.Latitude, 1}); err != nil {
			panic(err)
		}
		// Commit the transaction
		txn.Commit()
	} else if geoip.Country.IsoCode != "KZ" && ip == lastIpFromTraceroute {
		// Create a write transaction
		txn := memDb.Txn(true)
		// Insert some ips in to memDb
		if err := txn.Insert("ip", &IP{ip, lastIpFromTraceroute, cityName, citySubdivision, geoip.Location.Longitude, geoip.Location.Latitude, 0}); err != nil {
			panic(err)
		}
		// Commit the transaction
		txn.Commit()
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

func CheckAllRegionsStatus(memDb *memdb.MemDB) {
	ips := []string{
		//"Almaty",
		"188.0.151.149", // Almaty STS
		"178.91.21.193", // JSC
		"2.133.20.69",   // JSC
		"217.76.64.24",  // ALTEL
		"37.99.46.127",  // KAR-TEL-AS - Kar-Tel LLC
		//"Astana",
		"2.133.19.254", // Random Astana
		"5.34.45.254",
		"37.17.180.126",
		"89.218.79.230",
	}

	for _, ip := range ips {
		lastIpFromTraceroute := TracerouteLastHoopIP(ip)
		geoip := GetIpInfo(lastIpFromTraceroute)
		CheckIfIPFromKZ(ip, lastIpFromTraceroute, geoip, memDb)
	}

}

func TracerouteLastHoopIP(destIP string) string {
	ipHops := trace(destIP)

	return ipHops[len(ipHops)-1].String()
}
