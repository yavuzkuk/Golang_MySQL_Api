package controller

// func CreatePost(w http.ResponseWriter, r *http.Request) {

// 	var DbConnection = database.Db

// 	_, err := DbConnection.Prepare("INSERT INTO posts(title) VALUES(?)")
// 	if err != nil {
// 		panic(err)
// 	}

// 	body, err := io.ReadAll(r.Body)

// 	if err != nil {
// 		panic(err)
// 	}

// 	keyVal := make(map[string]string)
// 	json.Unmarshal(body, &keyVal)
// 	title := keyVal["title"]

// 	fmt.Println(title)

// }
