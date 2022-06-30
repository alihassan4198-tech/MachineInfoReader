package csvbuilder

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"machine_info_gatherer/common"
	"machine_info_gatherer/model"
	"os"
	"strings"

	"github.com/tidwall/pretty"
)

type CSVWriter struct {
	Writer *csv.Writer
}

func CreateCSVFile(info *model.ComputerInfoType) {
	csvFile, err := os.Create("info.csv")
	if err != nil {
		fmt.Println(err)
	}

	cw := CSVWriter{}
	cw.Writer = csv.NewWriter(csvFile)
	cw.WriteStructInCsv(info)
}

func (cw *CSVWriter) WriteStructInCsv(info *model.ComputerInfoType) {

	jsonInfo, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err)
	}

	prettyJsonInfo := pretty.Pretty(jsonInfo)

	prettyJsonInfoStr := fmt.Sprint(string(prettyJsonInfo))

	cw.JsonToCSVWriter(prettyJsonInfoStr)

}

func (cw *CSVWriter) JsonToCSVWriter(prettyJsonInfoStr string) {
	fmt.Println(prettyJsonInfoStr)

	splitted := strings.Split(prettyJsonInfoStr, "\n")
	splitted = splitted[:len(splitted)-1]
	for i, str := range splitted {
		str = common.RemoveCSVExtras(str)
		if common.DoesStringContainAlphaNumeric(str) && str != "" {
			// count double spaces for csv commas
			tabs := strings.Count(str, "  ")

			// clean string
			str = strings.TrimSpace(str)

			// key-val string
			key, val, _ := strings.Cut(str, ":")
			fmt.Println(i, tabs, key, val)

			var record []string
			for i := 0; i < tabs; i++ {
				record = append(record, "  ")
			}
			record = append(record, key)
			if common.DoesStringContainAlphaNumeric(val) {
				record = append(record, val)
			}

			if err := cw.Writer.Write(record); err != nil {
				log.Fatalln("error writing record to file", err)
			}
		}
	}
}
