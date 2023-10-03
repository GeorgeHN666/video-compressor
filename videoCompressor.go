package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
)

func CompressVideos(r *http.Request) ([]*Video, error) {

	// Get the array of multipart files
	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		return nil, errors.New("no files attached")
	}

	// Create a slice to store the encoded videos
	var Videos []*Video

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		var TempFile bytes.Buffer

		encoder := base64.NewEncoder(base64.StdEncoding, &TempFile)

		// Encode the compressed video to base64
		// Copy the file content to the encoder
		_, err = io.Copy(encoder, file)
		if err != nil {
			return nil, err
		}

		// Close the encoder to flush any remaining data
		if err := encoder.Close(); err != nil {
			return nil, err
		}

		Videos = append(Videos, &Video{
			URI:   TempFile.String(),
			Title: fileHeader.Filename,
		})
	}

	// At this point, 'videos' contains the encoded videos
	// You can do whatever you want with the 'videos' slice, such as sending it as JSON response.
	// For simplicity, we'll just print the results here.

	return Videos, nil
}
