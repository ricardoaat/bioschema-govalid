package validator

import (
	"encoding/csv"
	"net/http"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

//LoadFile loads CSV file into an spec object
func LoadFile(fp string) {
	re := regexp.MustCompile(`.*\.csv$`)
	isCsv := re.MatchString(fp)
	if isCsv {
		f, err := os.Open(fp)
		if err != nil {
			log.Panic(err)
		}
		defer f.Close()

		parseCSV(csv.NewReader(f))
	} else {
		log.WithFields(log.Fields{
			"file": fp,
		}).Error("Sorry, that doesn't looks like a CSV file")
	}
}

//LoadURL loads CSV from url into an spec object
func LoadURL(u string) {
	res, err := http.Get(u)
	if err != nil {
		log.Panic("Error while downloading", u, "-", err)
		return
	}
	defer res.Body.Close()

	parseCSV(csv.NewReader(res.Body))
}

func parseCSV(r *csv.Reader) {

	r.FieldsPerRecord = 7

	csv, err := r.ReadAll()
	if err != nil {
		log.Panic("An error has occurred trying to parse the CVS file ", err)
	}

	if len(csv) <= 1 {
		log.Error("Empty file, please check its content")
	}
	log.Debug("CSV lenght ", len(csv))

	specificationParams := make([]SpecificationParam, 0)
	for i, row := range csv {
		if i <= 1 {
			log.Debug("Skipped ", row)
			continue
		}
		s := SpecificationParam{
			strings.Replace(row[0], "\n", "", -1),
			strings.Replace(row[1], "\n", "", -1),
			strings.Replace(row[2], "\n", "", -1),
			strings.Replace(row[3], "\n", "", -1),
			strings.Replace(row[4], "\n", "", -1),
			strings.Replace(row[5], "\n", "", -1),
			strings.Replace(row[6], "\n", "", -1),
		}
		specificationParams = append(specificationParams, s)
	}

	Specifications = append(Specifications, Specification{"le name", specificationParams})

	d, err := yaml.Marshal(Specifications[0].SpecificationParams)
	if err != nil {
		log.Panic(err)
	}
	log.Info(string(d))

}
