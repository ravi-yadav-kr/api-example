package handler

import (
    "api-example/models"
    "api-example/service"
    "encoding/json"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
)

type handler struct {
    s service.Service
}

func New(s service.Service) handler {
    return handler{s: s}
}

func (h handler) Create(ctx *gin.Context) {
    var emp models.Employee

    err := ctx.Bind(&emp)
    if err != nil {
        log.Println(err)
        return
    }

    _, err = h.s.Create(ctx, emp)
    if err != nil {
        log.Println(err)
        return
    }

    data, err := json.Marshal(emp)

    ctx.Data(201, "", data)

    log.Println("handler: Successfully created employee")

    return

}

func (h handler) Get(ctx *gin.Context) {
    res, err := h.s.Get(ctx)
    if err != nil {
        ctx.Data(400, "", nil)
        return
    }

    data, err := json.Marshal(res)

    ctx.Data(200, "json", data)

    log.Println("handler: Successfully fetched data for employee")

    return
}

func (h handler) Update(ctx *gin.Context) {
    id := ctx.Query("id")

    var emp models.Employee

    ctx.Bind(&emp)

    res, err := h.s.Update(ctx, id, emp)
    if err != nil {
        ctx.Data(400, "", nil)
        return
    }

    data, err := json.Marshal(res)

    ctx.Data(200, "json", data)

    log.Println("handler: Successfully updated employee")

    return
}

func (h handler) Delete(ctx *gin.Context) {
    id := ctx.Query("id")

    res, err := h.s.Delete(ctx, id)
    if err != nil {
        ctx.Data(500, "", nil)
        return
    }

    if res == 0 {
        ctx.Data(204, "", nil)
        return
    }

    ctx.Data(200, "", nil)

    log.Println("handler: Successfully deleted employee")

    return
}
