package store

import (
    "api-example/models"
    "database/sql"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    log "github.com/sirupsen/logrus"
)

type empStore struct {
    db *sql.DB
}

func New(db *sql.DB) empStore {
    return empStore{db: db}
}

func (s empStore) Create(ctx *gin.Context, data models.Employee) (*models.Employee, error) {
    log.Debug("Store: starting creating")

    id := uuid.New().String()

    log.Info(id)

    _, err := s.db.Exec(insertQuery, id, data.Name, data.Age, data.DeptName, data.PhoneNumber, data.Address, data.JoiningDate)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    log.Debug("Store: Successfully created employee")

    return &data, nil
}

func (s empStore) Get(ctx *gin.Context) (*[]models.Employee, error) {
    log.Debug("Store: starting fetching")

    rows, err := s.db.Query(getQuery)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    data := make([]models.Employee, 0)

    for rows.Next() {
        var temp models.Employee
        err = rows.Scan(&temp.ID, &temp.Name, &temp.Age, &temp.DeptName, &temp.PhoneNumber, &temp.Address, &temp.JoiningDate)
        if err != nil {
            continue
        }

        data = append(data, temp)
    }

    log.Debug("Store: Successfully fetched employee")

    return &data, nil
}

func (s empStore) Update(ctx *gin.Context, id string, data models.Employee) (*models.Employee, error) {
    log.Debug("Store: starting update")

    _, err := s.db.Exec(updateQuery, data.Name, data.Age, data.DeptName, data.PhoneNumber, data.Address, data.JoiningDate, id)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    log.Debug("Store: Successfully updated employee")
    return &data, nil
}

func (s empStore) Delete(ctx *gin.Context, id string) (int, error) {
    log.Debug("Store: starting delete")
    res, err := s.db.Exec(deleteQuery, id)
    if err != nil {
        return 0, err
    }

    code, err := res.RowsAffected()
    if err != nil {
        return 0, err
    }

    if code == 0 {
        log.Debug("Store: Employee Not found")
        return int(code), nil
    }

    log.Debug("Store: Successfully deleted employee")

    return int(code), nil
}
