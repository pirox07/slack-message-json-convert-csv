package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Message struct {
	Ts          string `json:"ts"`
	Text        string `json:"text"`
	UserProfile struct {
		RealName string `json:"real_name"`
	} `json:"user_profile"`
	ThreadTs string `json:"thread_ts"`
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func unixTimeToJst(unixTimeStamp string) string {
	loc, err := time.LoadLocation("Asia/Tokyo")
	failOnError(err)

	f64, _ := strconv.ParseFloat(unixTimeStamp, 64)
	i64 := int64(f64)
	t := time.Unix(i64, 0).In(loc)

	return t.String()
}

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	failOnError(err)

	var paths []string
	for _, file := range files {
		extension := filepath.Ext(fmt.Sprint(file.Name()))
		if extension == ".json" {
			if file.IsDir() {
				paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
				continue
			}
			paths = append(paths, filepath.Join(dir, file.Name()))
		} else {
			log.Println(file.Name(), "is Non-Target.")
		}
	}

	return paths
}

func main() {
	log.Println("----------- start -----------")

	// Create CSV-File
	file, err := os.OpenFile("./SlackMessages.csv", os.O_WRONLY|os.O_CREATE, 0600)
	failOnError(err)

	// Initialize
	err = file.Truncate(0)
	failOnError(err)

	// Set Header
	writer := csv.NewWriter(file)
	writer.Write([]string{"thread_ts", "ts", "real_name", "text"})

	// Get FileName
	srcFile := dirwalk("./")

	for i := 0; i < len(srcFile); i++ {
		log.Println("Open File: ", srcFile[i])

		jsonString, err := ioutil.ReadFile(srcFile[i])
		failOnError(err)

		var msg []Message
		err = json.Unmarshal(jsonString, &msg)
		failOnError(err)

		for l := 0; l < len(msg); l++ {
			writer.Write([]string{
				msg[l].ThreadTs,
				unixTimeToJst(msg[l].Ts),
				msg[l].UserProfile.RealName,
				msg[l].Text})
			writer.Flush()
		}
	}

	log.Println("------------ end ------------")
}
