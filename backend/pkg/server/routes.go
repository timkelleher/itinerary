package backend

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/timkelleher/itinerary/pkg/db"
	"github.com/timkelleher/itinerary/pkg/models"
)

func healthRoute(c *gin.Context) {
	useItineraryMiddleware(c)
	c.JSON(http.StatusOK, "OK")
}

func statusesRoute(c *gin.Context) {
	useItineraryMiddleware(c)
	c.JSON(http.StatusOK, Statuses())
}

func methodsRoute(c *gin.Context) {
	useItineraryMiddleware(c)
	c.JSON(http.StatusOK, Methods())
}

func optionsRoute(c *gin.Context) {
	useItineraryMiddleware(c)
}

func collectionRoute(c *gin.Context) {
	useItineraryMiddleware(c)

	client, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	data, err := client.Collection(uint(id))
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, nilResponse{})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func collectionsRoute(c *gin.Context) {
	useItineraryMiddleware(c)

	client, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	data, err := client.Collections(db.CollectionsFilter{})
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, nilResponse{})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func createCollectionRoute(c *gin.Context) {
	useItineraryMiddleware(c)

	var request models.CollectionRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	client, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	var collection models.Collection
	collection.Name = request.Name
	collection.PathPrefix = request.PathPrefix
	collection.ContentType = request.ContentType
	err = client.InsertCollection(&collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, collection)
}

func updateCollectionRoute(c *gin.Context) {
	useItineraryMiddleware(c)

	var request models.CollectionRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	client, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	collection, err := client.Collection(uint(request.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	collection.Name = request.Name
	collection.PathPrefix = request.PathPrefix
	collection.ContentType = request.ContentType
	err = client.UpdateCollection(collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, collection)
}

func endpointRoute(c *gin.Context) {
	useItineraryMiddleware(c)

	client, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	data, err := client.Endpoint(uint(id))
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, nilResponse{})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func createEndpointRoute(c *gin.Context) {
	useItineraryMiddleware(c)

	var request models.EndpointRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	client, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	collection, err := client.Collection(uint(request.CollectionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	var endpoint models.Endpoint
	endpoint.Name = request.Name
	endpoint.Method = request.Method
	endpoint.ResponseStatusCode = request.ResponseStatusCode
	endpoint.Output = request.Output
	endpoint.Collection = collection
	endpoint.CollectionID = collection.ID
	err = client.InsertEndpoint(endpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, endpoint)
}

func updateEndpointRoute(c *gin.Context) {
	useItineraryMiddleware(c)

	var request models.EndpointRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	client, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	endpoint, err := client.Endpoint(uint(request.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	endpoint.Name = request.Name
	endpoint.Method = request.Method
	endpoint.Path = request.Path
	endpoint.ResponseStatusCode = request.ResponseStatusCode
	endpoint.Output = request.Output
	err = client.UpdateEndpoint(*endpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, endpoint)
}

func customEndpointRoute(c *gin.Context) {
	client, err := db.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	path := strings.ToLower(c.Request.URL.Path)
	method := strings.ToUpper(c.Request.Method)

	endpoint, err := client.EndpointFromPath(path, method)
	if err != nil && err == sql.ErrNoRows {
		c.String(http.StatusNotFound, "404 Not Found")
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	status, err := strconv.Atoi(endpoint.ResponseStatusCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	var output map[string]interface{}
	json.Unmarshal([]byte(endpoint.Output), &output)
	c.JSON(status, output)
}
