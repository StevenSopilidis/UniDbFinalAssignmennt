package main

import (
	d "github.com/StevenSopilidis/DbFinalAssignment/db"
	q "github.com/StevenSopilidis/DbFinalAssignment/queries"
	r "github.com/StevenSopilidis/DbFinalAssignment/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	addr := ":8080"
	db := d.InitDB()
	queries := q.GetQueries()
	r.InitController(engine, db, queries)
	engine.Run(addr)
}
