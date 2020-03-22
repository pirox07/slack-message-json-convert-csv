package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Ts       string `json:"ts"`
	Text     string `json:"text"`
	User     string `json:"user"`
	ThreadTs string `json:"thread_ts"`
}

type User struct {
	ID      string `json:"id"`
	Profile struct {
		RealName    string `json:"real_name"`
		DisplayName string `json:"display_name"`
	} `json:"profile"`
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

func getFileList(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	failOnError(err)

	var paths []string

	for _, file := range files {
		if !file.IsDir() {
			//if path.Ext(file.Name()) == ".json" {
			is_json, err := filepath.Match(`*.json`, file.Name())
			failOnError(err)
			if is_json == true {
				paths = append(paths, file.Name())
			} else {
				log.Println(file.Name(), "is not target.")
			}
		}
	}
	return paths
}

func userIDtoName(users []User, userID string) string {

	var userName string

	for _, v := range users {
		if userID == v.ID {
			userName = v.Profile.RealName + "/" + v.Profile.DisplayName
			break
		}
	}

	return userName
}

func main() {
	log.Println("----------- start -----------")

	// Create CSV file
	selfPath, _ := os.Executable()
	prevDir := filepath.Dir(selfPath)
	os.Chdir(prevDir)
	file, err := os.OpenFile("./SlackMessages.csv", os.O_WRONLY|os.O_CREATE, 0600)
	failOnError(err)

	// Initialize
	err = file.Truncate(0)
	failOnError(err)

	// Set Header
	writer := csv.NewWriter(file)
	writer.Write([]string{"thread_ts", "ts", "user_name", "text"})

	// Get FileName
	fmt.Println("Current dirctory is ", prevDir)
	srcFile := getFileList(prevDir)

	jsonString, err := ioutil.ReadFile("../users.json")
	failOnError(err)

	var users []User
	err = json.Unmarshal(jsonString, &users)
	failOnError(err)

	for i := 0; i < len(srcFile); i++ {
		log.Println("Open File: ", srcFile[i])

		jsonString, err := ioutil.ReadFile(srcFile[i])
		failOnError(err)

		var msg []Message
		err = json.Unmarshal(jsonString, &msg)
		failOnError(err)

		for l := 0; l < len(msg); l++ {

			// Convert user ID to user name
			userID := msg[l].User
			postUserName := userIDtoName(users, userID)

			// Convert mention user ID to user name
			tmpText := msg[l].Text
			re := regexp.MustCompile("<@" + `\w+` + ">")

			var mention []string
			mention = re.FindAllString(tmpText, -1)
			for n := 0; n < len(mention); n++ {
				userID = strings.TrimLeft(mention[n], "<@")
				userID = strings.TrimRight(userID, ">")
				mentionUserName := userIDtoName(users, userID)
				tmpText = strings.Replace(tmpText, mention[n], "<@"+mentionUserName+">", 1)
			}

			writer.Write([]string{
				msg[l].ThreadTs,
				unixTimeToJst(msg[l].Ts),
				postUserName,
				tmpText})
			writer.Flush()
		}
	}

	log.Println("------------ end ------------")
}
