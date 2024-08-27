// pkg/data/load_data.go
package data

import (
    "database/sql"
    "fmt"
    "time"
    "nyiso-electricity-data-analysis/config"
    "nyiso-electricity-data-analysis/pkg/models"
    _ "github.com/go-sql-driver/mysql"
)

func NewDBConnection(cfg config.Config) (*sql.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}

func FetchLoadData(db *sql.DB, offset, batchSize int) ([]models.LoadData, error) {
    query := fmt.Sprintf("SELECT id, time_stamp, capitl, centrl, dunwod, genese, hud_vl, longil, mhk_vl, millwd, n_y_c, north, west, nyiso FROM load_data LIMIT %d OFFSET %d", batchSize, offset)
    rows, err := db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error executing query: %v", err)
    }
    defer rows.Close()

    var data []models.LoadData
    for rows.Next() {
        var ld models.LoadData
        var timeStampStr string
        err := rows.Scan(&ld.ID, &timeStampStr, &ld.Capital, &ld.Central, &ld.Dunwood, &ld.Genesee, &ld.Hudvl, &ld.Longil, &ld.Mhkvl, &ld.Millwd, &ld.Nyc, &ld.North, &ld.West, &ld.Nyiso)
        if err != nil {
            return nil, fmt.Errorf("error scanning row: %v", err)
        }

        // Parse the time stamp string to time.Time
        ld.Timestamp, err = time.Parse("2006-01-02 15:04:05", timeStampStr)
        if err != nil {
            return nil, fmt.Errorf("error parsing timestamp: %v", err)
        }
        data = append(data, ld)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }

    return data, nil
}
