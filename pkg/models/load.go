// pkg/models/load.go
package models

import (
    "time"
)

// LoadData represents a row in the load_data table.
type LoadData struct {
    ID        int       `json:"id"`
    Timestamp time.Time `json:"time_stamp"`
    Capital   float64   `json:"capitl"`
    Central   float64   `json:"centrl"`
    Dunwood   float64   `json:"dunwod"`
    Genesee   float64   `json:"genese"`
    Hudvl     float64   `json:"hud_vl"`
    Longil    float64   `json:"longil"`
    Mhkvl     float64   `json:"mhk_vl"`
    Millwd    float64   `json:"millwd"`
    Nyc       float64   `json:"n_y_c"`
    North     float64   `json:"north"`
    West      float64   `json:"west"`
    Nyiso     float64   `json:"nyiso"`
}
