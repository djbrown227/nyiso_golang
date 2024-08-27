// pkg/data/feature_engineering.go
package data

import (
    "math"
    "nyiso-electricity-data-analysis/pkg/models"
)

// Moving Average
func CalculateMovingAverage(data []models.LoadData, windowSize int) []float64 {
    var movingAverages []float64
    for i := range data {
        if i < windowSize-1 {
            movingAverages = append(movingAverages, math.NaN())
            continue
        }
        sum := 0.0
        for j := 0; j < windowSize; j++ {
            sum += data[i-j].Nyiso
        }
        movingAverages = append(movingAverages, sum/float64(windowSize))
    }
    return movingAverages
}

// Rate of Change
func CalculateRateOfChange(data []models.LoadData) []float64 {
    var rates []float64
    for i := 1; i < len(data); i++ {
        rate := (data[i].Nyiso - data[i-1].Nyiso) / data[i-1].Nyiso
        rates = append(rates, rate)
    }
    return rates
}

// Lag Features
func CreateLagFeatures(data []models.LoadData, lag int) []float64 {
    var lagFeatures []float64
    for i := range data {
        if i < lag {
            lagFeatures = append(lagFeatures, math.NaN())
            continue
        }
        lagFeatures = append(lagFeatures, data[i-lag].Nyiso)
    }
    return lagFeatures
}

// Anomaly Detection (using Z-scores)
func DetectAnomalies(data []models.LoadData, threshold float64) []bool {
    var anomalies []bool
    mean, std := calculateMeanAndStdDev(data)
    for _, v := range data {
        zScore := (v.Nyiso - mean) / std
        anomalies = append(anomalies, math.Abs(zScore) > threshold)
    }
    return anomalies
}

// Helper function to calculate mean and standard deviation
func calculateMeanAndStdDev(data []models.LoadData) (mean float64, std float64) {
    var sum float64
    for _, v := range data {
        sum += v.Nyiso
    }
    mean = sum / float64(len(data))

    var varianceSum float64
    for _, v := range data {
        varianceSum += math.Pow(v.Nyiso-mean, 2)
    }
    std = math.Sqrt(varianceSum / float64(len(data)))
    return
}
