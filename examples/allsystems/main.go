package main

import (
	"context"
	"flag"
	"fmt"
	v1 "github.com/merlindorin/go-airzone/api/v1"
	"github.com/merlindorin/go-airzone/pkg"
	"go.uber.org/zap"
	"net/url"
	"os"
)

func Usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Simple command line for getting all Airzone Systems\n\n")
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	level := zap.LevelFlag("level", zap.InfoLevel, "set log level (debug|info|warn|error|panic|fatal)")
	baseURL := flag.String("baseurl", "http://192.168.1.235:3000", "Base URL")

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

	err = cl.V1().System.List(context.TODO(), &ss)
	if err != nil {
		panic(err)
	}

	fmt.Println("| SystemID | Manufacturer         | SystemFirmware |")
	fmt.Println("|----------|----------------------|----------------|")
	for _, s := range ss {
		fmt.Printf("| %8.d | %-20s | %-14s |\n", s.SystemID, s.Manufacturer, s.SystemFirmware)
	}
}
