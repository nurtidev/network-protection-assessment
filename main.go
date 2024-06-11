package main

import (
	"log"

	"github.com/nurtidev/network-protection-assessment/internal/console"
	"github.com/nurtidev/network-protection-assessment/internal/exporter"
	"github.com/nurtidev/network-protection-assessment/internal/interceptor"
	"github.com/nurtidev/network-protection-assessment/pkg/db"
)

func main() {
	database := db.InitDB("metrics.db")
	defer database.Close()

	device := "\\Device\\NPF_{402C6739-E35A-40A2-A62C-202DB58DE6CA}"

	go func() {
		err := interceptor.StartCapture(device)
		if err != nil {
			log.Fatalf("Failed to start capture: %v", err)
		}
	}()

	go exporter.RecordMetrics()
	console.StartConsole()
}
