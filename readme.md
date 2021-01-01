
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
Time taken for tests:   615.684 seconds
Complete requests:      100000
Failed requests:        0
Keep-Alive requests:    100000
Total transferred:      19000000 bytes
HTML transferred:       5800000 bytes
Requests per second:    162.42 [#/sec] (mean)
Time per request:       184.705 [ms] (mean)
Time per request:       6.157 [ms] (mean, across all concurrent requests)
Transfer rate:          30.14 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       2
Processing:     9  185  64.1    191     696
Waiting:        9  185  64.1    191     696
Total:          9  185  64.1    191     696

Percentage of the requests served within a certain time (ms)
  50%    191
  66%    200
  75%    207
  80%    213
  90%    283
  95%    297
  98%    313
  99%    379
 100%    696 (longest request)

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
Time taken for tests:   709.041 seconds
Complete requests:      100000
Failed requests:        0
Keep-Alive requests:    100000
Total transferred:      20500000 bytes
Total body sent:        22400000
HTML transferred:       7300000 bytes
Requests per second:    141.04 [#/sec] (mean)
Time per request:       212.712 [ms] (mean)
Time per request:       7.090 [ms] (mean, across all concurrent requests)
Transfer rate:          28.23 [Kbytes/sec] received
                        30.85 kb/s sent
                        59.09 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       2
Processing:    14  213  67.8    202    1291
Waiting:       13  213  67.8    202    1291
Total:         15  213  67.8    202    1291

Percentage of the requests served within a certain time (ms)
  50%    202
  66%    213
  75%    272
  80%    282
  90%    297
  95%    309
  98%    383
  99%    395
 100%   1291 (longest request)

```