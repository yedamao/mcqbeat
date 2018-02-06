# Mcqbeat

Welcome to Mcqbeat.

Mcqbeat is based on [Beats](https://github.com/elastic/beats) ship   memcacheq queue stats to elasticsearch.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/yedamao/mcqbeat`

## Getting Started with Mcqbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7
* [Beats](https://github.com/elastic/beats) 6.0

### Init Project
To get running with Mcqbeat and also install the
dependencies, run the following command:

```
make setup
```


### Build

To build the binary for Mcqbeat run the command below. This will generate a binary
in the same directory with the name mcqbeat.

```
make
```


### Run

To run Mcqbeat with debugging output enabled, run:

```
./mcqbeat -c mcqbeat.yml -e -d "*"
```


### Test

To test Mcqbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Kibana Screenshots
![image](https://user-images.githubusercontent.com/8220938/35849582-14bb83f0-0b5d-11e8-95b6-43715a352cfb.png)
