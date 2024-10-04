package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func CORSMiddleware() Middleware {
	return func(next LambdaHandler) LambdaHandler {
		return func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
			allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
			origin := req.Headers["origin"]

			log.Printf("Received request with Origin: %s", origin)
			log.Printf("Allowed Origins: %v", allowedOrigins)

			corsHeaders := map[string]string{
				"Access-Control-Allow-Headers":     "Content-Type,Authorization,X-Amz-Date,X-Api-Key,X-Amz-Security-Token",
				"Access-Control-Allow-Methods":     "GET,POST,PUT,DELETE,OPTIONS",
				"Access-Control-Allow-Credentials": "true",
			}

			originAllowed := false
			for _, allowedOrigin := range allowedOrigins {
				if origin == strings.TrimSpace(allowedOrigin) {
					corsHeaders["Access-Control-Allow-Origin"] = origin
					originAllowed = true
					log.Printf("Origin allowed: %s", origin)
					break
				}
			}

			if !originAllowed {
				log.Printf("Origin not allowed: %s", origin)
				// For development purposes, you might want to allow all origins
				// Remove this in production
				corsHeaders["Access-Control-Allow-Origin"] = "*"
				log.Printf("Temporarily allowing all origins for debugging")
			}

			// Handle preflight request
			if req.HTTPMethod == "OPTIONS" {
				log.Printf("Handling OPTIONS preflight request")
				return events.APIGatewayProxyResponse{
					Headers:    corsHeaders,
					StatusCode: 204,
				}, nil
			}

			// Call the next handler
			resp, err := next(req)

			// Add CORS headers to the response
			if resp.Headers == nil {
				resp.Headers = make(map[string]string)
			}
			for key, value := range corsHeaders {
				resp.Headers[key] = value
			}

			log.Printf("Response headers: %v", resp.Headers)

			return resp, err
		}
	}
}
