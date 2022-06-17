package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practice/entity"
	"strconv"
	"strings"
)

var PORT = ":8080"

var userMap = map[int]entity.User{
	1: {
		ID:       1,
		Username: "Hallo1",
		Email:    "hallo1@mail.com",
		Password: "1234567",
		Age:      10,
	},
	2: {
		ID:       2,
		Username: "Hallo2",
		Email:    "hallo2@mail.com",
		Password: "1234567",
		Age:      10,
	},
	3: {
		ID:       3,
		Username: "Hallo3",
		Email:    "hallo3@mail.com",
		Password: "1234567",
		Age:      10,
	},
	4: {
		ID:       4,
		Username: "Hallo4",
		Email:    "hallo4@mail.com",
		Password: "1234567",
		Age:      10,
	},
	5: {
		ID:       5,
		Username: "Hallo5",
		Email:    "hallo5@mail.com",
		Password: "1234567",
		Age:      10,
	},
	6: {
		ID:       6,
		Username: "Hallo6",
		Email:    "hallo6@mail.com",
		Password: "1234567",
		Age:      10,
	},
}

func main() {
	http.HandleFunc("/user/", getUser)
	fmt.Println("Listen on port localhost" + PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path[1:]
	index := strings.Split(path, "/")

	var firstParam string

	if len(index) == 2 && index[1] != "" {
		firstParam = index[1]
	}

	switch r.Method {
	case "GET":
		if firstParam != "" {
			for i, v := range userMap {
				id := strconv.Itoa(i)
				if id == firstParam {
					getDataById, _ := json.Marshal(v)
					_, err := w.Write(getDataById)
					if err != nil {
						return
					}
				}
			}
		} else {
			var getAllResponse []entity.User
			for _, v := range userMap {
				getAllResponse = append(getAllResponse, v)
			}
			getDataAll, _ := json.Marshal(&getAllResponse)
			w.Write(getDataAll)
			return
		}
	case "POST":
		jsonDecoder := json.NewDecoder(r.Body)
		var newUser entity.User
		err := jsonDecoder.Decode(&newUser)

		if err != nil {
			return
		}

		if _, ok := userMap[int(newUser.ID)]; !ok {
			_, err := w.Write([]byte("user with given id already exist"))
			if err != nil {
				return
			}
		} else {
			userMap[int(newUser.ID)] = newUser
			createData, _ := json.Marshal(&userMap)
			_, err = w.Write(createData)
			if err != nil {
				return
			}
		}
	case "PUT":
		jsonDecoder := json.NewDecoder(r.Body)
		var newUser entity.User
		err := jsonDecoder.Decode(&newUser)

		if err != nil {
			return
		}

		userMap[int(newUser.ID)] = newUser
		createData, _ := json.Marshal(&userMap)
		_, err = w.Write(createData)
		if err != nil {
			return
		}

	case "DELETE":
		idx, _ := strconv.Atoi(firstParam)

		delete(userMap, idx)

		_, err := w.Write([]byte("Deleted successfully"))
		if err != nil {
			return
		}
	default:
		_, err := w.Write([]byte("Method not found"))
		if err != nil {
			return
		}
	}
}
