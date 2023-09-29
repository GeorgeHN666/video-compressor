package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
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

		// Create a pipe to stream the video data from FFmpeg to base64 encoding
		pipeReader, pipeWriter := io.Pipe()

		// Use FFmpeg to compress and encode the video directly "scale=640:480"
		cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-c:v", "libx264", "-crf", "23", "-f", "base64", "-")
		cmd.Stdin = file
		cmd.Stdout = pipeWriter
		cmd.Stderr = os.Stderr

		go func() {
			defer pipeWriter.Close()
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error running FFmpeg: %v\n", err)
			}
		}()

		// Read the base64-encoded video data from the pipe
		var dataBuffer bytes.Buffer
		_, err = io.Copy(&dataBuffer, pipeReader)
		if err != nil {
			return nil, err
		}

		// Append the encoded video to the videos slice
		Videos = append(Videos, &Video{
			URI:   dataBuffer.String(),
			Title: fileHeader.Filename,
		})
	}

	// At this point, 'videos' contains the encoded videos
	// You can do whatever you want with the 'videos' slice, such as sending it as JSON response.
	// For simplicity, we'll just print the results here.

	return Videos, nil
}
