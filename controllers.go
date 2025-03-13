package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func UploadFile(c *gin.Context) {
	// Example function for handling file uploads
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

// Read and process the template with Excel data
func ProcessTemplate(c *gin.Context) {
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Filename parameter is required"})
		return
	}

	// Load Excel data
	data, err := ReadExcelData(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Read the template file
	templatePath := "templates/email_template.txt"
	templateContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read template file"})
		return
	}

	// Process each row of data and replace placeholders
	var responses []string
	for _, user := range data {
		filledTemplate := string(templateContent)
		filledTemplate = strings.ReplaceAll(filledTemplate, "{Name}", user.Name)
		filledTemplate = strings.ReplaceAll(filledTemplate, "{Date}", user.Date)
		filledTemplate = strings.ReplaceAll(filledTemplate, "{Amount}", user.Amount)
		filledTemplate = strings.ReplaceAll(filledTemplate, "{Status}", user.Status)

		responses = append(responses, filledTemplate)
	}

	c.JSON(http.StatusOK, gin.H{"filled_templates": responses})
}

// ReadExcelData extracts structured data from an Excel file
func ReadExcelData(filename string) ([]models.UserData, error) {
	filePath := filepath.Join("uploads", filename)
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open Excel file")
	}
	defer file.Close()

	rows, err := file.GetRows(file.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("failed to read Excel sheet")
	}

	var users []models.UserData
	for i := 1; i < len(rows); i++ { // Skip header row
		data := rows[i]
		if len(data) < 4 {
			continue
		}

		user := models.UserData{
			Name:   data[0],
			Date:   data[1],
			Amount: data[2],
			Status: data[3],
		}
		users = append(users, user)
	}

	return users, nil
}
