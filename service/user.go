package service

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/irvanherz/goblog/model"
)

// UserService struct
type UserService struct {
	connection *sql.DB
}

// NewUserService func
func NewUserService(connection *sql.DB) *UserService {
	return &UserService{connection: connection}
}

// Read users
func (s *UserService) Read(filters model.UserQuery) (*[]model.User, *model.PageInfo, error) {
	resultUsers := make([]model.User, 0)
	resultPageInfo := model.PageInfo{}

	whereConds := make([]string, 0)
	whereVals := make([]interface{}, 0)
	currentPage := int64(1)
	itemsPerPage := int64(10)
	sortBy := "created_at"
	sortOrder := "DESC"

	if filters.Page != nil {
		currentPage = *filters.Page
	}
	if filters.ItemsPerPage != nil {
		itemsPerPage = *filters.ItemsPerPage
	}
	if filters.SortBy != nil {
		sortBy = *filters.SortBy
	}
	if filters.SortOrder != nil {
		sortOrder = *filters.SortOrder
	}

	whereClause := ""
	orderbyClause := "ORDER BY x." + sortBy + " " + sortOrder
	limitClause := "LIMIT " + strconv.FormatInt((currentPage-1)*itemsPerPage, 10) + "," + strconv.FormatInt(itemsPerPage, 10)
	if len(whereConds) != 0 {
		whereClause = "WHERE " + strings.Join(whereConds, " AND ")
	}
	query := `
		SELECT * 
		FROM (
			SELECT 
				u.id,
				u.email, 
				u.name, 
				u.gender, 
				u.dob, 
				u.password,
				u.created_at, 
				u.updated_at
			FROM user u
		) x 
	` + whereClause + `
	` + orderbyClause + `
	` + limitClause
	userRows, userErr := s.connection.Query(query, whereVals...)

	if userErr != nil {
		return nil, nil, model.NewRequestError("1234", "Error reading database")
	}
	defer userRows.Close()

	for userRows.Next() {
		entry := model.User{}
		if err := userRows.Scan(
			&entry.ID,
			&entry.Email,
			&entry.Name,
			&entry.Gender,
			&entry.Dob,
			&entry.Password,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		); err != nil {
			return nil, nil, model.NewRequestError("12345", "Error reading database")
		}
		resultUsers = append(resultUsers, entry)
	}

	query = `
		SELECT COUNT(*) AS count 
		FROM (
			SELECT 
				u.id,
				u.email, 
				u.name, 
				u.gender, 
				u.dob, 
				u.password,
				u.created_at, 
				u.updated_at
			FROM user u
		) x 
	` + whereClause
	countRow := s.connection.QueryRow(query, whereVals...)
	if countRow == nil {
		return nil, nil, model.NewRequestError("1234", "Error reading database")
	}

	totalItems := int64(0)

	if err := countRow.Scan(&totalItems); err != nil {
		return nil, nil, model.NewRequestError("12345", "Error reading database")
	}

	resultPageInfo.CurrentPage = currentPage
	resultPageInfo.TotalItems = totalItems
	resultPageInfo.ItemsPerPage = itemsPerPage
	resultPageInfo.TotalPages = totalItems / itemsPerPage
	if totalItems%itemsPerPage > 0 {
		resultPageInfo.TotalPages++
	}

	return &resultUsers, &resultPageInfo, nil
}

// ReadByEmail func
func (s *UserService) ReadByEmail(email string) (*model.User, error) {
	query := `
		SELECT
			u.id,
			u.email, 
			u.name, 
			u.gender, 
			u.dob, 
			u.password,
			u.created_at, 
			u.updated_at
		FROM user u
		WHERE u.email=?
	`
	row := s.connection.QueryRow(query, email)
	if row == nil {
		return nil, model.NewRequestError("1234", "Error reading database")
	}

	entry := model.User{}
	err := row.Scan(
		&entry.ID,
		&entry.Email,
		&entry.Name,
		&entry.Gender,
		&entry.Dob,
		&entry.Password,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, model.NewRequestError("12345", "Error reading database")
	}
	return &entry, nil
}

// ReadByID func
func (s *UserService) ReadByID(ID int64) (*model.User, error) {
	query := `
		SELECT
			u.id,
			u.email, 
			u.name, 
			u.gender, 
			u.dob, 
			u.password,
			u.created_at, 
			u.updated_at
		FROM user u
		WHERE u.id=?
	`
	row := s.connection.QueryRow(query, ID)
	if row == nil {
		return nil, model.NewRequestError("1234", "Error reading database")
	}

	entry := model.User{}
	if err := row.Scan(
		&entry.ID,
		&entry.Email,
		&entry.Name,
		&entry.Gender,
		&entry.Dob,
		&entry.Password,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	); err != nil {
		return nil, model.NewRequestError("12345", "Error reading database")
	}
	return &entry, nil
}
