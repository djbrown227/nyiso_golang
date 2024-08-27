Here's a README file in Markdown format for your GitHub project:

```markdown
# NYISO Electricity Data Analysis

This repository contains a Go-based project focused on analyzing New York Independent System Operator (NYISO) electricity load data. The project includes server-side APIs for various data analysis tasks, such as calculating moving averages, detecting anomalies, generating lag features, and computing the rate of change.

## Features

### APIs
- **Moving Average:** Calculates the moving average of the NYISO electricity load data over a specified window.
- **Rate of Change:** Computes the rate of change between consecutive NYISO load data points.
- **Lag Features:** Generates lag features for the NYISO load data to aid in time-series analysis.
- **Anomaly Detection:** Identifies anomalies in the NYISO load data using Z-score thresholding.

### Data Processing
- **Fetch Load Data:** Retrieves electricity load data from a MySQL database, with support for batch processing.
- **Mean and Standard Deviation Calculation:** Computes mean and standard deviation, used for Z-score calculation in anomaly detection.

## Project Structure

```
nyiso_golang/
│
├── pkg/
│   ├── api/
│   │   └── api.go               # Contains HTTP handlers for the API endpoints
│   └── data/
│       ├── feature_engineering.go  # Implements data processing and feature engineering functions
│       └── load_data.go            # Manages database connection and data retrieval
│
└── config/                         # Configuration files for the database connection
```

## How to Run

1. **Clone the Repository**
   ```bash
   git clone https://github.com/djbrown227/nyiso_golang.git
   cd nyiso_golang
   ```

2. **Set Up the MySQL Database**
   - Create a MySQL database and load your NYISO data into it.
   - Configure the database connection settings in the `config` package.

3. **Run the Server**
   ```bash
   go run main.go
   ```
   The server will start on `http://localhost:8080`.

## API Endpoints

- **GET** `/api/load/moving-average`: Returns the moving average of NYISO load data.
- **GET** `/api/load/rate-of-change`: Returns the rate of change in NYISO load data.
- **GET** `/api/load/lag-features`: Returns lag features generated from NYISO load data.
- **GET** `/api/load/anomalies`: Returns detected anomalies in NYISO load data.

## Requirements

- Go 1.16+
- MySQL 5.7+
- `github.com/go-sql-driver/mysql` Go MySQL driver

## Future Work

- **Visualization**: Integrate a frontend for visualizing the processed data.
- **Advanced Analytics**: Implement additional features like predictive modeling and deeper anomaly detection.

## Contributing

Contributions are welcome! Please fork this repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```

This README provides a clear overview of your project, its features, how to set it up, and how to run it, which will be useful for potential employers reviewing your portfolio.
