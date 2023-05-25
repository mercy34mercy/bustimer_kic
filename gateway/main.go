package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

// MongoDB接続情報
const (
	MongoDBURL     = "mongodb://root:example@mongodb:27017"
	DatabaseName   = "bustimer"
	CollectionName = "requests"
)

// RequestModelはMongoDBに保存するリクエストのデータモデルです
type RequestModel struct {
	Path        string    `bson:"path"`
	From        string    `bson:"from"`
	To          string    `bson:"to"`
	Timestamp   time.Time `bson:"timestamp"`
	IPAddress   string    `bson:"ip_address"`
	UserAgent   string    `bson:"user_agent"`
	DeviceType  string    `bson:"device_type"`
	AppPlatform string    `bson:"app_platform"`
	IsPWA       bool      `bson:"is_pwa"`
}

func main() {
	clientOptions := options.Client().ApplyURI(MongoDBURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()

	// リバースプロキシ先のURLを設定
	proxyURL, err := url.Parse("https://busdes-kic-prod-pxydxwabba-an.a.run.app")
	if err != nil {
		panic(err)
	}

	// リバースプロキシハンドラを作成
	proxy := httputil.NewSingleHostReverseProxy(proxyURL)

	// ルーティング
	router.Any("/*path", func(c *gin.Context) {
		// リバースプロキシヘッダーを設定
		c.Request.Host = proxyURL.Host
		c.Request.URL.Scheme = proxyURL.Scheme
		c.Request.URL.Host = proxyURL.Host

		// リクエストをリバースプロキシに転送
		proxy.ServeHTTP(c.Writer, c.Request)
		// 指定されたパスのみMongoDBにリクエストデータを保存
		if shouldSaveToMongoDB(c.Request.URL.Path) && !isLocalhostRequest(c.Request) {
			// Queryからフィールド値を取得
			from := c.Query("fr")
			to := c.Query("to")

			// ユーザのIPアドレスを取得
			ipAddress := c.ClientIP()

			// ユーザーエージェントからデバイスタイプとアプリプラットフォームを取得
			userAgent := c.Request.UserAgent()
			deviceType, appPlatform := parseUserAgent(userAgent)

			// PWAかどうかを判定
			isPWA := isRequestPWA(c.Request)

			// MongoDBにリクエストデータを保存
			collection := client.Database(DatabaseName).Collection(CollectionName)
			requestData := RequestModel{
				Path:        c.Request.URL.Path,
				From:        from,
				To:          to,
				Timestamp:   time.Now(),
				IPAddress:   ipAddress,
				UserAgent:   userAgent,
				DeviceType:  deviceType,
				AppPlatform: appPlatform,
				IsPWA:       isPWA,
			}
			_, err := collection.InsertOne(context.Background(), requestData)
			if err != nil {
				fmt.Println(err)
			}
		}
	})

	// サーバを起動
	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

// shouldSaveToMongoDBは指定されたパスが保存対象かどうかを判定します
func shouldSaveToMongoDB(path string) bool {
	allowedPaths := []string{"/nextbus", "/timetable", "/timetable/multi"}

	for _, allowedPath := range allowedPaths {
		if strings.HasPrefix(path, allowedPath) {
			return true
		}
	}
	return false
}

// parseUserAgentはユーザーエージェントを解析してデバイスタイプとアプリプラットフォームを返します
func parseUserAgent(userAgent string) (string, string) {
	deviceType := ""
	appPlatform := ""

	if strings.Contains(userAgent, "Android") {
		deviceType = "Android"
		appPlatform = "Native App"
	} else if strings.Contains(userAgent, "iOS") {
		deviceType = "iOS"
		appPlatform = "Native App"
	} else {
		deviceType = "Web"
		appPlatform = "Web"
	}

	return deviceType, appPlatform
}

// isRequestPWAはリクエストがPWAかどうかを判定します
func isRequestPWA(r *http.Request) bool {
	// ユーザーエージェントにPWAを示す文字列が含まれるかどうかを判定
	userAgent := r.UserAgent()
	if strings.Contains(userAgent, "wv") {
		return true
	}

	// リクエストヘッダーに"Service-Worker"が含まれるかどうかを判定
	if r.Header.Get("Service-Worker") != "" {
		return true
	}

	// Web App Manifestの存在を判定
	linkHeader := r.Header.Get("Link")
	if strings.Contains(linkHeader, "rel=\"manifest\"") {
		return true
	}

	return false
}

// isLocalhostRequestはリクエストがlocalhostからのアクセスかどうかを判定します
func isLocalhostRequest(r *http.Request) bool {
	remoteAddr := r.RemoteAddr
	if strings.HasPrefix(remoteAddr, "[::1]") || strings.HasPrefix(remoteAddr, "127.0.0.1") {
		return true
	}
	return false
}
