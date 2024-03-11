# Регламент

## Как добавить модель

1. Инициализация модели в `repositories.go`

Всегда любым моделям нужно добавлять 4 поля: `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`

```golang
// Модель примера
type Example struct {
	// Основа
	ID string `json:"id" gorm:"primarykey"`

	// Динамические поля
	Name string `json:"name"`

	// Дети
	Childs []Child `json:"childs"`

	// Родители
	ExampleID string `json:"exmapleId"`
	ParentID string `json:"parentId"

	// Мета информация
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
```

Модель обязательно должна иметь структуру для создания. В нее помещаются все динамические поля, им прописывается JSON на случай, если это будет POST запрос

```golang
type ExampleCreate struct {
	Name string `json:"name"
}
```

У модели обязательно должна быть структура для поиска. В нее добавляется `Limit`, динамические поля и отношения. Все типы должны быть с `*`

```golang
// Структура поиска примера
type ExampleGet struct {
	ID *string

	// Динамические поля
	Name *string

	// Отношения
	ExampleID *string
	ParentID *string

	// Обязательно
	Limit *int
}
```

У модели обязательно должна быть структура репозитория. Все функции должны возвращать ошибку

```golang
// Репозиторий примера
type ExampleRepository interface {
	Create(example *ExampleCreate) (*Example, error)
	GetOne(query *ExampleGet) (*Example, error)
	GetAll(query *ExampleGet) (*[]Example, error)
	UpdateOne(example *Example) error
	DeleteOne(example *Example) error
}
```

2. Добавление репозитория

В папке `repositories` необходимо создать папку `имя_модели_rep` (`example_rep`),
а в ней аналогичный файл с расширением .go

Файл должен содержать определение репозитория

```golang
var Repository repositories.ExampleRepository
```

Структуру репозитория

```golang
type ExampleRepository struct {
	db gorm.DB
}
```

И функцию инициализации вместе с автомиграцией

```golang
// Init создает репозиторий примера
func Init() {
	go migration()

	Repository = &ExampleRepository{
		db: *storage.Database,
	}
}

func migration() {
	if err := Database.AutoMigrate(
		&repositories.Example{},
	); err != nil {
		logger.Fatal(`Migration failed`)
	}
}
```

3. Инициализация репозитория

Далее в `server.go` в функцию `initReps` нужно добавить инициализацию только что созданного репозитория

```golang
go example_rep.Init()
```

## Как писать Swagger комментарии к эндпоинтам

### Комментарий хендлера

-   `@Tag.name` - Название категории на русском
-   `@Tag.description` - Краткое описание цели категории

```golang
// @Tag.name Пример
// @Tag.description Эндпоинты для показания примера написания комментариев
```

Такие комментарии указываются в `handlers` в файле, относящемся к категории

### Комментарий контроллера

1. Сначала идет общий комментарий функции
2. `@Tags` - определяет категорию
3. `@Summary` - Краткое описание эндпоинта
4. `@Success 200 {object} СТРУКТУРА "КОММЕНТАРИЙ"` - Пример ответа сервера с кодом 200
5. `@Failure КОД {object}` - Пример ответа сервера с кодом отличным от 200
6. `@Router ПУТЬ [МЕТОД]` - как вызвать эндпоинт

```golang
// Create создает пример
//
// @Tags Пример
// @Summary Создает пример
// @Success 200 {object} CreateResponse200 "Успех"
// @Failure 400 {object} CreateResponse400 "Ошибка, смотри параметр error"
// @Router /example [post]
func (s *Controller) Create(c *gin.Context) {
```

Такие комментарии указываюся в `modules` в файле, относящемся к категории и являющийся контроллером (`категория.go` в названии)
