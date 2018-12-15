package model

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type HttpInfo struct {
	URL          string
	Host         string
	Time         string
	Status       string
	Protocol     string
	Uncompressed string
	Ips          []string
	Headers      map[string][]string
	Requests     int
	BadRequests  int
}

func (h *HttpInfo) PrintTable() {

	fmt.Println("")

	table := tablewriter.NewWriter(os.Stdout)
	table.Append([]string{"domain", h.URL})

	for _, ip := range h.Ips {
		table.Append([]string{"IP", ip})
	}

	if h.Requests > 1 {
		table.Append([]string{"N. of Requests", strconv.Itoa(h.Requests)})
		table.Append([]string{"Non HTTP 200 Responses", strconv.Itoa(h.BadRequests)})
		table.Append([]string{"Time Average", h.Time})
	} else {
		table.Append([]string{"Time", h.Time})
	}

	table.Append([]string{"Status", h.Status})
	table.Append([]string{"Protocol", h.Protocol})
	table.Append([]string{"Compressed", h.Uncompressed})
	table.SetHeader([]string{"Action / Header", "Value"})

	// Header Info
	setcookie := ""
	for i, val := range h.Headers {
		if strings.Compare(i, "Set-Cookie") == 0 {
			setcookie = strings.Join(val, " ")
		} else {
			table.Append([]string{i, strings.Join(val, " ")})
		}
	}

	table.SetHeader([]string{"Header", "Value"})
	table.SetRowLine(true)
	table.Render()

	// Set-Cookie value ill be show separately,
	// since it is often too long and breaks the table
	if strings.Compare(setcookie, "") != 0 {
		fmt.Println("Set-Cookie:")
		fmt.Println(setcookie)
	}

	fmt.Println("")
}
