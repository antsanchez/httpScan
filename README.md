# What is httpscan?

httpScan is a very simple CLI app which accepts a URL and print some Meta data from it, like the headers or the time the URL took to respond.

To execute, just type in your CLI the command and indicate with the flag -u the url to scan: 

```
$ httpscan -u https://tool-seo.com
```

If you want to perform more than one request in order to test the response time of the URL, add the flag -n with the number of requests to perform:

```
$ httpscan -u https://tool-seo.com -n 5
```

### Download and Install

#### Install From Source

Just download and compile it as any other standard Go software. More info about Go: https://golang.org/
