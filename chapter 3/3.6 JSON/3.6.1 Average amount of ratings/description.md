На стандартный ввод подаются данные о студентах университетской группы в формате JSON:
```json
{
"ID": 134,
"Number": "ИЛМ-1274",
"Year": 2,
"Students": [{
"LastName": "Вещий",
"FirstName": "Лифон",
"MiddleName": "Вениаминович",
"Birthday": "4апреля1970года",
"Address": "632432,г.Тобольск,ул.Киевская,дом6,квартира23",
"Phone": "+7(948)709-47-24",
"Rating": [1, 2, 3]
},
{
// ...
}
]
}
```
В сведениях о каждом студенте содержится информация о полученных им оценках (Rating). Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы. Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:
````json
{
    "Average": 14.1
}
````
В выводе такой формат: префикс - пустая строка, отступ - 4 пробела.