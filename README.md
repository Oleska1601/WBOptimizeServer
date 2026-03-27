
# Оптимизация простого API-сервиса с профилировкой

## Задание
Взять простой HTTP API (например, сервис, который складывает числа или возвращает JSON, или сервис из L0), создать нагрузку и оптимизировать его по CPU и памяти.

- pprof и net/http/pprof
- benchstat, go test -bench
- анализ trace

Результат: проект с кодом API, профилировкой с историей коммитов, README с описанием изменений.

## Реализация
### Версия V1 (до оптимизации)
V1 представляет собой реализацию с проблемами производительности:

**CPU-bound (Fibonacci):**
- Использует рекурсивный алгоритм с экспоненциальной сложностью O(2ⁿ) = Каждый вызов создает глубокий стек рекурсии

**Memory-bound (JSON):**

- Создает новый буфер на каждый запрос через `&bytes.Buffer{}`
- выполняет копирование данных при сериализации
- Отсутствует переиспользование памяти


### Версия V2 (после оптимизации)
V2 содержит оптимизированные алгоритмы и эффективное управление памятью:

**CPU-bound (Fibonacci):**
- Использует итеративный алгоритм с линейной сложностью O(n) = Отсутствие рекурсии

**Memory-bound (JSON):**
- Применяет sync.Pool для переиспользования буферов
- Предварительное выделение памяти (1024 байта)
- Работа с указателями вместо копирования структур
- Работа с выравниваем структур в памяти

### Бенчмарки
```bash
make bench-all
```

```bash
go test -bench=. -benchmem ./internal/service/v1/...
goos: windows
goarch: amd64
pkg: github.com/Oleska1601/WBOptimizeServer/internal/service/v1
cpu: AMD Ryzen 5 4500U with Radeon Graphics
BenchmarkFibonacci-6                 259           4579825 ns/op               0 B/op          0 allocs/op
BenchmarkProcessJSON-6            667281             13343 ns/op             688 B/op          6 allocs/op
PASS
ok      github.com/Oleska1601/WBOptimizeServer/internal/service/v1      10.823s
```

```bash
go test -bench=. -benchmem ./internal/service/v2/...
goos: windows
goarch: amd64
pkg: github.com/Oleska1601/WBOptimizeServer/internal/service/v2
cpu: AMD Ryzen 5 4500U with Radeon Graphics 
BenchmarkFibonacci-6            60297265                17.16 ns/op            0 B/op          0 allocs/op
BenchmarkProcessJSON-6            854785              1541 ns/op             304 B/op          3 allocs/op
PASS
ok      github.com/Oleska1601/WBOptimizeServer/internal/service/v2      2.547s
```

### Сравнение результатов
**CPU-bound (Fibonacci)**

| Метрика | V1         | V2       | Улучшение  |
| ------- | ---------- | -------- | ---------- |
| Время   | 4579825 ns | 17.16 ns | ↓ 99.9996% |

**Memory-bound (JSON)**

| Метрика   | V1       | V2      | Улучшение |
| --------- | -------- | ------- | --------- |
| Время     | 13343 ns | 1541 ns | ↓ 88.5%   |
| Память    | 688 B    | 304 B   | ↓ 55.8%   |
| Аллокации | 6        | 3       | ↓ 50%     |


### Запуск
1. Запуск приложения:
```bash
make up 
```
2. Запуск нагрузки на приложение + сбор профилей:
```bash
make run 
```

3. Просмотр результатов профилирования:
**cpu**
```bash
make view-cpu-v1
make view-cpu-v2
make view-cpu-compare
```

**heap**
```bash
make view-heap-v1
make view-heap-v2
make view-heap-compare
```

**allocs**
```bash
make view-allocs-v1
make view-allocs-v2
```

**trace**
```bash
make view-trace-v1
make view-trace-v2
```

