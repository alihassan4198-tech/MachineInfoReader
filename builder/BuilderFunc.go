package builder

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	jsonFile, err := os.Create("" + fileName + ".json")
	if err != nil {
		fmt.Println(err)
	}
	cw := CSVWriter{}
	cw.Writer = csv.NewWriter(jsonFile)
	j := cw.WriteStructInJson(info)
	jsonFile.WriteString(j)

}

func readJSONFile(filename string) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readJSON(f)
}

func CreateCSVFile(JsonFileName string, args ...interface{}) {

	data, err := readJSONFile("./" + JsonFileName + ".json")
	if err != nil {
		fmt.Println(err)
	}

	r, err := json2csv.JSON2CSV(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(r) == 0 {
		return
	}

	headerStyle := headerStyleTable["jsonpointer"]
	myscvfile, err := os.Create("" + JsonFileName + ".csv")
	if err != nil {
		log.Fatal(err)
	}
	err = printCSV(myscvfile, r, headerStyle, false)
	if err != nil {
		log.Fatal(err)
	}
	if JsonFileName == "ComputerSoftwaresInstalled" {

		mycswriter := csv.NewWriter(myscvfile)
		vals := []string{"total", "500"}
		_ = mycswriter.Write([]string{})
		_ = mycswriter.Write(vals)
		mycswriter.Flush()
		myscvfile.Close()

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
