package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ProcessArticles( 
	filename string, 
	maxCount int, 
	handler func(articleLines []string, title, id string, links []string) error,
) error {

	fmt.Println("Processing articles from:", filename)

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var articleLines []string
	inPage := false
	count := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if strings.Contains(line, "<page>") {
			inPage = true
			articleLines = []string{line}
		} else if strings.Contains(line, "</page>") && inPage {
			articleLines = append(articleLines, line)
			inPage = false

			title := extractTitle(articleLines)
			id := extractID(articleLines)
			links := extractLinks(articleLines)

			if err := handler(articleLines, title, id, links); err != nil {
				return err
			}

			count++
			if count >= maxCount {
				break
			}

		} else if inPage {
			articleLines = append(articleLines, line)
		}
	}

	return nil
}
