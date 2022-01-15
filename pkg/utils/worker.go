package utils

import (
	"fmt"
	"github.com/alecthomas/template"
	"github.com/hashicorp/go-memdb"
	"kaznet-status/database"

	"log"
)

func notify(templateFileName string, templateFuncMap template.FuncMap, data interface{}) {
	text, err := GetTemplate(templateFileName, templateFuncMap, data)
	if err != nil {
		log.Panic(err)
	}

	SendToTelegram(text)
}

func sendIpsWithoutInternetToTelegram(ipsWithoutInternet []database.IP) {
	var ipsWithoutInternetText string

	for _, ipWithoutInternet := range ipsWithoutInternet {
		ipWithoutInternetText := fmt.Sprintf(
			"\n IP: %s \n Город: %s \n\n",
			ipWithoutInternet.IP,
			ipWithoutInternet.City,
		)

		ipsWithoutInternetText = ipsWithoutInternetText + ipWithoutInternetText
	}

	type Info struct {
		Text string
	}

	templateData := Info{
		Text: ipsWithoutInternetText,
	}

	go notify("ipsWithoutInternet.gohtml", nil, templateData)
}

func InitWorker(memDb *memdb.MemDB) {
	getCityWithInternet(memDb)

	ipsWithoutInternet := getCityWithoutInternet(memDb)
	if len(ipsWithoutInternet) > 0 {
		sendIpsWithoutInternetToTelegram(ipsWithoutInternet)
	}
}

func getCityWithInternet(memDb *memdb.MemDB) []database.IP {
	// Create read-only transaction
	txn := memDb.Txn(false)
	defer txn.Abort()

	// List all the IPs
	it, err := txn.Get("ip", "id")
	if err != nil {
		panic(err)
	}

	fmt.Println("IPs with Internet (ok):")
	var ips []database.IP
	for obj := it.Next(); obj != nil; obj = it.Next() {
		ipData := obj.(*database.IP)
		if ipData.Status == 1 {
			ips = append(ips, *ipData)
			go ChangeIpData(ipData)
			fmt.Printf("%s %s %s  %b\n", ipData.IP, ipData.Lip, ipData.City, ipData.Status)
		}
	}

	fmt.Println("============================")

	return ips
}

func getCityWithoutInternet(memDb *memdb.MemDB) []database.IP {
	// Create read-only transaction
	txn := memDb.Txn(false)
	defer txn.Abort()

	// List all the IPs
	it, err := txn.Get("ip", "id")
	if err != nil {
		panic(err)
	}

	fmt.Println("IPs without Internet (ERROR):")
	var ips []database.IP
	for obj := it.Next(); obj != nil; obj = it.Next() {
		ipData := obj.(*database.IP)
		if ipData.Status == 0 {
			ips = append(ips, *ipData)
			go ChangeIpData(ipData)
			fmt.Printf("%s %s %s  %b\n", ipData.IP, ipData.Lip, ipData.City, ipData.Status)
		}
	}

	fmt.Println("============================")
	return ips
}
