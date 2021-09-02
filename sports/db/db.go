package db

import (
	"time"

	"syreclabs.com/go/faker"
)

func (r *sportsRepo) seed() error {

	statement, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS sports (id INTEGER PRIMARY KEY, meeting_id INTEGER, name TEXT, number INTEGER, visible INTEGER, home_team TEXT, away_team TEXT, advertised_start_time DATETIME)`)
	if err == nil {
		_, err = statement.Exec()
	}

	for i := 1; i <= 200; i++ {
		statement, err = r.db.Prepare(`INSERT OR IGNORE INTO sports (id, meeting_id, name, number, visible,home_team, away_team, advertised_start_time) VALUES (?,?,?,?,?,?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				faker.Number().Between(1, 10),
				faker.Team().Name(),
				faker.Number().Between(1, 12),
				faker.Number().Between(0, 1),
				faker.Team().Name(),
				faker.Team().Name(),
				faker.Time().Between(time.Now().UTC().AddDate(0, 0, -1), time.Now().UTC().AddDate(0, 0, 2)).Format(time.RFC3339),
			)
		}
	}

	return err
}
