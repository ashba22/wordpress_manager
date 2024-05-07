package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

// fetchData improved with better error handling and possibly a retry mechanism
func fetchData(url string) (interface{}, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

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
		db.Find(&websites)

		websitesRemap := make([]map[string]interface{}, len(websites))
		for i, website := range websites {
			websitesRemap[i] = map[string]interface{}{
				"id":          website.ID,
				"name":        website.Name,
				"url":         website.URL,
				"api_key":     website.Api_Key,
				"active":      website.Active,
				"description": website.Description,
			}
		}

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
		return c.JSON(website)
	})

	app.Get("/website/:id", func(c *fiber.Ctx) error {
		timeStart := time.Now()
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

		timeEnd := time.Since(timeStart)
		logger.Infof("Time taken for /website/%s: %v", id, timeEnd)
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
			Prefork:              true,
			DisableKeepalive:     true,
			ServerHeader:         "Fiber",
			StrictRouting:        true,
			CaseSensitive:        true,
			Immutable:            true,
			UnescapePath:         true,
			ETag:                 true,
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

	setupRoutes(app, db)
	logger.Fatal(app.Listen(":5001"))
}
