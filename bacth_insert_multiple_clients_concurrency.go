package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
	"time"
)

// Function to create a multipart form for the file and JSON
func createMultipartRequest(filePath, listID, url, authToken string) (*http.Request, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file
	part, err := writer.CreateFormFile("file", "4lakh1.csv")
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, err
	}

	// Add JSON field
	jsonField := fmt.Sprintf(`{"list_id": "%s"}`, listID)
	if err := writer.WriteField("json", jsonField); err != nil {
		return nil, err
	}

	writer.Close()

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", authToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

// Function to send a single request
func sendRequest(wg *sync.WaitGroup, client *http.Client, filePath, listID, url, authToken string, results chan<- error) {
	defer wg.Done()

	req, err := createMultipartRequest(filePath, listID, url, authToken)
	if err != nil {
		results <- err
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		results <- err
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		results <- fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		return
	}

	results <- nil
}

func testConcurrency(concurrent int, filePath, listID, url, authToken string) (successCount, errorCount int, totalTime time.Duration) {
	client := &http.Client{}
	var wg sync.WaitGroup
	results := make(chan error, concurrent)

	startTime := time.Now()

	for i := 0; i < concurrent; i++ {
		wg.Add(1)
		go sendRequest(&wg, client, filePath, listID, url, authToken, results)
	}

	wg.Wait()
	close(results)

	totalTime = time.Since(startTime)

	for err := range results {
		if err != nil {
			errorCount++
		} else {
			successCount++
		}
	}

	return
}

func main() {
	const (
		filePath  = "path/to/file"
		listID    = "list_id"
		url       = "https://filesync.vinmail.io/v1/file-sync"
		authToken = "Auth Token"
	)

	concurrent := 10 // Starting number of concurrent requests
	maxConcurrent := 1000 // Maximum concurrency to test
	step := 10            // Increment step

	fmt.Println("Starting concurrency test...")

	for concurrent <= maxConcurrent {
		fmt.Printf("\nTesting with %d concurrent requests...\n", concurrent)
		successCount, errorCount, totalTime := testConcurrency(concurrent, filePath, listID, url, authToken)

		fmt.Printf("Concurrency: %d, Success: %d, Errors: %d, Total Time: %s\n", concurrent, successCount, errorCount, totalTime)

		if errorCount > 0 {
			fmt.Printf("Reached error threshold at %d concurrent requests. Stopping test.\n", concurrent)
			break
		}

		concurrent += step
	}

	fmt.Println("Test completed.")
}















package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
	"time"
)

// Function to create a multipart form for the file and JSON
func createMultipartRequest(filePath, listID, url, authToken string) (*http.Request, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file
	part, err := writer.CreateFormFile("file", "4lakh1.csv")
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, err
	}

	// Add JSON field
	jsonField := fmt.Sprintf(`{"list_id": "%s"}`, listID)
	if err := writer.WriteField("json", jsonField); err != nil {
		return nil, err
	}

	writer.Close()

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", authToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

// Function to send a single request
func sendRequest(wg *sync.WaitGroup, client *http.Client, filePath, listID, url, authToken string, results chan<- error) {
	defer wg.Done()

	req, err := createMultipartRequest(filePath, listID, url, authToken)
	if err != nil {
		results <- err
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		results <- err
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		results <- fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		return
	}

	results <- nil
}

func testConcurrency(concurrent int, filePath, listID, url, authToken string) (successCount, errorCount int, totalTime time.Duration) {
	client := &http.Client{}
	var wg sync.WaitGroup
	results := make(chan error, concurrent)

	startTime := time.Now()

	for i := 0; i < concurrent; i++ {
		wg.Add(1)
		go sendRequest(&wg, client, filePath, listID, url, authToken, results)
	}

	wg.Wait()
	close(results)

	totalTime = time.Since(startTime)

	for err := range results {
		if err != nil {
			errorCount++
		} else {
			successCount++
		}
	}

	return
}

func main() {
	const (
		filePath  = "path/to/file"
		listID    = "list_id"
		url       = "https://filesync.vinmail.io/v1/file-sync"
		authToken = "Auth Token"
	)

	concurrent := 10 // Starting number of concurrent requests
	maxConcurrent := 1000 // Maximum concurrency to test
	step := 10            // Increment step

	fmt.Println("Starting concurrency test...")

	for concurrent <= maxConcurrent {
		fmt.Printf("\nTesting with %d concurrent requests...\n", concurrent)
		successCount, errorCount, totalTime := testConcurrency(concurrent, filePath, listID, url, authToken)

		fmt.Printf("Concurrency: %d, Success: %d, Errors: %d, Total Time: %s\n", concurrent, successCount, errorCount, totalTime)

		if errorCount > 0 {
			fmt.Printf("Reached error threshold at %d concurrent requests. Stopping test.\n", concurrent)
			break
		}

		concurrent += step
	}

	fmt.Println("Test completed.")
}















