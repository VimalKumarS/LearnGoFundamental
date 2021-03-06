package main

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const watchedPath = "./source"

func main() {

	for {
		d, _ := os.Open(watchedPath)
		files, _ := d.Readdir(-1)
		for _, fi := range files {
			filepath := watchedPath + "/" + fi.Name()
			f, _ := os.Open(filepath)
			data, _ := ioutil.ReadAll(f)
			f.Close()
			os.Remove(filepath)

			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					invoice := new(Invoice)
					invoice.Number = r[0]
					invoice.Amount, _ = strconv.ParseFloat(r[1], 64)
					invoice.PurchaseOrderNumber, _ = strconv.Atoi(r[2])
					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					invoice.InvoiceDate = time.Unix(unixTime, 0)

				}
			}(string(data))
		}
	}
}

type Invoice struct {
	Number              string
	Amount              float64
	PurchaseOrderNumber int
	InvoiceDate         time.Time
}
