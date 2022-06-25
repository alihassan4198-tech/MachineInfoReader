package csvbuilder

import (
	"fmt"
	"machine_info_gatherer/gatherer"

	"github.com/jszwec/csvutil"
)

func CreateCSVFile(info gatherer.BaseGatherer) {
	b, err := csvutil.Marshal(info)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
