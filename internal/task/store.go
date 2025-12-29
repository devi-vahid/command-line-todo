package task

import (
	"encoding/csv"
	"strconv"
	"time"
)

func LoadTasks(path string) ([]Task, error) {
	f, err := loadFile(path)
	if err != nil {
		return nil, err
	}
	defer closeFile(f)
	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if nil != err {
		return nil, err
	}
	var tasks []Task
	for i, record := range records {
		if 0 == i {
			continue
		}
		id, _ := strconv.Atoi(record[0])
		description := record[1]
		created_at, _ := time.Parse(time.RFC3339, record[2])
		is_complete, _ := strconv.ParseBool(record[3])
		tasks = append(tasks, Task{
			ID:          id,
			Description: description,
			CreatedAt:   created_at,
			IsComplete:  is_complete,
		})
	}
	return tasks, nil
}

func SaveTasks(path string, tasks []Task) error {
	f, err := loadFile(path)
	if nil != err {
		return err
	}

	defer closeFile(f)

	f.Truncate(0)
	f.Seek(0, 0)

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"ID", "Description", "CreatedAt", "IsComplete"})

	for _, task := range tasks {
		w.Write([]string{
			strconv.Itoa(task.ID),
			task.Description,
			task.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(task.IsComplete),
		})
	}
	return nil
}
