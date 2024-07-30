package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gdwr/nuget"
)

const PollInterval = 1 * time.Minute

func main() {
	// The cursor is used to only gather nuget events that have happened since the last time we checked.
	cursor := time.Now().UTC()

	// Create a Catalog, which is the interface into reading the NuGet catalog.
	//    Provide your own http client for ratelimiting and cursor to set what you wish to consume.
	nugetCatalog, err := nuget.NewCatalog(http.DefaultClient, cursor)
	handleErr(err)

	println("Polling for updates...")

	for {
		// Process our leaves, in this case printing them out.
		for leaf := range nugetCatalog.StreamLeaves() {
			fmt.Printf("%s version=%s created=%s\n", leaf.PackageId, leaf.Version, leaf.Created)
		}

		// Update the cursor to the current time, so we only get new events next time.
		nugetCatalog.Cursor = time.Now().UTC()

		time.Sleep(PollInterval)
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
