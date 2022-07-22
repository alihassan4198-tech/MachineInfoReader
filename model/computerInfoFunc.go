package model

import (
	"encoding/csv"
	"os"
	"strconv"
)

func (cs *ComputerSoftwaresInstalledType) AppendAllMapsInCSV(myscvfile *os.File) {

	mycswriter := csv.NewWriter(myscvfile)
	arg := cs.Total_software
	// numOfSoft := arg.(int)
	str := strconv.Itoa(arg)
	vals := []string{"/Total Softwares Installed"}
	vals1 := []string{str}
	_ = mycswriter.Write([]string{})
	_ = mycswriter.Write(vals)
	_ = mycswriter.Write(vals1)
	mycswriter.Flush()
	myscvfile.Close()

}

func (cs *ComputerCPUType) AppendAllMapsInCSV(myscvfile *os.File) {

	mycswriter := csv.NewWriter(myscvfile)
	arg := cs.Caption
	vals := []string{"/Caption"}
	vals1 := []string{arg}
	_ = mycswriter.Write([]string{})
	_ = mycswriter.Write(vals)
	_ = mycswriter.Write(vals1)
	mycswriter.Flush()
	myscvfile.Close()

}

func (cs *ComputerFirewallRulesType) AppendAllMapsInCSV(myscvfile *os.File) {

	mycswriter := csv.NewWriter(myscvfile)
	activeState := cs.Active_state
	totalRules := cs.Total_rules
	totalRulesNum := strconv.Itoa(totalRules)
	vals := []string{"/Active State", "/Total Rules"}
	vals1 := []string{activeState, totalRulesNum}
	_ = mycswriter.Write([]string{})
	_ = mycswriter.Write(vals)
	_ = mycswriter.Write(vals1)
	mycswriter.Flush()
	myscvfile.Close()

}

func (cs *ComputerServicesType) AppendAllMapsInCSV(myscvfile *os.File) {

	mycswriter := csv.NewWriter(myscvfile)
	arg := cs.TotalServciesRunning
	str := strconv.Itoa(arg)
	vals := []string{"/Total Servcies Running"}
	vals1 := []string{str}
	_ = mycswriter.Write([]string{})
	_ = mycswriter.Write(vals)
	_ = mycswriter.Write(vals1)
	mycswriter.Flush()
	myscvfile.Close()

}
