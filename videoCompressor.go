package main

import (
	"bytes"
	"encoding/base64"
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

		// Create a buffer to store the base64-encoded video data
		var dataBuffer bytes.Buffer

		// Create an io.Pipe to redirect FFmpeg output to both dataBuffer and base64 encoder
		pipeReader, pipeWriter := io.Pipe()

		// Compress the video using FFmpeg
		cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-c:v", "libx264", "-crf", "23", "./output.mp4")
		cmd.Stdin = file
		cmd.Stdout = pipeWriter
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			return nil, err
		}
		fmt.Println("After exec")
		// Encode the FFmpeg output to base64
		defer pipeWriter.Close()
		encoder := base64.NewEncoder(base64.StdEncoding, &dataBuffer)
		fmt.Println("Got heree")
		_, err = io.Copy(encoder, pipeReader)
		if err != nil {
			return nil, err
		}
		fmt.Println("Got after")

		if err := encoder.Close(); err != nil {
			return nil, err
		}

		// Read the compressed video data from FFmpeg and store it in dataBuffer
		_, err = io.Copy(&dataBuffer, pipeReader)
		if err != nil {
			return nil, err
		}
		fmt.Println("Got to copy")

		// Wait for FFmpeg to finish
		if err := cmd.Wait(); err != nil {
			return nil, err
		}

		// // Encode the compressed video to base64
		// encodedData := base64.StdEncoding.EncodeToString(dataBuffer.Bytes())
		// // Append the encoded video to the videos slice
		// Videos = append(Videos, &Video{
		// 	URI:   encodedData,
		// 	Title: fileHeader.Filename,
		// })
	}

	// At this point, 'videos' contains the encoded videos
	// You can do whatever you want with the 'videos' slice, such as sending it as JSON response.
	// For simplicity, we'll just print the results here.

	return Videos, nil
}
