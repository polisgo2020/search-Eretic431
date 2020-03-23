package main

import (
	"encoding/json"
	"fmt"
	"github.com/caneroj1/stemmer"
	"github.com/polisgo2020/Akhmedov_Abdulla/invertedIndex"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

var invertedIn invertedIndex.Index

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Not enough arguments")
	}

	file, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(file, &invertedIn)
	if err != nil {
		log.Println(err)
		return
	}

	stopWords, err := invertedIndex.ReadStopWords(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	var searchPhrase []string
	for i := 3; i < len(os.Args); i++ {
		if _, ok := stopWords[os.Args[i]]; ok {
			continue
		}

		tmp := strings.ToLower(stemmer.Stem(os.Args[i]))
		if _, ok := invertedIn[tmp]; ok {
			searchPhrase = append(searchPhrase, tmp)
		}
	}

	var answer []float64
	answerMap := make (map[float64][]string)
	if len(searchPhrase) == 1 {
		if filesMap, ok := invertedIn[searchPhrase[0]]; ok {
			for file := range filesMap {
				tmp := float64(len(filesMap[file]))

				answer = append(answer, tmp)
				if answerMap[tmp] == nil {
					answerMap[tmp] = make([]string, 0, 0)
				}

				answerMap[tmp] = append(answerMap[tmp], file)
			}

			_ = sort.Reverse(sort.Float64Slice(answer))
			for _, v := range answer {
				files := answerMap[v]
				for _, file := range files {
					fmt.Printf("%s - %f\n", file, v)
				}
			}
		} else {
			fmt.Print("None of files contains this search-phrase")
		}
	} else {
		tmp := invertedIn[""]
		for file, _ := range tmp {
			res := getInfo(searchPhrase, file)
			answer = append(answer, res)

			if answerMap[res] == nil {
				answerMap[res] = make([]string, 0, 0)
			}

			answerMap[res] = append(answerMap[res], file)
		}

		sort.Float64s(answer)
		for _, v := range answer {
			files := answerMap[v]
			for _, file := range files {
				if v > 0.0000001 {
					fmt.Printf("%s - %f\n", file, v)
				}
			}
		}
	}
}

func getInfo(phrase []string, file string) float64 {
	distance, count := findMinWay(phrase, 0, file, 1)
	return float64(distance) / float64(count)
}

func findMinWay(phrase []string, index int, file string, count int) (int, int) {
	if index >= len(phrase) {
		return 0, count
	}

	curIndex := findFirstExistTokenIndex(phrase, index, file)
	nextIndex := findFirstExistTokenIndex(phrase, curIndex+1, file)
	if nextIndex == -1 || curIndex == -1 {
		return 0, count
	}

	curList := invertedIn[phrase[curIndex]][file]
	nextList := invertedIn[phrase[nextIndex]][file]

	min := 99999
	var resCount int
	for _, v1 := range curList {
		var res int
		// TODO: кэшировать значения этого вызова в матрицу[index][count] -> сильно ускорит обход дерева
		res, resCount = findMinWay(phrase, nextIndex, file, count+1)
		for _, v2 := range nextList {
			delta := abs(v2 - v1)
			min = fMin(min, res+delta)
		}
	}

	return min, resCount
}

func fMin(a int, b int) int {
	if a <= b {
		return a
	}

	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}

func findFirstExistTokenIndex(phrase []string, index int, file string) int {
	for i := index; i < len(phrase); i++ {
		if _, ok := invertedIn[phrase[index]][file]; ok {
			return index
		}
	}

	return -1
}
