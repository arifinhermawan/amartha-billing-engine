package pgsql

var (
	queryCreateUserInDB = `
		INSERT INTO "user"(name, "password", created_at, updated_at)
		VALUES(
			:name,
			:password,
			:created_at,
			:updated_at
		)
	`
)
