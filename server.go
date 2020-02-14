package main

import (
	"net/http"
	"time"

	// "encoding/json"
	"fmt"
	// "os"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/labstack/echo"
	"gopkg.in/bluesuncorp/validator.v5" //validator
)

type (
	User struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func printString(s string) {
	fmt.Println(s)
}

func delaySecond(scrapeResult string, n time.Duration) {
	for _ = range time.Tick(n * time.Second) {
		str := "Hi! " + n.String() + " seconds have passed"
		printString(str)
	}
}

func main() {
	c := colly.NewCollector(
		colly.Async(true),                    // Turn on asynchronous requests
		colly.Debugger(&debug.LogDebugger{}), // Attach a debugger to the collector
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*", // when visiting links which domains' matches "*httpbin.*" glob
		Parallelism: 2,            // Limit the number of threads started by colly to two
		//Delay:      5 * time.Second,
	})

	e := echo.New() //create a server, inir xollt
	e.GET("/scrape", func(ec echo.Context) (err error) {
		var scrapeResult = ""
		//initialize scraper

		//SELECTORS LIST
		//Spotlight Deal = #refit-spf-container > div.sections-container > div.ebayui-dne-featured-card.ebayui-dne-featured-with-padding > div.ebayui-dne-summary-card.card.ebayui-dne-item-featured-card--topDeals.ebayui-dne-featured-with-carousel > div > div > div.dne-itemtile-detail
		//Trending Deals = #refit-spf-container > div.sections-container > div.ebayui-dne-featured-card.ebayui-dne-featured-with-padding > div.ebayui-dne-carousel.filmstrip-centered.ebayui-dne-carousel.ebayui-dne-trending-widget.filmstrip-1 > div
		//Trending Deals 2 (li:nth-child(9)) = #refit-spf-container > div.sections-container > div.ebayui-dne-featured-card.ebayui-dne-featured-with-padding > div.ebayui-dne-carousel.filmstrip-centered.ebayui-dne-carousel.ebayui-dne-trending-widget.filmstrip-1 > div > ul > li:nth-child(9) > div > div.dne-itemtile-detail
		//Featured Deals: #refit-spf-container > div.sections-container > div.ebayui-dne-featured-card.ebayui-dne-featured-with-padding > div.ebayui-dne-item-featured-card.ebayui-dne-item-featured-card > div
		//Featured Deals 2 (div:nth-child(1)): #refit-spf-container > div.sections-container > div.ebayui-dne-featured-card.ebayui-dne-featured-with-padding > div.ebayui-dne-item-featured-card.ebayui-dne-item-featured-card > div > div:nth-child(1) > div > div.dne-itemtile-detail

		// On every a element which has href attribute call callback
		c.OnHTML("#refit-spf-container > div.sections-container > div.ebayui-dne-featured-card.ebayui-dne-featured-with-padding > div.ebayui-dne-item-featured-card.ebayui-dne-item-featured-card > div > div > div > div.dne-itemtile-detail", func(e *colly.HTMLElement) {
			// link := e.Attr("href")
			// Print link
			print("\n==========================================================")
			print("\n++++++++E TEXT = ", e.Text, "\n\n")
			// print("\n++++++++E DOM = ", e.DOM.Text(), "\n\n")
			fmt.Printf("\n++++++++Link found: -> %s\n\n", e.Attr("href"))
			fmt.Println("PARAGRAPHS", e.DOM.Find("p").Text())
			scrapeResult = e.DOM.Find("p").Text()
			print("----------------------------------------------------------\n")
		})

		// Before making a request print "Visiting ..."
		// c.OnRequest(func(r *colly.Request) {
		// 	fmt.Println("Visiting", r.URL.String())
		// })

		// c.OnScraped(func(r *colly.Response) {
		// 	fmt.Println("Finished", r.Request.URL)
		// })

		// Start scraping on https://hackerspaces.org
		c.Visit("https://www.ebay.com/deals")
		// for i := 0; i < 1; i++ { // Start scraping in five threads on https://httpbin.org/delay/2
		// 	c.Visit(fmt.Sprintf("%s?n=%d", "https://www.ebay.com/deals", i))
		// }

		c.Wait() // Wait until threads are finished
		// time.Sleep(2 * time.Second)
		go delaySecond(scrapeResult, 5) // very useful for interval polling
		return ec.String(http.StatusOK, scrapeResult)
	})
	// time.Sleep(2 * time.Second)
	e.Logger.Fatal(e.Start(":1323"))
	select {} // this will cause the program to run forever
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
