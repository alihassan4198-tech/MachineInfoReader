package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"machine_info_gatherer/csvbuilder"
	"machine_info_gatherer/gatherer"
	"runtime"

	"github.com/wildducktheories/go-csv"
)

func configure(args []string) (*csv.JsonToCsvProcess, error) {
	var baseObject string
	var columns string

	flags := flag.NewFlagSet("json-to-csv", flag.ExitOnError)

	flags.StringVar(&baseObject, "base-object-key", "", "The column into which the remainder of each JSON object is read.")
	flags.StringVar(&columns, "columns", "", "The columns of the CSV file")

	if err := flags.Parse(args); err != nil {
		return nil, err
	}

	if columns == "" {
		return nil, fmt.Errorf("fatal: a --columns parameter must be specified")
	}

	if header, err := csv.Parse(columns); err != nil {
		return nil, fmt.Errorf("fatal: --columns could not be parsed as a CSV record")
	} else {

		return &csv.JsonToCsvProcess{
			BaseObject: baseObject,
			Header:     header,
		}, nil
	}
}

func main() {
	fmt.Println("Launching Machine Info Gatherer...")
	os := runtime.GOOS
	if os == "darwin" { // Mac
		macMachineInfo := gatherer.MacGatherer{}
		macMachineInfo.FetchAllBaseBoard()
		csvbuilder.CreateMacCSVFile(macMachineInfo)
	} else if os == "linux" {
		linuxMachineInfo := gatherer.LinuxGatherer{}
		linuxMachineInfo.FetchAllBaseBoard()
		csvbuilder.CreateLinuxCSVFile(linuxMachineInfo)
	}

	var p *csv.JsonToCsvProcess
	var err error
	if p, err = configure(os.Args[1:]); err == nil {
		errCh := make(chan error, 1)
		p.Run(json.NewDecoder(os.Stdin), csv.WithIoWriter(os.Stdout), errCh)
		err = <-errCh
	}
	if err != nil {
		fmt.Printf("fatal: %v\n", err)
		os.Exit(1)
	}

}
