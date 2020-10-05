package service

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/irvanherz/goblog/model"
)

// ArticleService struct
type ArticleService struct {
	connection *sql.DB
}

// NewArticleService func
func NewArticleService(connection *sql.DB) *ArticleService {
	return &ArticleService{connection: connection}
}

// Read articles
func (s *ArticleService) Read(filters model.ArticleQuery) (*[]model.Article, *model.PageInfo, error) {
	resultArticles := make([]model.Article, 0)
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
				p.id,
				p.title, p.content, 
				p.summary, 
				p.images, 
				p.tags, 
				p.created_at, 
				p.updated_at,
				u.id author_id,
				u.name author_name,
				u.gender author_gender,
				u.dob author_dob,
				u.photo author_photo
			FROM article p
			LEFT JOIN user u ON p.author_id=u.id
		) x 
	` + whereClause + `
	` + orderbyClause + `
	` + limitClause
	articleRows, articleErr := s.connection.Query(query, whereVals...)

	if articleErr != nil {
		return nil, nil, model.NewRequestError("1234", "Error reading database")
	}
	defer articleRows.Close()

	for articleRows.Next() {
		entry := model.Article{}
		if err := articleRows.Scan(
			&entry.ID,
			&entry.Title,
			&entry.Content,
			&entry.Summary,
			&entry.Images,
			&entry.Tags,
			&entry.CreatedAt,
			&entry.UpdatedAt,
			&entry.Author.ID,
			&entry.Author.Name,
			&entry.Author.Gender,
			&entry.Author.Dob,
			&entry.Author.Photo,
		); err != nil {
			return nil, nil, model.NewRequestError("12345", "Error reading database")
		}
		resultArticles = append(resultArticles, entry)
	}

	query = `
		SELECT COUNT(*) AS count 
		FROM (
			SELECT 
				p.id,
				p.title, p.content, 
				p.summary, 
				p.images, 
				p.tags, 
				p.created_at, 
				p.updated_at,
				u.id author_id,
				u.name author_name,
				u.gender author_gender,
				u.dob author_dob,
				u.photo author_photo
			FROM article p
			LEFT JOIN user u ON p.author_id=u.id
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

	return &resultArticles, &resultPageInfo, nil
}

// ReadByID func
func (s *ArticleService) ReadByID(ID int64) (*model.Article, error) {
	query := `
		SELECT
			p.id,
			p.title, p.content, 
			p.summary, 
			p.images, 
			p.tags, 
			p.created_at, 
			p.updated_at,
			u.id author_id,
			u.name author_name,
			u.gender author_gender,
			u.dob author_dob,
			u.photo author_photo
		FROM article p
		LEFT JOIN user u ON p.author_id=u.id
		WHERE p.id=?
	`
	row := s.connection.QueryRow(query, ID)
	if row == nil {
		return nil, model.NewRequestError("1234", "Error reading database")
	}

	e := model.Article{}
	if err := row.Scan(
		&e.ID,
		&e.Title,
		&e.Content,
		&e.Summary,
		&e.Images,
		&e.Tags,
		&e.CreatedAt,
		&e.UpdatedAt,
		&e.Author.ID,
		&e.Author.Name,
		&e.Author.Gender,
		&e.Author.Dob,
		&e.Author.Photo,
	); err != nil {
		return nil, model.NewRequestError("12345", "Error reading database")
	}
	return &e, nil
}

// DeleteByID func
func (s *ArticleService) DeleteByID(ID int64) error {
	query := `
		DELETE FROM article
		WHERE id=?
	`
	_, err := s.connection.Exec(query, ID)
	if err != nil {
		return model.NewRequestError("1234", "Error reading database")
	}
	return nil
}

// UpdateByID func
func (s *ArticleService) UpdateByID(ID int64, data *model.ArticleMutation) error {
	setKeys := make([]string, 0)
	setVals := make([]interface{}, 0)

	if data.Title != nil {
		setKeys = append(setKeys, "title=?")
		setVals = append(setVals, &data.Title)
	}
	if data.Content != nil {
		setKeys = append(setKeys, "content=?")
		setVals = append(setVals, &data.Content)
	}
	if data.Summary != nil {
		setKeys = append(setKeys, "summary=?")
		setVals = append(setVals, &data.Summary)
	}
	if data.Tags != nil {
		setKeys = append(setKeys, "tags=?")
		setVals = append(setVals, &data.Tags)
	}

	setClause := " "

	if len(setKeys) != 0 {
		setClause = " SET " + strings.Join(setKeys, ",")
	}

	query := `
		UPDATE article
	` + setClause + `
		WHERE id=?
	`
	// print(query)
	_, err := s.connection.Exec(query, append(setVals, ID)...)
	if err != nil {
		return model.NewRequestError("1234", "Error updating database")
	}
	return nil
}

// Create func
func (s *ArticleService) Create(data *model.ArticleMutation) error {
	setKeys := make([]string, 0)
	setVals := make([]interface{}, 0)

	if data.AuthorID != nil {
		setKeys = append(setKeys, "author_id=?")
		setVals = append(setVals, &data.AuthorID)
	}

	if data.Title != nil {
		setKeys = append(setKeys, "title=?")
		setVals = append(setVals, &data.Title)
	}
	if data.Content != nil {
		setKeys = append(setKeys, "content=?")
		setVals = append(setVals, &data.Content)
	}
	if data.Summary != nil {
		setKeys = append(setKeys, "summary=?")
		setVals = append(setVals, &data.Summary)
	}
	if data.Tags != nil {
		setKeys = append(setKeys, "tags=?")
		setVals = append(setVals, &data.Tags)
	}

	setClause := " "

	if len(setKeys) != 0 {
		setClause = " SET " + strings.Join(setKeys, ",")
	}

	query := `
		INSERT INTO article
	` + setClause

	// print(query)
	res, err := s.connection.Exec(query, setVals...)
	if err != nil {
		print(err.Error())
		return model.NewRequestError("1234", "Error mutating item in database")
	}
	insertID, err := res.LastInsertId()
	data.ID = &insertID
	return nil
}
