package builder

import "os"

type ICsvCreator interface {
	AppendAllMapsInCSV(file *os.File)
}
