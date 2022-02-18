package main

import (
	"encoding/csv"
	"fmt"
	"go-script-template/pkg/netHttp"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2" // imports as package "cli"
)

func main() {
	fmt.Println("I am GoLang script!")

	fmt.Println("### CLI example ###")
	cliExample()

	fmt.Println("### CSV read ###")
	records := csvRead("./data/example.csv")
	fmt.Println(records)

	fmt.Println("### CSV write ###")
	newFilePath := "tmp/new.csv"
	csvWrite(records, newFilePath)
	content, err := ioutil.ReadFile(newFilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	fmt.Println("### net/http ###")
	body := netHttp.GetBody("https://raw.githubusercontent.com/exsemt/ruby-script-template/master/data/example.json")
	fmt.Println("Body:", body)
}

func cliExample() {
	// https://github.com/urfave/cli/blob/master/docs/v2/manual.md
	app := &cli.App{
		Name:  "GoLang script",
		Usage: "CLI example",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Value:   "Foo",
				Usage:   "Your name",
			},
		},
		Action: func(c *cli.Context) error {
			name := c.String("name")
			fmt.Printf("Hello %q, welcome to the world of CLI.", name)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func csvRead(filePath string) []csvRow {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	records := make([]csvRow, 0)
	for _, row := range rows {
		if row[0] == "id" {
			continue
		}

		id, _ := strconv.ParseInt(row[0], 10, 64)
		foo, _ := strconv.ParseInt(row[1], 10, 64)
		bar, _ := strconv.ParseInt(row[2], 10, 64)

		records = append(records, csvRow{id: id, foo: foo, bar: bar})
	}

	return records
}

type csvRow struct {
	id  int64
	foo int64
	bar int64
}

func csvWrite(records []csvRow, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"id", "foo", "bar"})
	for _, record := range records {
		newRow := []string{fmt.Sprint(record.id + 1), fmt.Sprint(record.foo + 2), fmt.Sprint(record.bar + 3)}
		if err := w.Write(newRow); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
