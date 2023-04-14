# coinGecko
## Задание:
- ### Парсинг Cryptorank:
  + Ресурс: [Cryptorank](https://cryptorank.io/)
  + Данные для парсинга: теги нескольких валют (первых трех)
  + Метод парсинга: любой
  + Метод полученных результатов: Запись в гугл таблицы (API) по запуску
  + Количество столбцов 3: Наименование, Теги, Timestamp.
  
- ### Парсинг CoinGecko:
  + Ресурс: [CoinGecko](https://www.coingecko.com)
  + Данные для парсинга: валюты их стоимость относительно доллара
  + Метод парсинга: любой
  + Метод полученных результатов: Запись в гугл таблицы (API) по запуску
  + Количество столбцов 65: Наименование, Теги, Timestamp.
___
## Запуск программы:

*make run*


for parsing 3 first position on [Cryptorank](https://cryptorank.io/)

localhost:8080/crypto_rank


for parsing 65 first position on [CoinGecko](https://www.coingecko.com)

localhost:8080/coin_gecko

