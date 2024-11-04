package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
)

type AdministratorRepository interface {
	CreateAdministrator(firstName string, lastName string, authId int, phoneNumber string, email string, office string) (models.Administrator, error)
	GetAdministrators(queryParams models.AdministratorQueryParams) ([]models.Administrator, error)
	GetAdministrator(username string) (models.Administrator, error)
	UpdateAdministrator(username string, administratorUpdates models.AdministratorUpdate) (models.Administrator, error)
}

type PostgresAdministratorRepository struct {
	db *sql.DB
}

func CreatePostgresAdministratorRepository(db *sql.DB) *PostgresAdministratorRepository {
	return &PostgresAdministratorRepository{db: db}
}

func (r *PostgresAdministratorRepository) CreateAdministrator(firstName string, lastName string, authId int, phoneNumber string, email string, office string) (models.Administrator, error) {
	var administrator models.Administrator

	query := `SELECT * FROM public.create_administrator($1,$2,$3,$4,$5,$6)`

	err := r.db.QueryRow(query, firstName, lastName, authId, phoneNumber, email, office).Scan(&administrator.ID, &administrator.FirstName, &administrator.LastName, &administrator.AuthID, &administrator.PhoneNumber, &administrator.Email, &administrator.Office, &administrator.CreatedAt, &administrator.UpdatedAt)

	return administrator, err
}

func (r *PostgresAdministratorRepository) GetAdministrators(queryParams models.AdministratorQueryParams) ([]models.Administrator, error) {
	//administrator array
	var administrators = make([]models.Administrator, 0)

	query := `SELECT * FROM public.get_admins($1,$2,$3,$4,$5,$6,$7,$8)`

	rows, err := r.db.Query(query, queryParams.Limit, queryParams.Offset, queryParams.FirstName, queryParams.LastName, queryParams.Username, queryParams.Email, queryParams.PhoneNumber, queryParams.Office)
	if err != nil {
		return nil, err
	}
	//Be sure to close connection
	defer rows.Close()

	//Loop through returned rows
	for rows.Next() {
		administrator := models.Administrator{}
		if err := rows.Scan(&administrator.ID, &administrator.FirstName, &administrator.LastName, &administrator.AuthID, &administrator.PhoneNumber, &administrator.Email, &administrator.Office, &administrator.CreatedAt, &administrator.UpdatedAt); err != nil {
			return nil, err
		}
		administrators = append(administrators, administrator)
	}

	return administrators, nil
}

func (r *PostgresAdministratorRepository) GetAdministrator(username string) (models.Administrator, error) {
	//create Administrator
	var administrator models.Administrator

	//create query
	query := `SELECT * FROM public.get_administrator($1)`

	//execute function
	row := r.db.QueryRow(query, username)

	err := row.Scan(&administrator.ID, &administrator.FirstName, &administrator.LastName, &administrator.AuthID, &administrator.PhoneNumber, &administrator.Email, &administrator.Office, &administrator.CreatedAt, &administrator.UpdatedAt)

	return administrator, err
}

func (r *PostgresAdministratorRepository) UpdateAdministrator(username string, administratorUpdates models.AdministratorUpdate) (models.Administrator, error) {
	var administrator models.Administrator

	//create query
	query := `SELECT * FROM public.update_administrator($1,$2,$3,$4,$5,$6)`

	row := r.db.QueryRow(query, username, administratorUpdates.FirstName, administratorUpdates.LastName, administratorUpdates.PhoneNumber, administratorUpdates.Email, administratorUpdates.Office)

	err := row.Scan(&administrator.ID, &administrator.FirstName, &administrator.LastName, &administrator.AuthID, &administrator.PhoneNumber, &administrator.Email, &administrator.Office, &administrator.CreatedAt, &administrator.UpdatedAt)

	return administrator, err
}
