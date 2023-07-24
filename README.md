# pathfinder

Fuzz a list of URLs for potential path traversal/normalization issues.

Multiple comparisons are performed to establish a confidence level for findings that should be manually investigated.

## Install

```
▶ go install github.com/dptsec/pathfinder@latest
```

## Build locally

```
$ cd pathfinder
$ go build
```


## Usage

pathfinder accepts input from `stdin` with 1 URL per line. 

Input should be prepared beforehand with `awk -F '?' '{print $1}' | sed 's![^/]*$!!' | sort -u` to remove files and parameters.

Any lines that have no `scheme` will have `https://` prepended:

```
$ cat samples/input.txt
http://nmap.scanme.org
https://nmap.scanme.org
$ cat samples/input.txt | pathfinder
[*] Configuration:
- Confidence cut-off:   3
- Request method:       GET
- Proxy URL:
- Rate-limiting:        50 requests/sec
- Request timeout:      10 seconds
- Total threads:        5
- HTTP User-Agent:      Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36
- Stop on errors:       false

[*] Running 6 total queued jobs
[*] StatusCode mismatch:
        200 != 404
[*] Levenshtein distance:
        6734
[*] Word count:
        24 != 495
[*] Header mismatch: Content-Type
        text/html != text/html; charset=iso-8859-1

[*] Potential hit: http://nmap.scanme.org/%2e%2e%2f
```

An initial baseline request is sent to the first URL that will be used in further comparisons against the payloads.

## Options

The following options can be specified on the command line:
+ `-C`: Specify cookie data in the form of "COOKIENAME=VALUE;"
+ `-T`: HTTP request timeout in seconds
+ `-c`: Confidence cut-off level for positive findings
+ `-e`: Stop fuzzing when >=75% of requests have returned a 403 errors
+ `-m`: HTTP method to use
+ `-o`: Output file for results
+ `-p`: HTTP or SOCKS proxy for requests
+ `-r`: Rate-limit requests per second
+ `-t`: Number of concurrent threads
+ `-u`: User-Agent string
+ `-v`: Verbose mode - Show results of all confidence level checks

## TODO

+ Add support for multiple hosts on input. Currently only 1 baseline is set for comparisons
+ Output options
+ Better detection for traversals that aren't a result of path normalization
+ nginx alias traversal support with wordlist
+ Parameter fuzzing

## References
***[Breaking Parser Logic - Orange Tsai, 2018](https://i.blackhat.com/us-18/Wed-August-8/us-18-Orange-Tsai-Breaking-Parser-Logic-Take-Your-Path-Normalization-Off-And-Pop-0days-Out-2.pdf)
