package repository

import (
	"carRestAPI/internal/models"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RequestCarCatalog interface {
	Create(input []models.Car) error
	Delete(regNum string) error
	Update(regNum string, input models.UpdateCar) error
	GetAll(params models.Params) ([]models.Car, error)
}

type CarCatalogPostgres struct {
	db *sqlx.DB
}

func NewReqPostgres(db *sqlx.DB) *CarCatalogPostgres {
	return &CarCatalogPostgres{db: db}
}

func (cp *CarCatalogPostgres) Create(input []models.Car) error {
	tx, err := cp.db.Begin()
	if err != nil {
		return err
	}
	for _, elem := range input {

		var id int
		createPeopleQuery := fmt.Sprintf("INSERT INTO %s(name,surname,patronymic) VALUES ($1, $2, $3) RETURNING id", peopleTable)
		row := tx.QueryRow(createPeopleQuery, elem.Owner.Name, elem.Owner.Surname, elem.Owner.Patronymic)
		if err := row.Scan(&id); err != nil {
			tx.Rollback()
			return err
		}

		createCarQuery := fmt.Sprintf("INSERT INTO %s (regNum,mark,model,year) VALUES ($1 ,$2,$3,$4)", carTable)
		_, err = tx.Exec(createCarQuery, elem.RegNum, elem.Mark, elem.Model, elem.Year)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
func (cp *CarCatalogPostgres) Delete(regNum string) error {
	querty := fmt.Sprintf("DELETE FROM %s pl USING %s car WHERE pl.id =car.owner AND car.regNum=$1", peopleTable, carTable)
	_, err := cp.db.Exec(querty, regNum)

	return err
}

func (cp *CarCatalogPostgres) GetAll(params models.Params) ([]models.Car, error) {
	var cars []models.Car
	var args []interface{}
	argId := 1
	query := fmt.Sprintf(
		`SELECT regnum, mark, model, year, people.name, people.surname, people.patronymic 
					FROM %s car, %s people
					WHERE car.owner = people.id`, carTable, peopleTable)
	if params.Mark != "" {
		query += fmt.Sprintf(" and mark = $%d", argId)
		args = append(args, params.Mark)
		argId++
	}
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argId,argId+1)
	args = append(args, params.Limit, params.Offset)
	
	rows, err := cp.db.Query(query, args...)
	if err != nil {
		logrus.Error("Error query car in CarCatalogPostgres.GetAll")
		return nil, err
	}
	for rows.Next() {
		var car models.Car
		err = rows.Scan(&car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic)
		if err != nil {
			logrus.Error("Error scanning car")
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, err
}

func (cp *CarCatalogPostgres) Update(regNum string, input models.UpdateCar) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Mark != nil {
		setValues = append(setValues, fmt.Sprintf("mark=$%d", argId))
		args = append(args, *input.Mark)
		argId++
	}
	if input.Model != nil {
		setValues = append(setValues, fmt.Sprintf("model=$%d", argId))
		args = append(args, *input.Model)
		argId++
	}

	if input.Year != nil {
		setValues = append(setValues, fmt.Sprintf("year=$%d", argId))
		args = append(args, *input.Year)
		argId++
	}
	if input.Owner != nil {
		tx, err := cp.db.Begin()
		if err != nil {
			return err
		}
		setQuery := strings.Join(setValues, ", ")
		query := fmt.Sprintf("UPDATE %s car SET %s  WHERE car.regNum=$%d RETURNING owner", carTable, setQuery, argId)

		args = append(args, regNum)
		logrus.Debugf("updateQuery:%s", query)
		logrus.Debugf("args:%s", args)

		var owner string
		row := tx.QueryRow(query, args...)
		if err := row.Scan(&owner); err != nil {
			tx.Rollback()
			return err
		}
		setValues = make([]string, 0)
		args = make([]interface{}, 0)
		argId = 1
		if input.Owner.Name != nil {
			setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
			args = append(args, *input.Owner.Name)
			argId++
		}
		if input.Owner.Surname != nil {
			setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
			args = append(args, *input.Owner.Surname)
			argId++
		}
		if input.Owner.Patronymic != nil {
			setValues = append(setValues, fmt.Sprintf("patronymic=$%d", argId))
			args = append(args, *input.Owner.Patronymic)
			argId++
		}
		setQuery = strings.Join(setValues, ", ")
		query = fmt.Sprintf("UPDATE %s pl SET %s FROM %s car WHERE pl.id=car.owner and car.regNum=$%d", peopleTable, setQuery, carTable, argId)
		args = append(args, regNum)
		logrus.Debugf("updateQuery:%s", query)
		logrus.Debugf("args:%s", args)

		_, err = tx.Exec(query, args...)
		if err != nil {
			tx.Rollback()
			return err
		}
		return tx.Commit()

	} else {
		setQuery := strings.Join(setValues, ", ")
		query := fmt.Sprintf("UPDATE %s car SET %s  WHERE car.regNum=$%d", carTable, setQuery, argId)
		args = append(args, regNum)
		logrus.Debugf("updateQuery:%s", query)
		logrus.Debugf("args:%s", args)

		_, err := cp.db.Exec(query, args...)
		return err
	}
}
