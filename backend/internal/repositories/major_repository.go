package repositories

import (
	"database/sql"
	"enroll-tracker/internal/models"
)

type MajorRepository interface {
	CreateMajor(majorCreation models.MajorCreation) (models.Major, error)
	GetMajors(queryParams models.MajorQueryParams) ([]models.Major, error)
	GetMajor(majorID int) (models.Major, error)
	UpdateMajor(majorID int, majorUpdates models.MajorUpdate) (bool, error)
	DeleteMajor(majorID int) (bool, error)
}

type PostgresMajorRepository struct {
	db *sql.DB
}

func CreateMajorRepository(db *sql.DB) *PostgresMajorRepository {
	return &PostgresMajorRepository{db: db}
}

func (r *PostgresMajorRepository) CreateMajor(majorCreation models.MajorCreation) (models.Major, error) {
	var major models.Major

	query := `SELECT * FROM public.create_major($1,$2)`

	err := r.db.QueryRow(query, majorCreation.Name, majorCreation.Description).Scan(&major.ID, &major.Name, &major.Description, &major.Status, &major.LastUpdated, &major.CreatedAt)
	if err != nil {
		return models.Major{}, err
	}
	return major, nil
}
func (r *PostgresMajorRepository) GetMajors(queryParams models.MajorQueryParams) ([]models.Major, error) {
	majors := make([]models.Major, 0)

	query := `SELECT * FROM public.get_majors($1,$2,$3,$4,$5)`

	rows, err := r.db.Query(query, queryParams.Limit, queryParams.Offset, queryParams.Name, queryParams.Description, queryParams.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var major models.Major
		if err := rows.Scan(&major.ID, &major.Name, &major.Description, &major.Status, &major.LastUpdated, &major.CreatedAt); err != nil {
			return majors, err
		}

		majors = append(majors, major)
	}
	return majors, nil
}

func (r *PostgresMajorRepository) GetMajor(majorID int) (models.Major, error) {
	var major models.Major

	query := `SELECT * FROM public.get_major($1)`

	err := r.db.QueryRow(query, majorID).Scan(&major.ID, &major.Name, &major.Description, &major.Status, &major.LastUpdated, &major.CreatedAt)
	if err != nil {
		return models.Major{}, err
	}
	return major, nil
}

func (r *PostgresMajorRepository) UpdateMajor(majorID int, majorUpdates models.MajorUpdate) (bool, error) {
	var success bool

	query := `SELECT * FROM public.update_major($1,$2,$3)`

	err := r.db.QueryRow(query, majorID, majorUpdates.Description, majorUpdates.Status).Scan(&success)
	if err != nil {
		return false, err
	}
	if !success {
		return false, models.NoAffectedRows
	}
	return true, nil
}

func (r *PostgresMajorRepository) DeleteMajor(majorID int) (bool, error) {
	var success bool

	query := `SELECT * FROM public.delete_major($1)`

	err := r.db.QueryRow(query, majorID).Scan(&success)
	if err != nil {
		return false, err
	}
	if !success {
		return false, nil
	}
	return true, nil
}
