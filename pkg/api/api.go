// pkg/api/api.go
package api

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "nyiso-electricity-data-analysis/pkg/data"
)

func StartServer(db *sql.DB) {
    http.HandleFunc("/api/load/moving-average", func(w http.ResponseWriter, r *http.Request) {
        dataBatch, err := data.FetchLoadData(db, 0, 1000)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        movingAverage := data.CalculateMovingAverage(dataBatch, 10)
        json.NewEncoder(w).Encode(movingAverage)
    })

    http.HandleFunc("/api/load/rate-of-change", func(w http.ResponseWriter, r *http.Request) {
        dataBatch, err := data.FetchLoadData(db, 0, 1000)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        rateOfChange := data.CalculateRateOfChange(dataBatch)
        json.NewEncoder(w).Encode(rateOfChange)
    })

    http.HandleFunc("/api/load/lag-features", func(w http.ResponseWriter, r *http.Request) {
        dataBatch, err := data.FetchLoadData(db, 0, 1000)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        lagFeatures := data.CreateLagFeatures(dataBatch, 5)
        json.NewEncoder(w).Encode(lagFeatures)
    })

    http.HandleFunc("/api/load/anomalies", func(w http.ResponseWriter, r *http.Request) {
        dataBatch, err := data.FetchLoadData(db, 0, 1000)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        anomalies := data.DetectAnomalies(dataBatch, 2.0)
        json.NewEncoder(w).Encode(anomalies)
    })

    log.Println("Starting server on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
