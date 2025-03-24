package teacher

import (
	"fmt"
	"strings"
)

// Функция для форматирования информации о преподавателе в строку
func FormatTeacherInfo(teacher *Teacher) string {
	message := fmt.Sprintf("Имя: %s\n\n", teacher.Name)
	if teacher.BirthDate.Valid {
		message += fmt.Sprintf("Дата рождения: %s\n", 
								teacher.BirthDate.String[:10])
	}
	if teacher.AlmaMater.Valid {
		message += fmt.Sprintf("Альма-матер: %s\n", 
								teacher.AlmaMater.String)
	}
	if teacher.GraduationYear.Valid {
		message += fmt.Sprintf("Год выпуска: %s\n", 
								teacher.GraduationYear.String)
	}
	if teacher.Degree.Valid {
		message += fmt.Sprintf("Ученая степень: %s\n\n", 
								teacher.Degree.String)
	}
	if len(teacher.Departments) > 0 {
		message += fmt.Sprintf(fmt.Sprintf("Кафедры: %s\n", 
								strings.Join(teacher.Departments, ", ")))
	}
	if len(teacher.Subjects) > 0 {
		message += fmt.Sprintf("Предметы: %s\n\n", 
								strings.Join(teacher.Subjects, ", "))
	}
	if teacher.KnowledgeScore.Valid {
		message += fmt.Sprintf("Знания: %.2f (%d)\n", 
								teacher.KnowledgeScore.Float64,
								teacher.KnowledgeRatingNum.Int64)
	}
	if teacher.TeachingSkillScore.Valid {
		message += fmt.Sprintf("Умение преподавать: %.2f (%d)\n", 
								teacher.TeachingSkillScore.Float64,
								teacher.TeachingSkillRatingNum.Int64)
	}
	if teacher.CommunicationScore.Valid {
		message += fmt.Sprintf("В общении: %.2f (%d)\n", 
								teacher.CommunicationScore.Float64,
								teacher.CommunicationRatingNum.Int64)
	}
	if teacher.LeniencyScore.Valid {
		message += fmt.Sprintf("Халявность: %.2f (%d)\n", 
								teacher.LeniencyScore.Float64,
								teacher.LeniencyRatingNum.Int64)
	}
	if teacher.OverallScore.Valid {
		message += fmt.Sprintf("Общая оценка: %.2f (%d)\n", 
								teacher.OverallScore.Float64,
								teacher.OverallRatingNum.Int64)
	}

	return message
}
