package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// .envファイルを読み込んでいる。読み込めなかったらエラー。
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	// os.Getenv(キー名)で値を取得できる
	// .envはgitignoreさせることを忘れないように。
	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Println("s3Bucketenv:" + s3Bucket + "\n" + "secretKey:" + secretKey)

}
