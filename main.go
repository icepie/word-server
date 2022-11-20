package main

import (
	"crypto/rand"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Word struct {
	Word string
	CN   string
}

var (
	wordList []Word
)

const (
	csvFilePath = "data.csv"
)

func main() {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(csvFile)

	rows, err := csvReader.ReadAll() // `rows` is of type [][]string
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		var word Word

		for i, col := range row {
			switch i {
			case 0:
				word.Word = col
			case 1:
				word.CN = col
			}
		}

		wordList = append(wordList, word)
	}

	// log.Println(wordList)

	// 设置随机种子
	// rand.Seed(time.Hour.Microseconds())

	// 随机抽取一个单词

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(wordList))))

		log.Println(wordList[i.Int64()])

		word := wordList[i.Int64()]

		c.String(http.StatusOK, fmt.Sprintf("%s\n%s", word.Word, word.CN))
	})

	r.Run()
}
