package backend

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hack-city-net/city-net/backend/store"
	"net/http"
)

type LocalReversProxy struct {
	storeRepo store.Repository
	e         *gin.Engine
}

func NewLocalReverseProxy(repo store.Repository) *LocalReversProxy {
	e := gin.Default()

	return &LocalReversProxy{
		e:         e,
		storeRepo: repo,
	}
}

func (l *LocalReversProxy) proxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		server, err := l.storeRepo.GetStore(c.Request.Host)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
			c.Abort()
		}
		server.Serve(c)
	}
}

func (l *LocalReversProxy) Proxy(addr ...string) {
	l.e.GET("/*proxy", l.proxy())
	_ = l.e.Run(addr...)
}
