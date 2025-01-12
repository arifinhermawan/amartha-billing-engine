package pgsql

const (
	queryCreateUserInDB = `
		INSERT INTO "user"(name, "password", created_at, updated_at)
		VALUES(
			:name,
			:password,
			:created_at,
			:updated_at
		)
	`
	queryGetUserByIDFromDB = `
		SELECT
			id,
			name
		FROM
			"user"
		WHERE
			id = $1
	`
)
