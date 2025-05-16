package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"diary/database"
	"diary/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

)

func main() {
	// データベースの初期化
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ginのルーター設定
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// ルートハンドラー（HTML表示）
	r.GET("/", func(c *gin.Context) {
		diaries, err := database.GetAllDiaries(db)
		if err != nil {
			c.HTML(500, "index.html", gin.H{"error": err.Error()})
			return
		}
		c.HTML(200, "index.html", gin.H{"diaries": diaries})
	})


	
	// APIエンドポイント
	// 日記投稿API
	r.POST("/api/diary", func(c *gin.Context) {
		content := c.PostForm("content")
		if content == "" {
			c.JSON(400, gin.H{"error": "content is required"})
			return
		}

		// 日本時間で現在時刻を取得
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		now := time.Now().In(jst)

		// ハッシュ値の生成
		hash := uuid.New().String()

		diary := models.Diary{
			Content:   content,
			Hash:      hash,
			CreatedAt: now,
		}

		err := database.CreateDiary(db, diary)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "日記を投稿しました",
			"hash":    hash,
		})
	})

	// 日記削除API
	r.DELETE("/api/diary/:hash", func(c *gin.Context) {
		hash := c.Param("hash")
		err := database.DeleteDiary(db, hash)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "日記を削除しました"})
	})

	// 日記一覧取得API
	r.GET("/api/diaries", func(c *gin.Context) {
		diaries, err := database.GetAllDiaries(db)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, diaries)
	})

	// Webサーバー起動
	r.Run(":8080")
}

func postDiary(db *database.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("日記を入力してください（終了するには空行を入力）:")

	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	content := strings.Join(lines, "\n")
	if content == "" {
		fmt.Println("日記が空です。")
		return
	}

	// 日本時間で現在時刻を取得
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst)

	// ハッシュ値の生成
	hash := uuid.New().String()

	diary := models.Diary{
		Content:   content,
		Hash:      hash,
		CreatedAt: now,
	}

	err := database.CreateDiary(db, diary)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("日記を投稿しました。ハッシュ値: %s\n", hash)
}
