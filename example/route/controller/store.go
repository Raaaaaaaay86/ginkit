package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raaaaaaaay86/ginutil"
	"github.com/raaaaaaaay86/ginutil/example/route/controller/dto"
	"github.com/raaaaaaaay86/ginutil/example/route/middleware"
	"github.com/raaaaaaaay86/ginutil/example/route/service"
)

var _ ginutil.RouteGroup = (*Store)(nil)

type Store struct {
	stores service.Store
}

func (s *Store) GetRoutes() []ginutil.RouteFactory {
	return []ginutil.RouteFactory{
		s.Create,
		s.GetAll,
		s.IncrementTotalIncome,
	}
}

func (s *Store) v1(path string) ginutil.Path {
	return ginutil.Path{
		Name: fmt.Sprintf("/v1%s", path),
		Before: []gin.HandlerFunc{
			middleware.PrintMessage("authentication..."),
		},
	}
}

func (s *Store) Create() ginutil.Route {
	return ginutil.Route{
		Method: http.MethodPost,
		Path:   s.v1("/stores"),
		Handler: func(c *gin.Context) {
			id, err := s.stores.Create()
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return
			}

			c.JSON(http.StatusCreated, gin.H{"id": id})
		},
	}
}

func (s *Store) GetAll() ginutil.Route {
	return ginutil.Route{
		Method: http.MethodGet,
		Path:   s.v1("/stores"),
		Handler: func(c *gin.Context) {
			stores, err := s.stores.FindAll()
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return
			}

			c.JSON(http.StatusOK, dto.NewStoresFromEntities(stores))
		},
	}
}

func (s *Store) IncrementTotalIncome() ginutil.Route {
	return ginutil.Route{
		Method: http.MethodPut,
		Path:   s.v1("/stores/:id/total-income/inc"),
		Handler: func(c *gin.Context) {
			id, err := strconv.ParseInt(c.Param("id"), 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			income, err := s.stores.IncrementTotalIncome(id)
			if err != nil {
				if err.Error()[0:3] == "404" {
					c.Status(http.StatusNotFound)
					return
				}
				c.Status(http.StatusInternalServerError)
				return
			}

			c.JSON(http.StatusOK, gin.H{"id": id, "totalIncome": income})
		},
	}
}

func NewStore(
	stores service.Store,
) *Store {
	return &Store{
		stores: stores,
	}
}
