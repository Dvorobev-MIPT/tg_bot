package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	levenshtein "github.com/texttheater/golang-levenshtein/levenshtein"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "yourusername"
	password = "yourpassword"
	dbname   = "yourdbname"
	maxSuggestions = 3 // Максимальное количество предлагаемых вариантов
	similarityThreshold = 0.7 // Порог схожести (0-1)
)

type Teacher struct {
	ID                   int
	Name                 string
	BirthDate            sql.NullString
	AlmaMater            sql.NullString
	GraduationYear       sql.NullString
	Degree               sql.NullString
	KnowledgeScore       sql.NullFloat64
	TeachingSkillScore   sql.NullFloat64
	CommunicationScore   sql.NullFloat64
	LeniencyScore        sql.NullFloat64
	OverallScore         sql.NullFloat64
	KnowledgeRatingNum   sql.NullInt64
	TeachingSkillRatingNum sql.NullInt64
	CommunicationRatingNum sql.NullInt64
	LeniencyRatingNum    sql.NullInt64
	OverallRatingNum     sql.NullInt64
	Subjects             []string
	Departments          []string
}

func main() {
	// Подключение к базе данных
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка подключения к базе данных
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI("YOUR_TELEGRAM_BOT_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Обработка входящих сообщений
	for update := range updates {
		if update.Message != nil { // Если есть сообщение
			// Получаем ФИО преподавателя
			fullName := strings.TrimSpace(update.Message.Text)

			// Ищем преподавателя в базе данных
			teacher, err := getTeacherInfo(db, fullName)
			if err != nil {
				// Если преподаватель не найден, ищем похожие имена
				suggestions, err := findSimilarTeachers(db, fullName)
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка при поиске преподавателя.")
					bot.Send(msg)
					continue
				}

				if len(suggestions) > 0 {
					message := "Преподаватель не найден. Возможно, вы имели в виду:\n"
					for i, name := range suggestions {
						message += fmt.Sprintf("%d. %s\n", i+1, name)
					}
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
					bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Преподаватель не найден.")
					bot.Send(msg)
				}
				continue
			}

			// Формируем ответ
			response := formatTeacherInfo(teacher)

			// Отправляем ответ пользователю
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			bot.Send(msg)
		}
	}
}

// Функция для получения информации о преподавателе из базы данных
func getTeacherInfo(db *sql.DB, fullName string) (*Teacher, error) {
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

	var teacher Teacher
	var subjects, departments string

	err := row.Scan(
		&teacher.ID, &teacher.Name, &teacher.BirthDate, &teacher.AlmaMater, &teacher.GraduationYear, &teacher.Degree,
		&teacher.KnowledgeScore, &teacher.TeachingSkillScore, &teacher.CommunicationScore, &teacher.LeniencyScore, &teacher.OverallScore,
		&teacher.KnowledgeRatingNum, &teacher.TeachingSkillRatingNum, &teacher.CommunicationRatingNum, &teacher.LeniencyRatingNum, &teacher.OverallRatingNum,
		&subjects, &departments,
	)
	if err != nil {
		return nil, err
	}

	// Обработка массивов subjects и departments
	if subjects != "{}" {
		teacher.Subjects = strings.Split(subjects[1:len(subjects)-1], ",")
	}
	if departments != "{}" {
		teacher.Departments = strings.Split(departments[1:len(departments)-1], ",")
	}

	return &teacher, nil
}

// Функция для поиска похожих имен преподавателей
func findSimilarTeachers(db *sql.DB, queryName string) ([]string, error) {
	// Получаем все имена преподавателей из базы данных
	rows, err := db.Query("SELECT name FROM teachers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allNames []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		allNames = append(allNames, name)
	}

	// Вычисляем схожесть для каждого имени
	type similarity struct {
		name       string
		similarity float64
	}

	var similarities []similarity
	for _, name := range allNames {
		dist := levenshtein.DistanceForStrings([]rune(queryName), []rune(name), levenshtein.DefaultOptions)
		maxLen := max(len(queryName), len(name))
		sim := 1 - float64(dist)/float64(maxLen)
		if sim >= similarityThreshold {
			similarities = append(similarities, similarity{name: name, similarity: sim})
		}
	}

	// Сортируем по убыванию схожести
	for i := 0; i < len(similarities); i++ {
		for j := i + 1; j < len(similarities); j++ {
			if similarities[i].similarity < similarities[j].similarity {
				similarities[i], similarities[j] = similarities[j], similarities[i]
			}
		}
	}

	// Выбираем не более maxSuggestions самых похожих имен
	var suggestions []string
	for i := 0; i < len(similarities) && i < maxSuggestions; i++ {
		suggestions = append(suggestions, similarities[i].name)
	}

	return suggestions, nil
}

// Функция для форматирования информации о преподавателе в строку
func formatTeacherInfo(teacher *Teacher) string {
	response := fmt.Sprintf("Имя: %s\n", teacher.Name)
	if teacher.BirthDate.Valid {
		response += fmt.Sprintf("Дата рождения: %s\n", teacher.BirthDate.String)
	}
	if teacher.AlmaMater.Valid {
		response += fmt.Sprintf("Альма-матер: %s\n", teacher.AlmaMater.String)
	}
	if teacher.GraduationYear.Valid {
		response += fmt.Sprintf("Год выпуска: %s\n", teacher.GraduationYear.String)
	}
	if teacher.Degree.Valid {
		response += fmt.Sprintf("Ученая степень: %s\n", teacher.Degree.String)
	}
	if teacher.KnowledgeScore.Valid {
		response += fmt.Sprintf("Оценка знаний: %.2f\n", teacher.KnowledgeScore.Float64)
	}
	if teacher.TeachingSkillScore.Valid {
		response += fmt.Sprintf("Оценка преподавания: %.2f\n", teacher.TeachingSkillScore.Float64)
	}
	if teacher.CommunicationScore.Valid {
		response += fmt.Sprintf("Оценка коммуникации: %.2f\n", teacher.CommunicationScore.Float64)
	}
	if teacher.LeniencyScore.Valid {
		response += fmt.Sprintf("Оценка снисходительности: %.2f\n", teacher.LeniencyScore.Float64)
	}
	if teacher.OverallScore.Valid {
		response += fmt.Sprintf("Общая оценка: %.2f\n", teacher.OverallScore.Float64)
	}
	if len(teacher.Subjects) > 0 {
		response += fmt.Sprintf("Предметы: %s\n", strings.Join(teacher.Subjects, ", "))
	}
	if len(teacher.Departments) > 0 {
		response += fmt.Sprintf("Кафедры: %s\n", strings.Join(teacher.Departments, ", "))
	}

	return response
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}