package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"os"
)

type Server struct {
	ServerAddr []string `json:"servers"`
}

func main() {

	//var serverLinks Server
	//json.Unmarshal([]byte(byteValue), &serverLinks)
	//fmt.Printf("%+v", serverLinks)

	app := fiber.New()
	app.Get("/api", func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("servers.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		//log.Println(c.Request())

		headers := c.GetReqHeaders()
		fmt.Println(headers["Deviceid"])

		return c.Send(byteValue)
	})

	app.Listen(":8080")

}
