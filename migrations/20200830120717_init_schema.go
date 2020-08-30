package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20200830120717",
		Up:      mig_20200830120717_init_schema_up,
		Down:    mig_20200830120717_init_schema_down,
	})
}

func mig_20200830120717_init_schema_up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE users ( name varchar(255) );")
	if err != nil {
		return err
	}
	return nil
}

func mig_20200830120717_init_schema_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users")
	if err != nil {
		return err
	}
	return nil
}
