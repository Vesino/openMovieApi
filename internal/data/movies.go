package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

// Annotate the Movie struct with struct tags to control how the keys appear in the
// JSON-encoded output.
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // The - (hyphen) directive can be used when you never want a particular struct field to appear in the JSON output.
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`    // Add the omitempty directive
	Runtime   Runtime   `json:"runtime,omitempty"` // Add the omitempty directive
	Genres    []string  `json:"genres,omitempty"`  // Add the omitempty directive
	Version   int32     `json:"version"`
}

// Define a MovieModel struct type which wraps a sql.DB connection pool.
type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Insert(movie *Movie) error {
	query := `
		INSERT INTO movies (title, year, runtime, genres)
		VALUES  ($1, $2, $3, $4)
		RETURNING id, created_at, version
	`
	args := []interface{}{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}
	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func (m MovieModel) Get(id int64) (*Movie, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT id, created_at, title, year, runtime, genres, version
		FROM movies
		WHERE id = $1
	`
	var movie Movie
	err := m.DB.QueryRow(query, id).Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.Title,
		&movie.Year,
		&movie.Runtime,
		pq.Array(&movie.Genres),
		&movie.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &movie, nil
}

func (m MovieModel) Update(movie *Movie) error {
	return nil
}

func (m Movie) Delete(movie *Movie) error {
	return nil
}
