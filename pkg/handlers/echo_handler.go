package handlers

import (
	"fmt"
	"go-dependency-injection/pkg/services"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
)

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
	log            *zap.Logger
	pyramidService *services.PyramidService
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(log *zap.Logger, pyramidService *services.PyramidService) *EchoHandler {
	return &EchoHandler{log: log, pyramidService: pyramidService}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Handling request", zap.String("path", r.URL.Path))
	if _, err := io.Copy(w, r.Body); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	}

	h.pyramidService.PrintPyramidWithLines(5)
}
