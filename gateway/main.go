package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
	Path      string    `bson:"path"`
	From      string    `bson:"from"`
	To        string    `bson:"to"`
	Timestamp time.Time `bson:"timestamp"`
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
		if shouldSaveToMongoDB(c.Request.URL.Path) {
			// Queryからフィールド値を取得
			from := c.Query("fr")
			to := c.Query("to")

			// MongoDBにリクエストデータを保存
			collection := client.Database(DatabaseName).Collection(CollectionName)
			requestData := RequestModel{
				Path:      c.Request.URL.Path,
				From:      from,
				To:        to,
				Timestamp: time.Now(),
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
