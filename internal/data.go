package main

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	//"github.com/emirpasic/gods/lists/singlylinkedlist"
)

var userFile string = "/Users/Pedro Simoni/MyCalendarBot/data/users.csv"

func FindUser(currentUserID int64) (user user, err error) {
	file, err := os.Open(userFile)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	fileReader := csv.NewReader(file)

	header, err := fileReader.Read()
	if err != nil {
		return user, errors.New("error reading users file")
	}

	for {
		row, err := fileReader.Read()
		if err == nil {
			for i, h := range header {
				log.Printf("%v", h)
				if h == "UserID" {
					idRead, err := strconv.ParseInt(row[i], 10, 64)
					log.Printf("--%d--", idRead)
					if err != nil {
						return user, errors.New("error reading users file")
					} else if idRead != currentUserID {
						break
					} else {
						user.UserID = idRead
					}
				} else if h == " Notify" {
					notBoolRead, err := strconv.ParseBool(row[i])
					if err != nil {
						return user, errors.New("error reading users file")
					}

					user.Notify = notBoolRead
				} else if h == " DayLayout" {
					dayBoolRead, err := strconv.ParseBool(row[i])
					if err != nil {
						return user, errors.New("error reading users file")
					}

					user.DayLayout = dayBoolRead

					log.Printf("User found on data base")
					return user, nil
				}
			}
		} else if err == io.EOF {
			user.UserID = currentUserID
			user.Notify = false
			user.DayLayout = true

			StoreNewUser(user)

			log.Printf("New user stored in data base")
			return user, errors.New("user not found")
		} else {
			return user, errors.New("error reading users file")
		}
	}
}

func StoreNewUser(user user) (err error) {
	file, err := os.OpenFile(userFile, os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	fileWriter := csv.NewWriter(file)
	defer fileWriter.Flush()

	row := []string{strconv.FormatInt(user.UserID, 10), strconv.FormatBool(user.Notify), strconv.FormatBool(user.DayLayout)}
	err = fileWriter.Write(row)
	if err != nil {
		return err
	}

	return nil
}
