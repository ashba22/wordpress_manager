package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/earlydata"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Log setup
var logger = logrus.New()

func init() {
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Level = logrus.InfoLevel
}

type Website struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Name        string
	URL         string
	Api_Key     string
	Active      bool
	Description string
}

type APIResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type APIError struct {
	Error   string `json:"error"`
	Details string `json:"details"`
}

// Create a global Fasthttp client
var client = &fasthttp.Client{
	ReadTimeout:         10 * time.Second,
	WriteTimeout:        10 * time.Second,
	MaxIdleConnDuration: 30 * time.Second,
}

func fetchData(url string) (interface{}, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Connection", "close") // Add 'Connection: close' header

	retryCount := 3
	for i := 0; i < retryCount; i++ {
		if err := client.Do(req, resp); err != nil {
			log.Printf("Failed to fetch data from %s, error: %v", url, err)
			if i == retryCount-1 {
				return nil, fmt.Errorf("request failed after %d retries: %w", retryCount, err)
			}
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		log.Printf("Non-200 status from %s: %d", url, resp.StatusCode())
		body := resp.Body()
		log.Printf("Response body: %s", body)
		return nil, fmt.Errorf("non-200 response status: %d, body: %s", resp.StatusCode(), body)
	}

	var result interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		log.Printf("Failed to decode JSON from %s, error: %v", url, err)
		return nil, fmt.Errorf("JSON decoding failed: %w", err)
	}

	return result, nil
}

