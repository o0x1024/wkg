package portService

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/Ullaakut/nmap/v2"

	"time"
)

func TestPortScan(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5 minute timeout.


	scanner, err := nmap.NewScanner(
		nmap.WithPorts("80,443,843,9091"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Println("unable to create nmap scanner: ", err)
	}

	target := []string{"hotwheels.momo.com", "college.sf-express.com", "sfim.sf-express.com"}
	for _, v := range target {
		scanner.AddOptions(nmap.WithTargets(v))
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		log.Println("unable to run nmap scan: ", err)
	}

	if warnings != nil {
		log.Println("Warnings: \n ", warnings)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	fmt.Printf("Nmap done: %d hosts up scanned in %3f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)

}
