package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "start, s",
			Value: "1",
			Usage: "Start Page",
		},
		cli.StringFlag{
			Name:  "end, e",
			Value: "1",
			Usage: "End Page",
		},
		cli.StringFlag{
			Name:  "lNumber, l",
			Value: "72",
			Usage: "Lines Number per Page",
		},
	}

	app.Action = func(c *cli.Context) error {
		fileName := ""
		if c.NArg() > 0 {
			fileName = c.Args().Get(0)
		}
		fmt.Printf("File Path: %s\n", fileName)
		startPage, err := strconv.Atoi(c.String("start"))
		endPage, err := strconv.Atoi(c.String("end"))
		lineNumber, err := strconv.Atoi(c.String("lNumber"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Start Page: %d\nEnd Page: %d\nLines Number per page: %d\n", startPage, endPage, lineNumber)

		startLine := lineNumber*(startPage-1) + 1
		endLine := lineNumber * endPage
		fi, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return nil
		}
		defer fi.Close()

		br := bufio.NewReader(fi)
		currentLine := 1
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			if currentLine >= startLine && currentLine <= endLine {
				fmt.Printf("%s\n", string(a))
			}
			currentLine++
			if currentLine > endLine {
				break
			}
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
