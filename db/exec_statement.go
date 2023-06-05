package db

import (
	"database/sql"
	"fmt"
	"log"
)

func ExecQueryB(db *sql.DB, query string) ([]string, error) {
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		return []string{}, err
	}
	var res []string = make([]string, 0)
	for rows.Next() {
		var goals_scored, penalties, match_date, first_name, last_name string
		err := rows.Scan(&goals_scored, &penalties, &match_date, &first_name, &last_name)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return []string{}, err
		}
		data := fmt.Sprintf("GoalsScored: %s | WasPenalty: %s, | MatchDate: %s | FirstName: %s | LastName: %s",
			goals_scored, penalties, match_date, first_name, last_name)
		res = append(res, data)
	}
	return res, nil
}

func ExecQueryA(db *sql.DB, query string) ([]string, error) {
	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return []string{}, err
	}
	var res []string = make([]string, 0)
	for rows.Next() {
		var first_name, last_name string
		err := rows.Scan(&first_name, &last_name)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return res, err
		}
		data := fmt.Sprintf("FirstName: %s | LastName: %s", first_name, last_name)
		res = append(res, data)
	}
	return res, nil
}

func ExecQueryC(db *sql.DB, query string) ([]string, error) {
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		return []string{}, err
	}
	var res []string = make([]string, 0)
	for rows.Next() {
		var season_goals, season_penalties, season_cards, season_playtime, season_position string
		err := rows.Scan(&season_goals, &season_penalties, &season_cards, &season_playtime, &season_position)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return []string{}, err
		}
		data := fmt.Sprintf("Goals: %s | Penalties: %s | Cards: %s | Playtime: %s | Position: %s", season_goals, season_penalties, season_cards, season_playtime, season_position)
		res = append(res, data)
	}
	return res, nil
}

func ExecQueryD(db *sql.DB, query string) ([]string, error) {
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		return []string{}, err
	}
	var res []string = make([]string, 0)
	for rows.Next() {
		var season_matches, season_home_matches, season_away_matches, season_wins, season_losses, season_draws, season_home_wins, season_home_losses, season_home_draws, season_away_wins, season_away_losses, season_away_draws string
		err := rows.Scan(&season_matches, &season_home_matches, &season_away_matches, &season_wins, &season_losses, &season_draws, &season_home_wins, &season_home_losses, &season_home_draws, &season_away_wins, &season_away_losses, &season_away_draws)
		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return []string{}, err
		}
		data := fmt.Sprintf("season_matches: %s | season_home_matches: %s | season_away_matches: %s | season_wins: %s | season_losses: %s | season_draws: %s | season_home_wins: %s | season_home_losses: %s | season_home_draws: %s | season_away_wins: %s | season_away_losses: %s | season_away_draws: %s", season_matches, season_home_matches, season_away_matches, season_wins, season_losses, season_draws, season_home_wins, season_home_losses, season_home_draws, season_away_wins, season_away_losses, season_away_draws)
		res = append(res, data)
	}
	return res, nil
}
