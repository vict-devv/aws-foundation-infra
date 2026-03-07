package health

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

type Response struct {
	Status    string    `json:"status"`
	Uptime    string    `json:"uptime"`
	GoVersion string    `json:"go_version"`
	Timestamp time.Time `json:"timestamp"`
}

var startTime = time.Now()

// Handler responds with the health and other application information..
func Handler(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Status:    "UP",
		Uptime:    time.Since(startTime).String(),
		GoVersion: runtime.Version(),
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
