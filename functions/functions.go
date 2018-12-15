package functions

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/antsanchez/httpScan/model"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetInfo(domain string, times int) (info model.HttpInfo) {

	start := time.Now()

	res, err := http.Get(domain)
	check(err)

	badRequests := 0
	if res.StatusCode != 200 {
		badRequests++
	}

	if times > 1 {
		for i := 0; i < times; i++ {
			res, err := http.Get(domain)
			check(err)

			if res.StatusCode != 200 {
				badRequests++
			}
		}
	}

	elapsed := time.Now().Sub(start)
	duration := time.Duration(elapsed)

	if times > 1 {
		duration = duration / time.Duration(times)
	}

	host, err := url.Parse(domain)
	check(err)

	ips, err := net.LookupIP(host.Host)
	check(err)

	defer res.Body.Close()

	info.Requests = times
	info.BadRequests = badRequests

	info.URL = domain
	info.Host = host.Host
	info.Time = duration.String()
	info.Status = res.Status
	info.Protocol = res.Proto
	info.Uncompressed = strconv.FormatBool(res.Uncompressed)

	for _, ip := range ips {
		info.Ips = append(info.Ips, ip.String())
	}

	info.Headers = res.Header

	return
}
