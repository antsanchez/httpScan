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
	"flag"
	"fmt"

	"github.com/antsanchez/httpScan/functions"
)

func main() {

	domain := flag.String("u", "", "URL to scan")
	times := flag.Int("n", 1, "Number of times to do the request")

	flag.Parse()

	if len(*domain) == 0 {
		fmt.Println("You need to indicate at least the URL to scan")
		flag.PrintDefaults()
		return
	}

	info := functions.GetInfo(*domain, *times)

	info.PrintTable()
}
