package services

import (
	"fmt"
	"github.com/eampleev23/yatr/internal/api_requests"
	"github.com/eampleev23/yatr/internal/client_config"
	"github.com/eampleev23/yatr/internal/models"
	"github.com/eampleev23/yatr/internal/my_csv"
	"strconv"
	"unicode/utf8"
)

func GenerateIssues(c *client_config.Config) error {

	fmt.Println("Читаем файл из конфига и сохраняем в [][]string..")
	result := my_csv.CsvParse(c.FilePath)
	fmt.Println("Создаем массив пустых моделей..")
	var newIssues []models.NewIssue

	fmt.Println("Начинаем перебирать строки из файла, начиная с первой")

	fmt.Println("Заводим счетчик пропусков с нулевым значением..")
	countOfSkips := 0

	for i := 1; i < len(result); i++ {

		fmt.Println("Старт итерации с i=", i)

		if result[i][1] != "" {
			fmt.Println("По задаче", result[i][3], "уже заполнен KEY:", result[i][1], ". Пропускаем...")
			fmt.Println("Увеличиваем счетчик пропусков на 1..")
			countOfSkips++
			continue
		}
		fmt.Println("Создаем новую пустую модель и заполняем ее данными..")
		newIssues = append(newIssues, models.NewIssue{})
		newIssues[i-countOfSkips-1].Queue = result[i][2]
		newIssues[i-countOfSkips-1].Summary = result[i][3]
		newIssues[i-countOfSkips-1].Type = result[i][4]
		prj, err := strconv.Atoi(result[i][5])
		if err != nil {
			fmt.Errorf("strconv fail %s", err.Error())
		}
		newIssues[i-countOfSkips-1].Project = prj
		newIssues[i-countOfSkips-1].Start = trimFirstRune(result[i][6])
		newIssues[i-countOfSkips-1].DueDate = trimFirstRune(result[i][7])
		newIssues[i-countOfSkips-1].Description = result[i][8]
		newIssues[i-countOfSkips-1].Assignee = result[i][9]
		newIssues[i-countOfSkips-1].Author = result[i][10]
		newIssues[i-countOfSkips-1].Parent = result[i][11]
		newIssues[i-countOfSkips-1].Priority = result[i][12]

		fmt.Println("Отправляем запрос в апи трекера и получаем key созданной задачи..")
		createdKey, err := api_requests.Create(c, newIssues[i-countOfSkips-1])
		if err != nil {
			return fmt.Errorf("create issues: %w", err)
		}
		fmt.Println("Заносим key в массив..")
		result[i][1] = createdKey
	}
	fmt.Println("Перезаписываем файл с уже созданными keys задач..")
	err := my_csv.CsvSave(c.FilePath, result)
	if err != nil {
		return fmt.Errorf("save issues: %w", err)
	}
	fmt.Println("Работа успешно выполнена, сэр.")
	return nil
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
