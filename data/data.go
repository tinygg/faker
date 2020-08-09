package data

import (
	"github.com/tinygg/gofaker/data/person"
	"reflect"
)

// Data consists of the main set of fake information
func Data() map[string]map[string][]string {
	return map[string]map[string][]string{
		"person":    person.Provider(reflect.ValueOf(LOCALE).String()),
		"address":   Address,
		"company":   Company,
		"job":       Job,
		"lorem":     Lorem,
		"language":  Languages,
		"internet":  Internet,
		"file":      Files,
		"color":     Colors,
		"computer":  Computer,
		"hipster":   Hipster,
		"beer":      Beer,
		"hacker":    Hacker,
		"animal":    Animal,
		"currency":  Currency,
		"log_level": LogLevels,
		"timezone":  TimeZone,
		"car":       Car,
		"emoji":     Emoji,
		"word":      Word,
		"food":      Food,
	}
}

// IntData consists of the main set of fake information (integer only)
var IntData = map[string]map[string][]int{
	"status_code": StatusCodes,
}
