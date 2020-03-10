package inverted_index

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"unicode"
)

func readFiles(flag bool, files []string) (map[string]string, error) {
	m := make(map[string]string)
	i := 0

	if flag {
		for _, v := range files {
			data, err := ioutil.ReadFile(v)
			if err != nil {
				log.Print(err, "Could not read file!")
				return nil, err
			}

			tmp := strings.Split(v, "/")
			fileName := tmp[len(tmp)-1]
			m[fmt.Sprint(i)+"_"+fileName] = string(data)
			i++
		}
	} else {
		for _, v := range files {
			dir, err := ioutil.ReadDir(v)
			if err != nil {
				log.Print(err, "Could not read directory!")
				return nil, err
			}

			for _, file := range dir {
				data, err := ioutil.ReadFile(filepath.Join(v, file.Name()))
				if err != nil {
					log.Print(err, "Could not read file!")
					return nil, err
				}

				m[strings.Join([]string{fmt.Sprintf("%d", i), "_", file.Name()}, "")] = string(data)
				i++
			}
		}
	}

	return m, nil
}

// returns inverted index map that also stores position of each token in document
func GetInvertedIndex(flag bool, files []string) (map[string]map[string][]int, error) {
	invertedIndex := make(map[string]map[string][]int)
	filesMap, err := readFiles(flag, files)
	if err != nil {
		return nil, err
	}

	for file, str := range filesMap {
		tokens := strings.Fields(str)
		for position, token := range tokens {
			token = strings.TrimFunc(token, func(r rune) bool {
				return !unicode.IsLetter(r)
			})
			token = strings.ToLower(token)

			if invertedIndex[token] == nil {
				invertedIndex[token] = make(map[string][]int)
			}

			invertedIndex[token][file] = append(invertedIndex[token][file], position)
		}
	}

	return invertedIndex, nil
}
