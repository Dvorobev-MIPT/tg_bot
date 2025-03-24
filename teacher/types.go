package teacher

import (
	"database/sql"
)

// Тип учитель - получаем из БД информация
type Teacher struct {
	ID                     int
	Name                   string
	BirthDate              sql.NullString
	AlmaMater              sql.NullString
	GraduationYear         sql.NullString
	Degree                 sql.NullString
	KnowledgeScore         sql.NullFloat64
	TeachingSkillScore     sql.NullFloat64
	CommunicationScore     sql.NullFloat64
	LeniencyScore          sql.NullFloat64
	OverallScore           sql.NullFloat64
	KnowledgeRatingNum     sql.NullInt64
	TeachingSkillRatingNum sql.NullInt64
	CommunicationRatingNum sql.NullInt64
	LeniencyRatingNum      sql.NullInt64
	OverallRatingNum       sql.NullInt64
	Subjects               []string
	Departments            []string
}
// потребуется для вычисление наиболее подходящих ФИО
type Similarity struct {
	name       string
	similarity float64
}

const (
	MaxSuggestions 	 = 3 // Максимальное количество предлагаемых вариантов
	SimilarityThreshold = 0.55 // Порог схожести (0-1)
)