package utils

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func CommitOrRollback(tx *sqlx.Tx, err *error) {
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			*err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
	} else {
		*err = tx.Commit()
	}
	return
}
