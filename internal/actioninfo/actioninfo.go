package actioninfo

import "fmt"

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Printf("Ошибка парсинга данных %d: %v\n", i+1, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Ошибка формирования информации для данных %d: %v\n", i+1, err)
			continue
		}

		fmt.Print(info)
		fmt.Println("---")
	}
}
