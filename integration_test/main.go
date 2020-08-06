package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ccin2p3/go-puppetdb"
)

func main() {
	client := puppetdb.NewClient("localhost", 8080, true)
	query := `resources[title, type, parameters] { exported = true and type = "Nagios_host" and tags = "realm_production" }`

	rawBody, err := client.PQLRawQuery(query)
	if err != nil {
		panic(err)
	}
	defer rawBody.Close()

	b, err := ioutil.ReadAll(rawBody)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(b))
}
