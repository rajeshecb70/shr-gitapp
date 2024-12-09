package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
)

// UptimeResponse represents system uptime response structure
type UptimeResponse struct {
	Uptime string `json:"uptime"`
}

// CPUResponse represents CPU usage response structure
type CPUResponse struct {
	CPUUsage float64 `json:"cpu_usage"`
}

// LoadResponse represents system load response structure
type LoadResponse struct {
	Load1  float64 `json:"load_1"`
	Load5  float64 `json:"load_5"`
	Load15 float64 `json:"load_15"`
}

// writeJSON is a helper function to write a JSON response and handle errors.
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to write JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// uptimeHandler provides the system uptime in JSON format.
func uptimeHandler(w http.ResponseWriter, r *http.Request) {
	hostInfo, err := host.Info()
	if err != nil {
		log.Printf("Failed to get host info: %v", err)
		http.Error(w, "Failed to get uptime", http.StatusInternalServerError)
		return
	}

	uptime := time.Duration(hostInfo.Uptime) * time.Second
	response := UptimeResponse{Uptime: uptime.String()}
	writeJSON(w, http.StatusOK, response)
}

// cpuHandler provides CPU usage in JSON format.
func cpuHandler(w http.ResponseWriter, r *http.Request) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		log.Printf("Failed to get CPU usage: %v", err)
		http.Error(w, "Failed to get CPU usage", http.StatusInternalServerError)
		return
	}

	response := CPUResponse{CPUUsage: percentages[0]}
	writeJSON(w, http.StatusOK, response)
}

// loadHandler provides system load in JSON format.
func loadHandler(w http.ResponseWriter, r *http.Request) {
	loadStat, err := load.Avg()
	if err != nil {
		log.Printf("Failed to get system load: %v", err)
		http.Error(w, "Failed to get system load", http.StatusInternalServerError)
		return
	}

	response := LoadResponse{
		Load1:  loadStat.Load1,
		Load5:  loadStat.Load5,
		Load15: loadStat.Load15,
	}
	writeJSON(w, http.StatusOK, response)
}

func main() {
	http.HandleFunc("/uptime", uptimeHandler)
	http.HandleFunc("/cpu", cpuHandler)
	http.HandleFunc("/load", loadHandler)

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	fmt.Printf("Starting server at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
