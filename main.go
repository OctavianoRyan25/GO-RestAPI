package main

import (
	"OctavianoRyan25/GO-JWT/database"
	"OctavianoRyan25/GO-JWT/router"
)

func main()  {
	database.StartDB()
	r:= router.StrartServer()
	r.Run(":8080")
}