func setupRoutes(app *fiber.App, db *gorm.DB) {

	app.Get("/websites", func(c *fiber.Ctx) error {

		var websites []Website
		err := db.Find(&websites).Error
		if err != nil {
			return c.Status(500).SendString("Failed to fetch websites")
		}

		websitesRemap := make([]map[string]interface{}, len(websites)) // Initialize with the correct length

		// Use a WaitGroup to wait for all requests to finish
		var wg sync.WaitGroup

		for i, website := range websites {
			wg.Add(1)
			go func(i int, website Website) {
				defer wg.Done()

				// Send a HEAD request to the website to check if it is online
				resp, err := http.Head(website.URL)
				isOnline := false // Set the initial value to false

				if err == nil && resp.StatusCode == http.StatusOK {
					isOnline = true // Update the value to true if the request is successful and the status code is 200
				}

				// Create a map with the website details and online status
				websiteMap := map[string]interface{}{
					"id":          website.ID,
					"name":        website.Name,
					"url":         website.URL,
					"api_key":     website.Api_Key,
					"active":      website.Active,
					"description": website.Description,
					"is_online":   isOnline,
				}

				websitesRemap[i] = websiteMap
			}(i, website)
		}

		// Wait for all requests to finish
		wg.Wait()

		return c.JSON(websitesRemap)
	})

	app.Post("/websites", func(c *fiber.Ctx) error {
		website := new(Website)
		if err := c.BodyParser(website); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		if result := db.Create(&website); result.Error != nil {
			logger.Errorf("Failed to create website: %v", result.Error)
			return c.Status(500).JSON(APIError{"database_error", "Failed to create website"})
		}

		websiteMap := map[string]interface{}{
			"id":      website.ID,
			"name":    website.Name,
			"url":     website.URL,
			"api_key": website.Api_Key,
			"active":  website.Active,
		}

		return c.JSON(websiteMap)
	})
	app.Get("/websites/updates", func(c *fiber.Ctx) error {
		var websites []Website
		if err := db.Find(&websites).Error; err != nil {
			return c.Status(500).SendString("Failed to fetch websites")
		}

		websitesUpdatesInfo := make([]map[string]interface{}, len(websites))

		var wg sync.WaitGroup
		var mu sync.Mutex // Add a mutex to protect the shared slice
		for i, website := range websites {
			wg.Add(1)
			go func(i int, website Website) {
				defer wg.Done()

				client := &fasthttp.Client{}
				req := fasthttp.AcquireRequest()
				defer fasthttp.ReleaseRequest(req)
				req.SetRequestURI(fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/updates?api_key=%s", website.URL, website.Api_Key))

				resp := fasthttp.AcquireResponse()
				defer fasthttp.ReleaseResponse(resp)
				if err := client.Do(req, resp); err != nil {
					log.Printf("Failed to fetch updates from %s: %v", website.URL, err)
					// Handle the error and continue to the next website
					return
				}

				if resp.StatusCode() != fasthttp.StatusOK {
					log.Printf("Non-200 status from %s: %d", website.URL, resp.StatusCode())
					// Handle the error and continue to the next website
					return
				}

				var result interface{}
				if err := json.Unmarshal(resp.Body(), &result); err != nil {
					log.Printf("Failed to decode JSON from %s: %v", website.URL, err)
					// Handle the error and continue to the next website
					return
				}

				websiteMap := map[string]interface{}{
					"id":      website.ID,
					"name":    website.Name,
					"updates": result,
				}

				mu.Lock()
				websitesUpdatesInfo[i] = websiteMap
				mu.Unlock()
			}(i, website)
		}

		wg.Wait()

		return c.JSON(websitesUpdatesInfo)
	})

	app.Get("/website/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")
		var website Website
		if err := db.First(&website, "id = ?", id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("No website found with given ID")
		}

		responseObject := map[string]interface{}{
			"website": map[string]interface{}{
				"id":          website.ID,
				"name":        website.Name,
				"url":         website.URL,
				"api_key":     website.Api_Key,
				"active":      website.Active,
				"description": website.Description,
			},
		}

		urls := map[string]string{
			"plugins":  fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/installed-plugins?api_key=%s", website.URL, website.Api_Key),
			"themes":   fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/installed-themes?api_key=%s", website.URL, website.Api_Key),
			"info":     fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/website-info?api_key=%s", website.URL, website.Api_Key),
			"posts":    fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/posts?api_key=%s", website.URL, website.Api_Key),
			"pages":    fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/pages?api_key=%s", website.URL, website.Api_Key),
			"comments": fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/comments?api_key=%s", website.URL, website.Api_Key),
			"users":    fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/users?api_key=%s", website.URL, website.Api_Key),
			"media":    fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/media?api_key=%s", website.URL, website.Api_Key),
			"updates":  fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/updates?api_key=%s", website.URL, website.Api_Key),
		}

		ch := make(chan error, len(urls)) // Buffer the channel to avoid blocking
		for key, url := range urls {
			go func(key, url string) {
				data, err := fetchData(url)
				if err != nil {
					ch <- err // send error to channel
					return
				}
				responseObject["website"].(map[string]interface{})[key] = data
				ch <- nil
			}(key, url)
		}

		for range urls {
			if err := <-ch; err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
		}

		return c.JSON(responseObject)
	})

	app.Put("/website/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		updatedWebsite := new(Website)
		if err := c.BodyParser(updatedWebsite); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		var website Website
		db.First(&website, "id = ?", id)
		website.Name = updatedWebsite.Name
		website.URL = updatedWebsite.URL
		website.Active = updatedWebsite.Active
		website.Description = updatedWebsite.Description
		website.Api_Key = updatedWebsite.Api_Key
		db.Save(&website)
		return c.JSON(website)
	})

	app.Delete("/website/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		result := db.Delete(&Website{}, "id = ?", id)
		if result.Error != nil {
			return c.Status(500).SendString("Failed to delete website")
		}
		return c.SendString("Website successfully deleted")
	})

	app.Post("/website/:id/toggle-plugin", func(c *fiber.Ctx) error {
		id := c.Params("id")
		/// get plugin name and path from JSON body request
		var body struct {
			Name string `json:"name"`
			Path string `json:"path"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Failed to parse request body: %v", err))
		}

		pluginName := body.Name
		if pluginName == "" {
			return c.Status(fiber.StatusBadRequest).JSON(APIError{"bad_request", "Missing 'name' field in request body"})
		}

		pluginPath := body.Path
		if pluginPath == "" {
			return c.Status(fiber.StatusBadRequest).JSON(APIError{"bad_request", "Missing 'path' field in request body"})
		}

		var website Website
		if err := db.First(&website, "id = ?", id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("No website found with given ID")
		}

		// Send POST request to toggle the plugin
		urlWithPlugin := fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/toggle-plugin", website.URL)

		postData := map[string]string{
			"api_key": website.Api_Key,
			"name":    pluginName,
			"path":    pluginPath,
		}

		postBody, err := json.Marshal(postData)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to marshal request body: %v", err))
		}

		resp, err := http.Post(urlWithPlugin, "application/json", bytes.NewBuffer(postBody))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to toggle plugin: %v", err))
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to read response body: %v", err))
		}

		return c.JSON(APIResponse{
			Status:  "success",
			Data:    string(respBody),
			Message: "",
		})

	})

	app.Post("/website/:id/change-post-status", func(c *fiber.Ctx) error {
		id := c.Params("id")
		/// get post id and status from JSON body request
		var body struct {
			PostID int    `json:"id"`
			Status string `json:"status"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Failed to parse request body: %v", err))
		}

		postID := body.PostID
		if postID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(APIError{"bad_request", "Missing 'post_id' field in request body"})
		}

		status := body.Status
		if status == "" {
			return c.Status(fiber.StatusBadRequest).JSON(APIError{"bad_request", "Missing 'status' field in request body"})
		}

		var website Website
		if err := db.First(&website, "id = ?", id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("No website found with given ID")
		}

		// Send POST request to change post status
		urlWithPostStatus := fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/change-post-status", website.URL)

		println(urlWithPostStatus)
		println(postID)
		println(status)

		postData := map[string]interface{}{
			"api_key": website.Api_Key,
			"post_id": postID,
			"status":  status,
		}

		fmt.Println(postData)

		postBody, err := json.Marshal(postData)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to marshal request body: %v", err))
		}

		resp, err := http.Post(urlWithPostStatus, "application/json", bytes.NewBuffer(postBody))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to change post status: %v", err))
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to read response body: %v", err))
		}

		println(string(respBody))

		return c.JSON(APIResponse{
			Status:  "success",
			Data:    string(respBody),
			Message: "",
		})
	})

	app.Post("/website/:id/change-page-status", func(c *fiber.Ctx) error {
		id := c.Params("id")
		/// get page id and status from JSON body request
		var body struct {
			PageID int    `json:"id"`
			Status string `json:"status"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Failed to parse request body: %v", err))
		}

		pageID := body.PageID
		if pageID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(APIError{"bad_request", "Missing 'page_id' field in request body"})
		}

		status := body.Status
		if status == "" {
			return c.Status(fiber.StatusBadRequest).JSON(APIError{"bad_request", "Missing 'status' field in request body"})
		}

		var website Website
		if err := db.First(&website, "id = ?", id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("No website found with given ID")
		}

		// Send POST request to change page status
		urlWithPageStatus := fmt.Sprintf("%s/wp-json/wp-manager-plugin/v1/change-page-status", website.URL)

		println(urlWithPageStatus)

		postData := map[string]interface{}{
			"api_key": website.Api_Key,
			"page_id": pageID,
			"status":  status,
		}

		fmt.Println(postData)

		postBody, err := json.Marshal(postData)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to marshal request body: %v", err))
		}

		resp, err := http.Post(urlWithPageStatus, "application/json", bytes.NewBuffer(postBody))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to change page status: %v", err))
		}

		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to read response body: %v", err))
		}

		println(string(respBody))

		return c.JSON(APIResponse{
			Status:  "success",
			Data:    string(respBody),
			Message: "",
		})

	})

}

func main() {
	println("Starting server...")
	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				code := fiber.StatusInternalServerError
				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
				}
				return c.Status(code).JSON(fiber.Map{
					"error": err.Error(),
				})
			},
			Prefork:              false,
			ServerHeader:         "Fiber",
			StrictRouting:        true,
			CaseSensitive:        true,
			Immutable:            true,
			UnescapePath:         true,
			BodyLimit:            4 * 1024 * 1024,
			Concurrency:          256 * 1024,
			ReadTimeout:          10 * time.Second,
			WriteTimeout:         10 * time.Second,
			IdleTimeout:          120 * time.Second,
			CompressedFileSuffix: ".fiber.gz",
		},
	)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(earlydata.New())

	db, err := gorm.Open(sqlite.Open("websites.db"), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database")
	}

	if _, err := os.Stat("websites.db"); os.IsNotExist(err) {
		file, err := os.Create("websites.db")
		if err != nil {
			logger.Fatal("failed to create database file")
		}
		file.Close()
	}

	db.AutoMigrate(&Website{})

	var count int64
	db.Model(&Website{}).Count(&count)
	if count == 0 {
		websites := []Website{
			{Name: "WP1", URL: "https://wp1.goxcms.xyz/", Api_Key: "ewlg6v944t4qNrVb7R5iVijOgvETLEdqhIbheZwkJYs7RpeCyT7753cxStQLOPYp", Active: true, Description: "This is a WordPress website"},
			{Name: "WP2", URL: "https://wp2.goxcms.xyz/", Api_Key: "LaBslUiBe5M61DDgChntNaV2qSza4vUfEaoAAR6ASizrhRGRAf6a5p4RV6F3uTMc", Active: true, Description: "This is a WordPress website"},
		}
		db.Create(&websites)
	}

	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)
		logger.Printf("Request processed in %v\n", duration)
		return err
	})

	setupRoutes(app, db)

	/// server go app but add error handling for the server
	if err := app.Listen(":3000"); err != nil {
		logger.Fatal("failed to start server: ", err)
	}

}
