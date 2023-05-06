package service

import (
    "api-example/models"
    "api-example/store"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
)

type empService struct {
    s store.Store
}

func New(s store.Store) empService {
    return empService{s: s}
}

func (svc empService) Create(ctx *gin.Context, data models.Employee) (*models.Employee, error) {
    log.Debug("Service: starting create")
    res, err := svc.s.Create(ctx, data)
    if err != nil {
        return nil, err
    }

    log.Debug("Service: Successfully created employee")

    return res, nil
}

func (svc empService) Get(ctx *gin.Context) (*[]models.Employee, error) {
    log.Debug("Service: starting fetching")
    res, err := svc.s.Get(ctx)
    if err != nil {
        return nil, err
    }

    log.Debug("Service: Successfully fetched employee")

    return res, nil
}

func (svc empService) Update(ctx *gin.Context, id string, data models.Employee) (*models.Employee, error) {
    log.Debug("Service: starting update")
    res, err := svc.s.Update(ctx, id, data)
    if err != nil {
        return nil, err
    }

    log.Debug("Service: Successfully updated employee")

    return res, nil
}

func (svc empService) Delete(ctx *gin.Context, id string) (int, error) {
    log.Debug("Service: starting delete")
    res, err := svc.s.Delete(ctx, id)
    if err != nil {
        return res, err
    }

    log.Debug("Service: Successfully deleted employee")

    return res, nil
}
