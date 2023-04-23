package websocket

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
)

const (
	WebSocketServerURL = "ws://ws_service_backend:1006/api"
)

var (
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
)

func ConnectToWebSocketServer(ctx context.Context) (*websocket.Conn, error) {
	apiURL := url.URL{Scheme: "ws", Host: "ws_service_backend:1006", Path: "/api"}
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.DialContext(ctx, apiURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket server: %w", err)
	}
	return conn, nil
}

func AuthenticateWebSocketConnection(conn *websocket.Conn) {
	token, err := GenerateJWTToken("my-secret-key")
	if err != nil {
		log.Fatal("JWT token generation error:", err)
	}
	authMsg := []byte(fmt.Sprintf(`{"type": "Authentaction", "token": "%s"}`, token))
	err = conn.WriteMessage(websocket.TextMessage, authMsg)
	if err != nil {
		log.Fatal("WebSocket write error:", err)
	}
	_, resp, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("Client 1: WebSocket read error:", err)
	}
	fmt.Println("Server response:", string(resp))
}

func SubscribeToChannel(conn *websocket.Conn, channel string) {
	subMsg := []byte(fmt.Sprintf(`{"type": "subscribe", "channel": "%s"}`, channel))
	err := conn.WriteMessage(websocket.TextMessage, subMsg)
	if err != nil {
		log.Fatal("WebSocket write error:", err)
	}
}

func WaitForWebSocketMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("Client 2: WebSocket read error:", err)
		}
		fmt.Println("Received message:", string(message))
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(c *fiber.Ctx) error {
	var message map[string]string

	c.BodyParser(&message)
	conn, err := ConnectToWebSocketServer(context.Background())
	if err != nil {
		log.Fatal("Client 3: WebSocket read error:", err)
	}
	defer conn.Close()
	SendMessageToUser(conn, message)

	return nil
}

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func SendMessageToUser(conn *websocket.Conn, message map[string]string) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	if err := conn.WriteMessage(websocket.TextMessage, messageBytes); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	_, resp, err := conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("failed to read server response: %w", err)
	}
	fmt.Println("Server response:", string(resp))

	// Close the websocket connection
	if err := conn.Close(); err != nil {
		return fmt.Errorf("failed to close websocket connection: %w", err)
	}

	return nil
}

func GenerateJWTToken(secretKey string) (string, error) {
	// Generate a JWT token using a library like "github.com/dgrijalva/jwt-go"
	// Example code:
	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["sub"] = "user123"
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// tokenString, err := token.SignedString([]byte(secretKey))
	// if err != nil {
	//     return "", err
	// }
	// return tokenString, nil

	// Note: The above code is just an example, you'll need to modify it to suit your specific use case
	return "", nil
}
