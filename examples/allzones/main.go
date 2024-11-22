package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"

	v1 "github.com/merlindorin/go-airzone/api/v1"
	"github.com/merlindorin/go-airzone/pkg"
	"go.uber.org/zap"
)

func Usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Simple command line for getting all Airzone Zones in all Systems\n\n")
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	level := zap.LevelFlag("level", zap.InfoLevel, "set log level (debug|info|warn|error|panic|fatal)")
	baseURL := flag.String("baseurl", "http://192.168.1.235:3000", "Base URL")

	flag.Usage = Usage
	flag.Parse()

	l, err := zap.NewDevelopment(zap.IncreaseLevel(level))
	if err != nil {
		panic(err)
	}

	u, err := url.Parse(*baseURL)
	if err != nil {
		panic(err)
	}

	cl := pkg.NewClient(u, l)

	ss := []v1.System{}
	allZones := []v1.Zone{}

	err = cl.V1().System.List(context.TODO(), &ss)
	if err != nil {
		panic(err)
	}

	for _, s := range ss {
		zs := []v1.Zone{}

		err = cl.V1().Zone.List(context.TODO(), s.SystemID, &zs)
		if err != nil {
			panic(err)
		}

		allZones = append(allZones, zs...)
	}

	fmt.Println("| ZoneID   | SystemID | Name         | Setpoint | RoomTemp | Humidity | Sleep |")
	fmt.Println("|----------|----------|--------------|----------|----------|----------|-------|")
	for _, z := range allZones {
		fmt.Printf("| %8.d | %8.d | %-12s | %6.1f°C | %6.1f°C | %7.f%% | %5d |\n", z.ZoneID, z.SystemID, z.Name, z.Setpoint, z.RoomTemp, z.Humidity, z.Sleep)
	}
}
