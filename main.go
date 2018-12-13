// Copyright 2018 Antonio Sanchez. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

func main() {
	doAndPrintRequest(os.Args[1])
}

func doAndPrintRequest(domain string) {

	start := time.Now()

	res, err := http.Get(domain)
	check(err)

	host, err := url.Parse(domain)
	check(err)

	ips, err := net.LookupIP(host.Host)
	check(err)

	elapsed := time.Now().Sub(start)

	defer res.Body.Close()

	fmt.Println("")

	// Request Info
	table := tablewriter.NewWriter(os.Stdout)
	table.Append([]string{"domain", domain})

	for _, ip := range ips {
		table.Append([]string{"IP", ip.String()})
	}

	table.Append([]string{"Time", elapsed.String()})
	table.Append([]string{"Status", res.Status})
	table.Append([]string{"Protocol", res.Proto})
	table.Append([]string{"Compressed", strconv.FormatBool(res.Uncompressed)})
	table.SetHeader([]string{"Action / Header", "Value"})

	// Header Info
	setcookie := ""
	for i, val := range res.Header {
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

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
