# Automate System

## Overview
This project automates filling template fields based on an Excel sheet. It reads an Excel file, extracts data, and dynamically fills placeholders in a predefined template.

## Features
- Upload an Excel file with structured data.
- Process the template by replacing placeholders with Excel data.
- API endpoints to handle file uploads and data processing.

## Prerequisites
Ensure you have the following installed:
- [Go](https://go.dev/doc/install) (1.17+ recommended)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Excelize](https://github.com/xuri/excelize) for handling Excel files

## Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/your-repo/automate-system.git
   cd automate-system
   ```
2. Initialize Go modules:
   ```sh
   go mod init AutomateSystem
   ```
3. Install dependencies:
   ```sh
   go get github.com/gin-gonic/gin
   go get github.com/xuri/excelize/v2
   go mod tidy
   ```

## Project Structure
```
AutomateSystem/
│── Controllers/
│   ├── controllers.go  # Handles API logic
│── Models/
│   ├── userData.go     # Defines data structures
│── Templates/
│   ├── email_template.txt  # Example template file
│── Uploads/          # Directory for uploaded files
│── main.go          # Application entry point
│── go.mod           # Go module dependencies
│── README.md        # Project documentation
```

## API Endpoints
### 1. Upload an Excel File
**Endpoint:** `POST /upload`

**Usage:**
```sh
curl -X POST -F "file=@sample.xlsx" http://localhost:8080/upload
```

### 2. Process Template with Excel Data
**Endpoint:** `GET /process-template?filename=sample.xlsx`

**Response:**
```json
{
  "filled_templates": [
    "Hello John, your payment of $500 is due on 2025-03-20.",
    "Hello Alice, your payment of $300 is due on 2025-03-22."
  ]
}
```

## Running the Project
1. Start the application:
   ```sh
   go run main.go
   ```
2. Open your browser and test the endpoints using Postman or `curl`.

## Troubleshooting
- If `gin` or `excelize` is missing, run:
  ```sh
  go mod tidy
  ```
- If imports break, delete `go.mod` and `go.sum` and run:
  ```sh
  go mod init AutomateSystem
  go mod tidy
  ```

## License
This project is open-source and available under the MIT License.

