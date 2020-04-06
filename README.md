# Akhmedov Abdulla GoLang homeworks

### Примеры запуска из директории `build`
* `go run build.go ../inputFiles`
* `go run build.go -sw=../stopWords.txt ../inputFiles`
* `go run build.go -s ../inputFiles/t0.txt`
* `go run build.go -s -sw=../stopWords.txt ../inputFiles/t0.txt`

# Алгоритм поиска фразы
На этапе формирования обратного индекса берется основа слова, исключаются
шумовые слова, запоминается позиция каждого слова в файле.

Из пришедшей поисковой фразы убираются шумовые слова, от каждого слова берется
основа. Далее, если из поисковой фразы осталось только одно слово, то
тот файл лучше, в котором это слово встречается чаще. Иначе с помощью
поиска в глубину ищем кратчайщий путь между словами из поисковой фразы
(причем неважно в какой последовательности они стоят в файле) и делим эту
величину на количество слов, принимавших участие в этом пути. Таким образом,
наилучшим файлом будет тот, у которого это отношение меньше, то есть слова находятся кучнее.

### Примеры запуска из директории `search`
* `go run search.go ../stopWords.txt ../hm1/outputJSON.txt aspect association`
* `go run search.go ../stopWords.txt ../hm1/outputJSON.txt Gaidai`
* `go run search.go ../stopWords.txt ../hm1/outputJSON.txt Gaidai generation`

### Пример запуска из корня репозитория
Если аргументом не передадут стоп-слова или файл с обратным индексом, возьмутся дефолтные пути к этим файла, с учётом
того, что программа запускается из корня
* `go run ./web/web.go -sw=stopWords.txt -index=outputJSON.txt :8080`
* `go run ./web/web.go -index=outputJSON.txt :8080`
* `go run ./web/web.go :8080`

![Alt text](./golang.png)
