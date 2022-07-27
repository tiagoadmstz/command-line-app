# Command Line App
A simple command line application to practice golang

# Commands
## ip
- Usage: Search for IP adresses on the internet
- Sintax: ip --host \<host-name>
### e.g.:
```
$ go run main.go ip --host amazon.com.br
54.239.26.87
72.21.203.171
52.94.225.243
2022/07/27 19:40:32 <nil>
```
## servers
- Usage: Search for servers names on the internet
- Sintax: servers --host \<host-name>
### e.g.:
```
$ go run main.go servers --host mercadolivre.com.br
ns-341.awsdns-42.com.
ns-653.awsdns-17.net.
ns-1828.awsdns-36.co.uk.
ns-1387.awsdns-45.org.
2022/07/27 19:37:32 <nil>
exit status 1
```
# Build Application
### Generating executable file
```
$ GOOS=windows go build
$ GOOS=linux go build
$ GOOS=darwin go build
```
### Execute (Windows)
```
$ ./command-line-app ip --host google.com.br
$ ./command-line-app servers --host google.com.br
```