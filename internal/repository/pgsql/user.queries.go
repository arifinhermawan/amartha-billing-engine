package pgsql

const (
	queryGetDelinquentsUsers = `
		SELECT DISTINCT u.id AS id, u.name
		FROM "user" u
		JOIN loan l ON u.id = l.user_id
		JOIN payment p ON l.id = p.loan_id
		WHERE l.is_active = TRUE
		AND p.is_paid = FALSE
		AND p.due_date < $1
		AND (
			SELECT COUNT(*)
			FROM payment sub_p
			WHERE sub_p.loan_id = l.id
				AND sub_p.is_paid = FALSE
				AND sub_p.due_date < $1
		) >= 2;
	`

	queryCreateUser = `
		INSERT INTO "user"(name, "password", created_at, updated_at)
		VALUES(
			:name,
			:password,
			:created_at,
			:updated_at
		)
	`
	queryGetUserByID = `
		SELECT
			id,
			name
		FROM
			"user"
		WHERE
			id = $1
	`
)
