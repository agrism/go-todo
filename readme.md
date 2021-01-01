
# Before start

### 1. Build:

```bigquery
GOOS=linux GOARCH=amd64 go build -o go-todo-linux
```
### 2. Run docker
```bigquery
docker-compose up --build
```

### 3. Init database query

```bash
curl --request GET http://127.0.0.1:8089/init
```

# Use

#### 1. Create record

```bash
curl --header "Content-Type: application/json" \
  --request POST --data '{"name":"test name", "nickname": "test-nick-name"}' \
  http://127.0.0.1:8089/person
```

#### 2. Read all records

```bash
curl --request GET http://127.0.0.1:8089/person
```

#### 3. Read one record

```bash
curl --request GET http://127.0.0.1:8089/person/1
```

# Benchmarks

### 1. Read test

```bash
ab -n 100000 -k -c 30 -q http://127.0.0.1:8089/person/1
```

### 1.1. Result:

```bash
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient).....done


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8089

Document Path:          /person/1
Document Length:        58 bytes

Concurrency Level:      30
Time taken for tests:   664.587 seconds
Complete requests:      100000
Failed requests:        0
Keep-Alive requests:    100000
Total transferred:      19000000 bytes
HTML transferred:       5800000 bytes
Requests per second:    150.47 [#/sec] (mean)
Time per request:       199.376 [ms] (mean)
Time per request:       6.646 [ms] (mean, across all concurrent requests)
Transfer rate:          27.92 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       2
Processing:     7  199  69.3    196    1406
Waiting:        7  199  69.3    196    1406
Total:          7  199  69.3    196    1406

Percentage of the requests served within a certain time (ms)
  50%    196
  66%    205
  75%    214
  80%    229
  90%    291
  95%    304
  98%    381
  99%    398
 100%   1406 (longest request)
```

### 2. Write test

```bash
ab -p benchmark-post.json -T application/json -n 100000 -k -c 30 -q http://127.0.0.1:8089/person
```

### 2.1. Result:

```bash
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient).....done


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8089

Document Path:          /person
Document Length:        73 bytes

Concurrency Level:      30
Time taken for tests:   720.999 seconds
Complete requests:      100000
Failed requests:        0
Keep-Alive requests:    100000
Total transferred:      20500000 bytes
Total body sent:        22400000
HTML transferred:       7300000 bytes
Requests per second:    138.70 [#/sec] (mean)
Time per request:       216.300 [ms] (mean)
Time per request:       7.210 [ms] (mean, across all concurrent requests)
Transfer rate:          27.77 [Kbytes/sec] received
                        30.34 kb/s sent
                        58.11 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       2
Processing:    13  216  67.5    203     817
Waiting:       11  216  67.5    203     817
Total:         13  216  67.5    203     817

Percentage of the requests served within a certain time (ms)
  50%    203
  66%    216
  75%    278
  80%    285
  90%    299
  95%    311
  98%    386
  99%    398
 100%    817 (longest request)
```