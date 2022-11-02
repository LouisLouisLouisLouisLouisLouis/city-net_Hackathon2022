package store

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebsiteType string

const (
	WebsiteTypeRedirect = "REDIRECT"
	WebsiteTypePdf      = "PDF"
	WebsiteTypeHtml     = "HTML"
)

type Store struct {
	Name        string      `bson:"name"`
	Subdomain   string      `bson:"subdomain"`
	WebsiteType WebsiteType `bson:"website_type"`
	Content     []byte      `bson:"content"`
	Redirect    string      `bson:"redirect"`
}

func (w *Store) Serve(c *gin.Context) {
	switch w.WebsiteType {
	case WebsiteTypeRedirect:
		c.Redirect(http.StatusMovedPermanently, w.Redirect)
	case WebsiteTypeHtml:
		c.Data(http.StatusOK, "text/html", w.Content)
	case WebsiteTypePdf:
		c.Data(http.StatusOK, "application/pdf", w.Content)
	}
}

type Repository interface {
	AddStore(Store Store) error
	GetStore(subdomain string) (Store, error)
	UpdateStore() error
	DeleteStore(subdomain string) error
}
