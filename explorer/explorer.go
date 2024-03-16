package explorer

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"log"
	"main/core/sendtx"
)

func Explorer() {
	app := fiber.New()
	app.Post("/chainservice/uploadchain", uploadchain)
	app.Post("/chainservice/updatechain", updatechain)
	log.Fatal(app.Listen(":3005"))
}

type OrgInfo struct {
	Number     string `json:"number"`
	Workamount string `json:"workamount"`
	Persion    string `json:"persion"`
	Workmethod string `json:"workmethod"`
	Worktime   string `json:"worktime"`
	Remarks    string `json:"remarks"`
}
type ErrorResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

func uploadchain(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println(string(c.Body()))
	payload := &OrgInfo{}
	if err := c.BodyParser(payload); err != nil {
		fmt.Println("Upload Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	hash := sendtx.AddOrUpdateUserData(payload.Number, payload.Workamount, payload.Persion, payload.Workmethod, payload.Worktime, payload.Remarks)
	fmt.Println(hash)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: hash})
}

func updatechain(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println(string(c.Body()))
	payload := &OrgInfo{}
	if err := c.BodyParser(payload); err != nil {
		fmt.Println("Updete Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	hash := sendtx.AddOrUpdateUserData(payload.Number, payload.Workamount, payload.Persion, payload.Workmethod, payload.Worktime, payload.Remarks)
	fmt.Println(hash)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: hash})
}
