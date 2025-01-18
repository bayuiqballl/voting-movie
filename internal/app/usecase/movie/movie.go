package movie

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"vote-system/internal/entity"
	"vote-system/pkg/config"
	"vote-system/pkg/helper"

	"github.com/gofiber/fiber/v2"
)

func (ms *service) UploadMovie(c *fiber.Ctx, file *multipart.FileHeader) (resp entity.UploadMovieResponse, err error) {

	if file.Filename == "" {
		err = helper.Error(http.StatusBadRequest, "File cannot be empty", nil)
		return
	}

	// Check if the file is a video file
	if file.Header.Get("Content-Type") != "video/mp4" {
		err = helper.Error(http.StatusBadRequest, "Invalid file type", nil)
		return
	}

	// Check if the file size is within the allowed limit
	if file.Size > 1024*1024*10 { // 10 MB
		err = helper.Error(http.StatusBadRequest, "File size is too large", nil)
		return
	}

	// Create the uploads directory if it doesn't exist
	uploadDir := "../uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if mkdirErr := os.MkdirAll(uploadDir, os.ModePerm); mkdirErr != nil {
			err = helper.Error(http.StatusInternalServerError, "Failed to create uploads directory", mkdirErr)
			return entity.UploadMovieResponse{}, err
		}
	}
	fileName := file.Filename
	extension := filepath.Ext(fileName)
	newFileName := fmt.Sprintf("%s%s", helper.GenerateRandomString(10), extension) // Generate a random string and append it to the file name, extension)

	// Save the file to the uploads directory
	filePath := fmt.Sprintf("%s/%s", uploadDir, newFileName)
	if saveErr := c.SaveFile(file, filePath); saveErr != nil {
		err = helper.Error(http.StatusInternalServerError, "Failed to save file", saveErr)
		return entity.UploadMovieResponse{}, err
	}

	// Return the file's URL
	fileURL := fmt.Sprintf("%s/uploads/%s", config.SetupEnvFile().BaseURL, newFileName)

	resp = entity.UploadMovieResponse{
		WatchURL: fileURL,
	}

	return
}
