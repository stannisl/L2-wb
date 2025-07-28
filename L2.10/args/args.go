package args

import (
	"errors"
	"strconv"
	"strings"
)

/*
Flags содержит настройки, определяющие поведение сортировки.
-k N — сортировать по столбцу (колонке) №N (разделитель — табуляция по умолчанию).
Например, «sort -k 2» отсортирует строки по второму столбцу каждой строки.

-n — сортировать по числовому значению (строки интерпретируются как числа).

-r — сортировать в обратном порядке (reverse).

-u — не выводить повторяющиеся строки (только уникальные).
*/
type Flags struct {
	SortByColumn bool // k flag
	Column       int

	SortByNumber bool // n flag
	ReverseSort  bool // r flag
	UniqueSort   bool // u flag
}

// ParseArgs обрабатывает аргументы командной строки и возвращает структуру Flags.
func ParseArgs(flags []string) (Flags, []string, error) {
	parsedFlags := Flags{Column: -1}
	extractColumnNumber := false
	var nonFlags []string

	for _, arg := range flags {

		// Parsing single flags. ex "-k, -n, -u"
		if strings.HasPrefix(arg, "-") {
			for _, ch := range arg[1:] {
				switch ch {
				case 'n':
					parsedFlags.SortByNumber = true
				case 'r':
					parsedFlags.ReverseSort = true
				case 'u':
					parsedFlags.UniqueSort = true
				case 'k':
					{
						parsedFlags.SortByColumn = true
						extractColumnNumber = true
					}
				}
			}
		} else if extractColumnNumber {
			column, err := strconv.Atoi(arg)
			if err == nil {
				parsedFlags.Column = column
				extractColumnNumber = false
			} else {
				return Flags{}, nil, errors.New("wrong argument after -k flag. Should be number to sorting by columns")
			}
		} else {
			nonFlags = append(nonFlags, arg)
		}
	}

	return parsedFlags, nonFlags, nil
}
