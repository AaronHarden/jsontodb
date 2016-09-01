package jsontodb

import (
	"os"
	"strings"
	"testing"
)

func TestJSONPQ(t *testing.T) {
	expectedDriver := postgres
	expectedDataSource := "dbname=test host=127.0.0.1 user=test"
	file, err := os.Open("test_pq.json")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	driverName, dataSourceName := pqFromJSON(file)
	compareDriver(t, driverName, expectedDriver)
	compareDataSource(t, dataSourceName, expectedDataSource)
}

func TestEmptyPQ(t *testing.T) {
	expectedDriver := postgres
	expectedDataSource := ""
	driverName, dataSourceName := pqFromJSON(strings.NewReader("{}"))
	compareDriver(t, driverName, expectedDriver)
	compareDataSource(t, dataSourceName, expectedDataSource)
}

func compareDriver(t *testing.T, driverName, expectedDriver string) {
	if driverName != expectedDriver {
		t.Errorf("Expected driver to be:\n%v\nReceived:\n%v", expectedDriver, driverName)
	}
}

func compareDataSource(t *testing.T, dataSourceName, expectedDataSource string) {
	if dataSourceName != expectedDataSource {
		t.Errorf("Expected data source to be:\n%v\nReceived:\n%v", expectedDataSource, dataSourceName)
	}
}
