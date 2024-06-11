package processor

import (
	"database/sql"
	"log"
)

func SaveMetric(db *sql.DB, metricName string, value float64) {
	insertMetricSQL := `INSERT INTO metrics (metric_name, value) VALUES (?, ?)`
	statement, err := db.Prepare(insertMetricSQL)
	if err != nil {
		log.Fatalf("Failed to prepare statement: %v", err)
	}
	defer statement.Close()

	_, err = statement.Exec(metricName, value)
	if err != nil {
		log.Fatalf("Failed to insert metric: %v", err)
	}
}
