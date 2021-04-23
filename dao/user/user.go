package user

import (
	"database/sql"
	"fmt"
	"github.com/ethan510010/week02_hw/customerror"
	"github.com/ethan510010/week02_hw/db"
	"github.com/ethan510010/week02_hw/models"
	"github.com/pkg/errors"
)

func GetUserBy(ID int) (*models.User, error) {
	// assume no special character here
	query := fmt.Sprintf("SELECT * from users where id = %v", ID)
	user, err := db.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(sql.ErrNoRows, query)
		}
		return nil, errors.Wrap(customerror.ErrorOthers, query)
	}
	return user.(*models.User), nil
}


