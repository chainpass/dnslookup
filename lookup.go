package dnslookup

import (
	"errors"
	"net"
)

type (
	Record struct {
		Name     string
		Priority int
		Value    string
	}
)

var (
	ErrInvalidRecordType = errors.New("invalid record type")
)

func Lookup(recordType string, domain string) ([]Record, error) {
	switch recordType {
	case "CNAME":
		result, err := net.LookupCNAME(domain)
		if err != nil {
			return nil, err
		}

		return []Record{{
			Name:  domain,
			Value: result,
		}}, nil

	case "MX":
		records, err := net.LookupMX(domain)
		if err != nil {
			return nil, err
		}

		var result []Record
		for _, record := range records {
			result = append(result, Record{
				Name:     domain,
				Priority: int(record.Pref),
				Value:    record.Host,
			})
		}

		return result, nil

	case "TXT":
		records, err := net.LookupTXT(domain)
		if err != nil {
			return nil, err
		}

		var result []Record
		for _, record := range records {
			result = append(result, Record{
				Name:  domain,
				Value: record,
			})
		}

		return result, nil
	}

	return nil, ErrInvalidRecordType
}
