<<<<<<< HEAD
package main

import (
	"apigolang/api"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_"github.com/go-sql-driver/mysql"
)

func main() {
	dns := "root:root@tcp(127.0.0.1:3306)/myfile?parseTime=true"
	
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("Error opening database: ", err)

	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established successfully!")

	api.RegisterRoutes(db)

	log.Println("Server started on :3030")
	log.Fatal(http.ListenAndServe(":3030",nil))
}
=======
///package main

///import "fmt"

///func main() {
///fmt.Println("Hello world")
///}

package main

import (
	"fmt"
	"time"

	"github.com/varlaalekya/goproject/pkg/greet"
	"github.com/varlaalekya/goproject/pkg/mathx"
)

func main() {
	const course = "Go Basics"
	student := "Alekya"

	fmt.Println(greet.Hello(student))

	numbers := []int{1, 2, 3, 4, 5, 10}
	sum := mathx.Sum(numbers...)
	fmt.Println("Sum:", sum)

	evens := 0
	for _, n := range numbers {
		if mathx.IsEven(n) {
			evens++
		}
	}

	switch {
	case evens == 0:
		fmt.Println("No even numbers.")
	case evens < len(numbers):
		fmt.Printf("%d even numbers found.\n", evens)
	default:
		fmt.Println("All are even.")
	}

	switch now := time.Now().Hour(); {
	case now < 12:
		fmt.Println("Good morning!")
	case now < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	fmt.Println("Max(7, 11):", mathx.Max(7, 11))
}

