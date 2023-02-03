package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// for mapping json request body
type Request struct {
	Init_password string // password before checking
}

// return total action needed to make string pwd become strong password
func password_validation(pwd string) int {
	k := 0 // total of add or remove char
	// check short or long password
	var long_pwd bool
	if len(pwd) < 6 {
		k = 6 - len(pwd)
		long_pwd = false
	}
	if len(pwd) >= 20 {
		k = len(pwd) - 19
		long_pwd = true
	}

	num_of_step := 0
	flag := [3]bool{false, false, false} // hasUpper, hasLower, hasDigit
	con := 1                             // consecutive char
	re := 0                              // for total add or replace char needed
	for i, char := range pwd {
		if unicode.IsUpper(char) {
			flag[0] = true
		}
		if unicode.IsLower(char) {
			flag[1] = true
		}
		if unicode.IsDigit(char) {
			flag[2] = true
		}
		// count consecutive char
		if i != 0 && pwd[i-1] == pwd[i] {
			con += 1
		} else {
			con = 1
		}
		// have 3 consecutive char
		if con == 3 {
			if k > 0 {
				if long_pwd { // use remove char action
					con = 2
					num_of_step += 1
				} else { // use add char action
					con = 1
					re += 1
				}
				k -= 1
			} else { // use replace action
				con = 0
				re += 1
			}
		}
	}
	// check uppercase, lowercase and digit
	for _, value := range flag {
		if !value {
			if re > 0 {
				re -= 1
			} else if k > 0 && !long_pwd {
				k -= 1
			}
			num_of_step += 1
		}
	}
	num_of_step += k + re
	return num_of_step
}

func main() {
	// load .env file
	godotenv.Load()

	// connect to database
	db, err := sql.Open("postgres", os.Getenv("CONNECTION_STRING"))
	if err != nil {
		log.Fatalln(err)
	}

	// initial gin server
	server := gin.Default()

	server.GET("/api/strong_password_steps", func(ctx *gin.Context) {

		// read json request body
		defer ctx.Request.Body.Close()
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err})
		}

		// initail Request struct for mapping
		var res = &Request{}
		err = json.Unmarshal(body, res)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err})
		}

		// check password
		password := res.Init_password
		result := password_validation(password)

		// add to database ( table log value (init_password, num_of_step) )
		query_string := fmt.Sprintf("INSERT INTO log VALUES ('%v', %v)", password, result)
		_, err = db.Exec(query_string)
		if err != nil {
			ctx.JSON(400, gin.H{"error": fmt.Sprintf("An error occured while executing query: %v", err)})
		}

		// respone the result
		ctx.JSON(200, gin.H{"num_of_step": result})
	})

	server.Run()
}
