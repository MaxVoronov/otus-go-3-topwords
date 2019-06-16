# Otus Homework #3
## Частотный анализ

Написать функцию, которая получает на вход текст и возвращает.  
10 самых часто встречающихся слов без учета словоформ.

### Пример использования

```go
package main

import (
	"fmt"
	"github.com/maxvoronov/otus-go-3-topwords"
)

const TextSample = `
	Вот дом, который построил Джек.
	А это пшеница, Которая в тёмном чулане хранится
	В доме, который построил Джек.
	А это весёлая птица-синица,
	Которая часто ворует пшеницу,
	Которая в тёмном чулане хранится
	В доме, который построил Джек.
	Вот кот, который пугает и ловит синицу,
	Которая часто ворует пшеницу,
	Которая в тёмном чулане хранится
	В доме, который построил Джек.`

func main() {
	top10 := topwords.GetTopWords(TextSample, 10)

	fmt.Println("Toп-10 слов:")
	for _, word := range top10 {
		fmt.Printf("  - %s (total %d)\n", word.Name, word.Count)
	}
}

```

Результат:
```
Top-10 words:
  - в (total 6)
  - который (total 5)
  - которая (total 5)
  - построил (total 4)
  - джек (total 4)
  - хранится (total 3)
  - чулане (total 3)
  - тёмном (total 3)
  - доме (total 3)
  - ворует (total 2)
```

### Полезные ссылки
- [Golang: Sort package](https://golang.org/pkg/sort/)
- [The 3 ways to sort in Go](https://yourbasic.org/golang/how-to-sort-in-go/)
