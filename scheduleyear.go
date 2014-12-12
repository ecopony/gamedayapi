package gamedayapi

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

// OpeningAndFinalDatesForYear will inspect the schedule for the year and return dates for opening day and the final
// day of the season
func OpeningAndFinalDatesForYear(year int) (time.Time, time.Time) { // return an error
	var openingDay, finalDay time.Time
	scheduleFilePath, err := scheduleFilePath(year)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(scheduleFilePath)
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

// TeamsForYear returns a list of valid team codes for the given year, sorted in alphabetical order.
func TeamsForYear(year int) []string {
	teams := make([]string, 0, 30)
	scheduleFilePath, err := scheduleFilePath(year)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(scheduleFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		teamFromLine := firstTeamFromScheduleLine(line)
		teams = appendIfMissing(teams, teamFromLine)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Strings(teams)
	return teams
}

func scheduleFilePath(year int) (string, error) {
	absPath, err := filepath.Abs("../gamedayapi/schedules")
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	buffer.WriteString(absPath)
	buffer.WriteString("/")
	buffer.WriteString(strconv.Itoa(year))
	buffer.WriteString("SKED.TXT")
	return buffer.String(), nil
}

// Parses a line from the Retrosheet schedule files. Lines look like:
// "20110331","0","Thu","MIL","NL",1,"CIN","NL",1,"d","",""
func dateFromScheduleLine(line string) time.Time {
	year, _ := strconv.Atoi(line[1:5])
	month, _ := strconv.Atoi(line[5:7])
	day, _ := strconv.Atoi(line[7:9])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func firstTeamFromScheduleLine(line string) string {
	return line[22:25]
}
