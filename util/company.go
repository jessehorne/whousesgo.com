package util

import (
	"encoding/csv"
	"os"
)

type Company struct {
	Name     string
	Website  string
	Location string
}

const (
	CompanyCSVName     = 0
	CompanyCSVWebsite  = 1
	CompanyCSVLocation = 2
)

func GetAllCompanies() ([]Company, error) {
	csvFile, err := os.Open("./companies/active.csv")
	if err != nil {
		return []Company{}, err
	}
	defer csvFile.Close()
	csvFile.Seek(0, 0)

	reader := csv.NewReader(csvFile)
	reader.Comma = '|'
	rawData, err := reader.ReadAll()
	if err != nil {
		return []Company{}, err
	}

	var companies []Company

	for i, record := range rawData {
		if i == 0 {
			continue
		}

		newCompany := Company{
			Name:     record[CompanyCSVName],
			Website:  record[CompanyCSVWebsite],
			Location: record[CompanyCSVLocation],
		}

		companies = append(companies, newCompany)
	}

	return companies, nil
}
