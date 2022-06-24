package backend

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type errorResponse struct {
	Error string `json:"error"`
}

type nilResponse struct {
}

func Run(port uint16) {
	router := gin.Default()

	router.GET("/health", healthRoute)

	// Backend routes
	router.GET("/itinerary/statuses", statusesRoute)
	router.GET("/itinerary/methods", methodsRoute)

	router.OPTIONS("/itinerary/collection", optionsRoute)
	router.PUT("/itinerary/collection", createCollectionRoute)
	router.GET("/itinerary/collection/:id", collectionRoute)
	router.GET("/itinerary/collections", collectionsRoute)
	router.OPTIONS("/itinerary/collection/:id", optionsRoute)
	router.POST("/itinerary/collection/:id", updateCollectionRoute)

	router.OPTIONS("/itinerary/endpoint", optionsRoute)
	router.PUT("/itinerary/endpoint", createEndpointRoute)
	router.GET("/itinerary/endpoint/:id", endpointRoute)
	router.OPTIONS("/itinerary/endpoint/:id", optionsRoute)
	router.POST("/itinerary/endpoint/:id", updateEndpointRoute)

	// Custom collection/endpoint routes
	router.Any("/api/*path", customEndpointRoute)

	// Frontend
	router.StaticFile("/", "./frontend/dist/index.html")
	router.StaticFS("/js/", http.Dir("./frontend/dist/js"))

	log.WithField("port", port).Info("Server Starting")
	log.Panic(
		router.Run(fmt.Sprintf(":%d", port)),
	)
}
