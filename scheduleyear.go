package gamedayapi

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// OpeningAndFinalDatesForYear will inspect the schedule for the year and return dates for opening day and the final
// day of the season
func OpeningAndFinalDatesForYear(year int) (time.Time, time.Time) { // return an error
	var openingDay, finalDay time.Time
	f, err := os.Open("schedules/" + strconv.Itoa(year) + "SKED.TXT")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bf := bufio.NewReader(f)
	firstLine, _ := bf.ReadString('\n')
	openingDay = dateFromScheduleLine(firstLine)

	for {
		switch line, err := bf.ReadString('\n'); err {
		case nil:
		case io.EOF:
			if line > "" {
				finalDay = dateFromScheduleLine(line)
			}
			return openingDay, finalDay
		default:
			log.Fatal(err)
		}
	}

}

// Parses a line from the Retrosheet schedule files. Lines look like:
// "20110331","0","Thu","MIL","NL",1,"CIN","NL",1,"d","",""
func dateFromScheduleLine(line string) time.Time {
	year, _ := strconv.Atoi(line[1:5])
	month, _ := strconv.Atoi(line[5:7])
	day, _ := strconv.Atoi(line[7:9])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
