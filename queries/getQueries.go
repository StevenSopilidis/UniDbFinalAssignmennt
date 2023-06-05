package queries

import (
	"log"
	"os"
)

func GetQueries() map[string]string {
	queryA, error := os.ReadFile("./queries/Query_a.sql")
	if error != nil {
		log.Panicln(error.Error())
	}
	queryB, error := os.ReadFile("./queries/Query_b.sql")
	if error != nil {
		log.Panicln(error.Error())
	}
	queryC, error := os.ReadFile("./queries/Query_c.sql")
	if error != nil {
		log.Panicln(error.Error())
	}
	queryD, error := os.ReadFile("./queries/Query_d.sql")
	if error != nil {
		log.Panicln(error.Error())
	}
	queries := make(map[string]string)
	queries["queryA"] = string(queryA)
	queries["queryB"] = string(queryB)
	queries["queryC"] = string(queryC)
	queries["queryD"] = string(queryD)
	return queries
}
