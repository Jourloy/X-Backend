# Как писать Swagger комментарии к эндпоинтам

## Комментарий хендлера

-   `@Tag.name` - Название категории на русском
-   `@Tag.description` - Краткое описание цели категории

```golang
// @Tag.name Пример
// @Tag.description Эндпоинты для показания примера написания комментариев
```

Такие комментарии указываются в `handlers` в файле, относящемся к категории

## Комментарий контроллера

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