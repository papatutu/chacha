package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func main() {
	// Define command-line flags
	portFlag := flag.String("port", "", "Port to run the server on")
	serveDirFlag := flag.String("serveDir", "", "Directory to serve static files from")
	flag.Parse()

	// Set up Viper to read configuration
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("env")    // type of the config file
	viper.AutomaticEnv()          // read in environment variables that match

	// Set default values
	viper.SetDefault("PORT", "9000")
	viper.SetDefault("SERVE_DIR", "./public")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	// Override with command-line flags if provided
	if *portFlag != "" {
		viper.Set("PORT", *portFlag)
	}
	if *serveDirFlag != "" {
		viper.Set("SERVE_DIR", *serveDirFlag)
	}

	// Get the port and serve directory from the config
	port := viper.GetString("PORT")
	serveDir := viper.GetString("SERVE_DIR")

	app := fiber.New()

	// Enable access logging
	app.Use(logger.New())

	// Serve static files from the configured directory
	app.Static("/", serveDir)

	// Catch all routes and serve index.html
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile(fmt.Sprintf("%s/index.html", serveDir))
	})

	// Start the server on the configured port
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
