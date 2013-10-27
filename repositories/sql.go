package repositories

type rowScanner interface {
	Scan(dest ...interface{}) error
}

type rowCollector interface {
	collectRow(rs rowScanner) error
}

func exec(sql string, params ...interface{}) error {
	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(params...)

	return err
}

func queryRow(rc rowCollector, sql string, params ...interface{}) error {
	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(params...)

	return rc.collectRow(row)
}

func query(rc rowCollector, sql string, params ...interface{}) error {
	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(params...)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err := rc.collectRow(rows)

		if err != nil {
			return err
		}
	}

	return nil
}
