package date_base

import (
	"database/sql"
	"fmt"
	"strings"

	"tg_bot/teacher"
)

// Подключение к базе данных
func DbConnect() (*sql.DB, error){
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable", host, port, user, password, dataBase)
	db, err := sql.Open("postgres", dbInfo)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Проверка подключения к базе данных
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, nil
}

// Функция для получения информации о преподавателе из базы данных
func GetTeacherInfo(db *sql.DB, fullName string) (*teacher.Teacher, error) {
	query := `
		SELECT t.teacher_id, t.name, t.birth_date, t.alma_mater, t.graduation_year, t.degree,
			   t.knowledge_score, t.teaching_skill_score, t.communication_score, t.leniency_score, t.overall_score,
			   t.knowledge_rating_num, t.teaching_skill_rating_num, t.communication_rating_num, t.leniency_rating_num, t.overall_rating_num,
			   array_agg(DISTINCT s.name) AS subjects, array_agg(DISTINCT d.name) AS departments
		FROM teachers t
		LEFT JOIN subject_teacher st ON t.teacher_id = st.teacher_id
		LEFT JOIN subjects s ON st.subject_id = s.subject_id
		LEFT JOIN department_teacher dt ON t.teacher_id = dt.teacher_id
		LEFT JOIN departments d ON dt.department_id = d.department_id
		WHERE t.name = $1
		GROUP BY t.teacher_id
	`

	row := db.QueryRow(query, fullName)

	var teacher teacher.Teacher
	var subjects, departments string

	err := row.Scan(
		&teacher.ID, &teacher.Name, &teacher.BirthDate, &teacher.AlmaMater, 
		&teacher.GraduationYear, &teacher.Degree,
		&teacher.KnowledgeScore, &teacher.TeachingSkillScore, 
		&teacher.CommunicationScore, &teacher.LeniencyScore, 
		&teacher.OverallScore, &teacher.KnowledgeRatingNum, 
		&teacher.TeachingSkillRatingNum, &teacher.CommunicationRatingNum, 
		&teacher.LeniencyRatingNum, &teacher.OverallRatingNum,
		&subjects, &departments,
	)

	if err != nil {
		return nil, err
	}

	// Обработка массива subjects
	if subjects != "{}" && subjects != "{NULL}" {
		teacher.Subjects = strings.Split(subjects[1:len(subjects)-1], ",")
	}
	// Обработка массива departments
	if departments != "{}" && departments != "{NULL}" {
		teacher.Departments = strings.Split(departments[1:len(departments)-1], ",")
	}


	return &teacher, nil
}

// Функция для получения списка преподавателей по первой букве
func GetTeachersByLetter(db *sql.DB, letter string) ([]teacher.Teacher, error) {
	query := `
		SELECT name 
		FROM teachers 
		WHERE name ILIKE $1 || '%'
		ORDER BY name
		LIMIT 100
	`

	rows, err := db.Query(query, letter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []teacher.Teacher
	for rows.Next() {
		var teacher teacher.Teacher
		if err := rows.Scan(&teacher.Name); err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	return teachers, nil
}