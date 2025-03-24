package teacher

import (
	"database/sql"
	levenshtein "github.com/texttheater/golang-levenshtein/levenshtein"

)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Функция для поиска похожих имен преподавателей
func FindSimilarTeachers(db *sql.DB, queryName string) ([]string, error) {
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
	var similarities []Similarity
	for _, name := range allNames {
		dist := levenshtein.DistanceForStrings([]rune(queryName), []rune(name), levenshtein.DefaultOptions)
		maxLen := max(len(queryName), len(name))
		sim := 1 - float64(dist)/float64(maxLen)
		if sim >= SimilarityThreshold {
			similarities = append(similarities, Similarity{name: name, similarity: sim})
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

	// Выбираем не более MaxSuggestions самых похожих имен
	var suggestions []string
	for i := 0; i < len(similarities) && i < MaxSuggestions; i++ {
		suggestions = append(suggestions, similarities[i].name)
	}

	return suggestions, nil
}