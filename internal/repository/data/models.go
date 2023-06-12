package data

import "github.com/binsabit/url_shortener/internal/repository/data/models"

type User struct {
	ID       int
	Email    string
	password string
}

type URLModel interface {
	GetURLByID(id int) (*models.URL, error)
	GetAllURLOfUser(id int) ([]*models.URL, error)
	CreateURL(url string) (*models.URL, error)
	UpdateURLByID(id int) error
	DeleteURLByID(id int) error
}
