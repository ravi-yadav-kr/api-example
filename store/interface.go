package store

import (
	"api-example/models"
	"github.com/gin-gonic/gin"
)

type Store interface {
	Create(ctx *gin.Context, data models.Employee) (*models.Employee, error)
	Get(ctx *gin.Context) (*[]models.Employee, error)
	Update(ctx *gin.Context, id string, data models.Employee) (*models.Employee, error)
	Delete(ctx *gin.Context, id string) (int, error)
}
