package main

import (
	"fmt"
	"net/http"
	"time"

	// "encoding/json"

	// "os"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/labstack/echo"

	//validator
	//GORM
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type (
	User struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	goRoutineExample()

	// scrapeEbay()

	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// // Migrate the schema
	// db.AutoMigrate(&Product{})

	// // Create
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// // Read
	// var product Product
	// db.First(&product, 1)                   // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// fmt.Println("Product = ", product.Code, " ", product.Price)

	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// fmt.Println("Product = ", product.Code, " ", product.Price)

	// // Delete - delete product
	// db.Delete(&product)
}

func scrapeEbay() {
	c := colly.NewCollector(
		colly.Async(true),                    // Turn on asynchronous requests
		colly.Debugger(&debug.LogDebugger{}), // Attach a debugger to the collector
	)

	var deals []string

	c.OnHTML("#refit-spf-container > div.sections-container > div.ebayui-dne-featured-card.ebayui-dne-featured-with-padding > div.ebayui-dne-item-featured-card.ebayui-dne-item-featured-card > div > div > div > div.dne-itemtile-detail", func(e *colly.HTMLElement) { //when you find these selectors, then we complete the completion founder
		// print("\n==========================================================")
		// fmt.Println("PARAGRAPHS", e.DOM.Find("p").Text())
		scrapeResult := e.DOM.Find("p").Text()
		// print("\nSCRAPE RESULT IS ", scrapeResult)
		// deals += scrapeResult
		deals = append(deals, scrapeResult)
		// print("----------------------------------------------------------\n")
	})

	fmt.Println("EYOOO", deals)

	for index, element := range deals {
		fmt.Println("Each deals are: ", index, " === ", element)
	}

	e := echo.New() //create a server, inir xollt
	e.GET("/scrape", func(ec echo.Context) (err error) {
		c.Visit("https://www.ebay.com/deals")
		c.Wait()

		firstDeal := deals[0]
		fmt.Println("First deal is =====", firstDeal)

		return ec.String(http.StatusOK, firstDeal)
	})

	// time.Sleep(2 * time.Second)
	e.Logger.Fatal(e.Start(":1323"))
	// select {} // this will cause the program to run forever
}

// func getUser(){
// 	e := echo.New()
// 	e.POST("/", func(c echo.Context) error {

// 		user := User{Name: "Kobe", Email: "kobe@gmail.com"}
//         userJSON, _ := json.Marshal(user)
// 		fmt.Println(string(userJSON))

// 		// m := echo.Map{
// 		// 	&User{
// 		// 		Name:  "Jon",
// 		// 		Email: "jon@labstack.com",
// 		// 	},
// 		// }
// 		if err := c.Bind(&user); err != nil {
// 			return err
// 		}
// 		return c.JSON(200, user)
// 	})
// 	e.Logger.Fatal(e.Start(":1323"))
// }

// // Handler
// func sendJSON(c echo.Context) error { //https://echo.labstack.com/guide/response
// 	u := &User{
// 	  Name:  "Jon",
// 	  Email: "jon@labstack.com",
// 	}
// 	return c.JSON(http.StatusOK, u)
//   }

// func postUser() {
// 	e := echo.New()
// 	e.Validator = &CustomValidator{validator: validator.New("validate", validator.BakedInValidators)}
// 	e.POST("/users", func(c echo.Context) (err error) {
// 		u := new(User)
// 		if err = c.Bind(u); err != nil {
// 			return
// 		}
// 		if err = c.Validate(u); err != nil {
// 			return
// 		}
// 		return c.JSON(http.StatusOK, u)
// 	})
// 	e.Logger.Fatal(e.Start(":1323"))
// }

// func helloWorld() {
// 	// var users = []User{ //a slice of User
// 	// 	"user": User{ Name: "Kobe", Email: "kobe@gmail.com" },
// 	// }
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		// return c.JSON(http.StatusOK, users)
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})
// 	e.Logger.Fatal(e.Start(":1323"))
// }

///////////-----------------------------------------
func goRoutineExample() {

	theMine := [5]string{"ore1", "ore2", "ore3"}
	oreChan := make(chan string)

	// Finder
	go func(mine [5]string) { //a function that takes a mine with 5 elements
		print("\n1")
		for _, item := range mine {
			print("\n2")
			oreChan <- item //send item to the channel
			print("\n3")
		}
	}(theMine) //(theMine) this executes this go function

	// Ore Breaker
	go func() {
		print("\n4")
		for i := 0; i < 3; i++ {
			print("\n5")
			foundOre := <-oreChan //receive
			print("\n6")
			fmt.Println("\nMiner: Received " + foundOre + " from finder")
			print("\n7")
		}
		print("\n8")
	}()
	print("\n9")
	<-time.After(time.Second * 5) // Again, ignore this for now
	print("\n10")

}
