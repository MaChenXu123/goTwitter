package main

import (
	"goTwitter/router"
)

func main() {
	r := router.Router()
	r.Run(":8080")
}
