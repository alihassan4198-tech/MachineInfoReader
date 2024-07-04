package builder

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"machine_info_gatherer/common"
	"machine_info_gatherer/model"
	"os/user"
	"path/filepath"

	// "machine_info_gatherer/distro/systemprofiler"

	"os"

	"github.com/tidwall/pretty"
	"github.com/yukithm/json2csv"
)

type CSVWriter struct {
	Writer *csv.Writer
}

var headerStyleTable = map[string]json2csv.KeyStyle{
	"jsonpointer": json2csv.JSONPointerStyle,
	"slash":       json2csv.SlashStyle,
	"dot":         json2csv.DotNotationStyle,
	"dot-bracket": json2csv.DotBracketStyle,
}

func (cw *CSVWriter) WriteStructInJson(info interface{}) string {

	jsonInfo, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err)
	}
	prettyJsonInfo := pretty.Pretty(jsonInfo)
	prettyJsonInfoStr := fmt.Sprint(string(prettyJsonInfo))
	return prettyJsonInfoStr
}

func CreateJsonFile(info interface{}, fileName string) {
	jsonString, err := json.Marshal(info)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return
	}
	homeDir := currentUser.HomeDir
	dir := filepath.Join(homeDir, "machinedataJsonInfo")
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	filePath := filepath.Join(dir, fileName+".json")
	err = os.WriteFile(filePath, jsonString, os.ModePerm)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Printf("JSON data written to %s\n", filePath)
}

func ReadJSONFile(filename string) (interface{}, error) {

	f, err := os.Open(common.PathGetter() + filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readJSON(f)
}

func CreateCSVFile(JsonFileName string, Info *model.ComputerInfoType) {

	data, err := ReadJSONFile(JsonFileName + ".json")
	if err != nil {
		fmt.Println(err)
	}

	r, err := json2csv.JSON2CSV(data)
	if err != nil {
		fmt.Println(err)
	}

	// if len(r) == 0 {
	// 	return
	// }

	headerStyle := headerStyleTable["jsonpointer"]
	myscvfile, err := os.Create(common.PathGetter() + Info.ComputerBaseboard.Computer_name + "_" + JsonFileName + ".csv")

	if err != nil {
		log.Fatal(err)
	}
	err = printCSV(myscvfile, r, headerStyle, false)
	if err != nil {
		log.Fatal(err)
	}

	switch JsonFileName {
	case "ComputerCPU":
		Info.ComputerCPU.AppendAllMapsInCSV(myscvfile)
	case "ComputerSoftwaresInstalled":
		Info.ComputerSoftwaresInstalled.AppendAllMapsInCSV(myscvfile)
	case "ComputerFirewallRules":
		Info.ComputerFirewallRules.AppendAllMapsInCSV(myscvfile)
	case "ComputerServices":
		Info.ComputerServices.AppendAllMapsInCSV(myscvfile)
	}

	e := os.Remove(common.PathGetter() + JsonFileName + ".json")
	if e != nil {
		fmt.Println(e)
	}
}

func readJSON(r io.Reader) (interface{}, error) {

	decoder := json.NewDecoder(r)
	decoder.UseNumber()

	var data interface{}
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func printCSV(w io.Writer, results []json2csv.KeyValue, headerStyle json2csv.KeyStyle, transpose bool) error {

	csv := json2csv.NewCSVWriter(w)
	csv.HeaderStyle = headerStyle
	csv.Transpose = transpose
	if err := csv.WriteCSV(results); err != nil {
		return err
	}
	return nil
}
