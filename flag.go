package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0, "database port")

	/*
		dipake flag.Parse() biar waktu go run flag.go ambil dari yang diinput
		contoh jalanin ini
		go run .\flag.go -username=Tirta -password="rahasia banget" -host=123.123.23.3 -port=5505
		kalo tanpa parse, akan ambil defail value
	*/
	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)

}
