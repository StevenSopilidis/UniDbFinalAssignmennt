package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	d "github.com/StevenSopilidis/DbFinalAssignment/db"
	e "github.com/StevenSopilidis/DbFinalAssignment/errors"
	"github.com/gin-gonic/gin"
)

var queries map[string]string
var database *sql.DB

type SuccesffullResponse struct {
	Data []string
}

func InitController(r *gin.Engine, db *sql.DB, q map[string]string) {
	// get queries
	queries = q
	// initialize connection to database
	database = db
	// setup routes
	r.GET("/api/players/goals_and_penalties/match/:id", getGoalsAndPenalties)
	r.GET("/api/team/:teamId/match/:matchId", getTeamCoachAtMatch)
	r.GET("/api/players/:playeId/stats/season/:startYear/:endYear", getPlayerStatsForSeason)
	r.GET("/api/teams/:teamId/stats/season/:startYear/:endYear", getTeamsStatsForSeason)
}

func getTeamsStatsForSeason(c *gin.Context) {
	startYear := c.Param("startYear")
	endYear := c.Param("endYear")
	startYearNum, err := strconv.Atoi(startYear)
	endYeaNum, err2 := strconv.Atoi(endYear)
	if err != nil || err2 != nil || startYearNum > endYeaNum {
		c.JSON(http.StatusBadRequest, e.GenerateInvalidArgumentError("Invalid years provided"))
		return
	}
	teamId := c.Param("teamId")
	_, err = strconv.Atoi(teamId)
	if err != nil {
		c.JSON(http.StatusBadRequest, e.GenerateInvalidArgumentError("Invalid playerId provided"))
		return
	}
	startYearDate := fmt.Sprintf("DATE '%s-9-1'", startYear)
	endYearDate := fmt.Sprintf("DATE '%s-5-28'", endYear)
	query := strings.ReplaceAll(queries["queryD"], "<teamId>", teamId)
	query = strings.ReplaceAll(query, "<season_start_placeholder>", startYearDate)
	query = strings.ReplaceAll(query, "<season_end_placeholder>", endYearDate)

	res, err := d.ExecQueryD(database, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, e.GenerateInternalError("Something went wrong"))
		return
	}
	c.JSON(http.StatusOK, SuccesffullResponse{
		Data: res,
	})
}

func getPlayerStatsForSeason(c *gin.Context) {
	startYear := c.Param("startYear")
	endYear := c.Param("endYear")
	startYearNum, err := strconv.Atoi(startYear)
	endYeaNum, err2 := strconv.Atoi(endYear)
	if err != nil || err2 != nil || startYearNum > endYeaNum {
		c.JSON(http.StatusBadRequest, e.GenerateInvalidArgumentError("Invalid years provided"))
		return
	}
	playerId := c.Param("playeId")
	_, err = strconv.Atoi(playerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, e.GenerateInvalidArgumentError("Invalid playerId provided"))
		return
	}

	// season start at first september and ends at 28 MAY
	startYearDate := fmt.Sprintf("DATE '%s-9-1'", startYear)
	endYearDate := fmt.Sprintf("DATE '%s-5-28'", endYear)
	query := strings.ReplaceAll(queries["queryC"], "<playerId>", playerId)
	query = strings.ReplaceAll(query, "<season_start_placeholder>", startYearDate)
	query = strings.ReplaceAll(query, "<season_end_placeholder>", endYearDate)

	res, err := d.ExecQueryC(database, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, e.GenerateInternalError("Something went wrong"))
		return
	}
	c.JSON(http.StatusOK, SuccesffullResponse{
		Data: res,
	})
}

func getTeamCoachAtMatch(c *gin.Context) {
	teamId := c.Param("teamId")
	matchId := c.Param("matchId")
	_, err := strconv.Atoi(teamId)
	_, err2 := strconv.Atoi(matchId)
	if err != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, e.GenerateInvalidArgumentError("Invalid matchId or teamId provided"))
		return
	}
	query := strings.ReplaceAll(queries["queryA"], "<teamId>", teamId)
	query = strings.ReplaceAll(query, "<matchId>", matchId)
	println(query)
	res, err := d.ExecQueryA(database, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, e.GenerateInternalError("Something went wrong"))
		return
	}
	c.JSON(http.StatusOK, SuccesffullResponse{
		Data: res,
	})
}

func getGoalsAndPenalties(c *gin.Context) {
	matchId := c.Param("id")
	_, err := strconv.Atoi(matchId)
	if err != nil {
		c.JSON(http.StatusBadRequest, e.GenerateInvalidArgumentError(err.Error()))
		return
	}
	query := strings.Replace(queries["queryB"], "<match_id>", matchId, 1)
	res, err := d.ExecQueryB(database, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, e.GenerateInternalError("Something went wrong"))
		return
	}
	c.JSON(http.StatusOK, SuccesffullResponse{
		Data: res,
	})
}
