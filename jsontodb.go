package jsontodb

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

const postgres = "postgres"

func sortKeys(m map[string]string) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}

func jsonToMap(in io.Reader) (dbArguments map[string]string) {
	dec := json.NewDecoder(in)
	err := dec.Decode(&dbArguments)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func JSONToPQ(in io.Reader) (*sql.DB, error) {
	return sql.Open(pqFromJSON(in))
}

func pqFromJSON(in io.Reader) (driverName, dataSourceName string) {
	driverName = postgres
	dbArguments := jsonToMap(in)
	keys := sortKeys(dbArguments)
	buf := bytes.Buffer{}
	for _, k := range keys {
		fmt.Fprintf(&buf, "%s=%s ", k, dbArguments[k])
	}
	dataSourceName = strings.TrimSpace(buf.String())
	return
}
