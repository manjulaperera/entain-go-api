package db

const (
	sportsList = "list"
	sportById  = "tuple"
)

func getSportQueries() map[string]string {
	return map[string]string{
		sportsList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible,
				home_team,
				away_team, 
				advertised_start_time,
				betting_closed_time
			FROM sports
		`,
		sportById: `
			SELECT
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				home_team,
				away_team,
				advertised_start_time,
				betting_closed_time
			FROM sports
			WHERE id= ?
		`,
	}
}
