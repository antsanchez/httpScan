package functions

import (
	"fmt"
	"io/ioutil"
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
	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)

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

	info.Requests = times
	info.BadRequests = badRequests
	info.URL = domain
	info.Host = host.Host
	info.Time = duration.String()
	info.Status = res.Status
	info.Protocol = res.Proto
	info.Uncompressed = strconv.FormatBool(res.Uncompressed)
	info.BodySize = len(bodyBytes)

	for _, ip := range ips {
		info.Ips = append(info.Ips, ip.String())
	}

	info.Headers = res.Header

	return
}
