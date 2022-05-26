package event

type Movie struct {
	ID       int64  `db:"id,omitempty"`
	Name     string `db:"name"`
	Director string `db:"director"`
	Year     int64  `db:"year"`
}
