# report

Отчет по заданию "Взаимодействие с помощью асинхронного обмена сообщениями (messaging)"

| конфигурация / средняя задержка (latency avg) | RPS-100                 | RPS-500 | RPS-1000 | RPS-1500 | RPS-2000 | Что изменялось? Почему? |
| --------------------------------------------- | ----------------------- | ------- | -------- | -------- | -------- | ----------------------- |
| RabbitMQ + 3 api + 1 processor                | TODO: заполнить таблицу |         |          |          |          |                         |
| RabbitMQ + 1 api + 1 processor                |                         |         |          |          |          |                         |
| Nats + 3 api + 1 processor                    |                         |         |          |          |          |                         |
| ...ваши варианты,,,                           |                         |         |          |          |          |                         |

Выводы: TODO: написать выводы

## Подсказки

### Где смотреть результаты теста

- Ожидаемый RPS - который передается в параметрах - см скрипт в makefile-e `for rps in 100 500 1000 1500 2000`  
- Факический RPS - который в отчете wrk, строка `Requests/sec:    100.56`
- Средняя задержка - в отчете wrk, ячейка в таблице `Latency/Avg     4.01ms`
- Важно что б не было записи `Socket errors: connect 0, read 0, write 0, timeout 160`, она говорит о том, что были проблемы при запросах

```bash
Running 15s test @ http://proxy:8080
  4 threads and 4 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.01ms    3.11ms  37.98ms   94.37%
    Req/Sec       -nan      -nan   0.00      0.00%
  604 requests in 6.01s, 71.37KB read
Requests/sec:    100.56
Transfer/sec:     11.88KB
```

### Если фактический RPS сильно отличается от ожидаемого

Значит сильная деградация сервиса и задержка будет гораздо больше 1секунды.  

### а что если - We are about to crash and die (recoridng a negative #)This wil never ever ever happen

У вас не хватает ресурсов или специфичные настройки операционной системы.  
Cделайте шаги RPS поменьше, допустим 50 100 200 300 500

```bash
Running 15s test @ http://proxy:8080
  4 threads and 4 connections


 ---------- 

We are about to crash and die (recoridng a negative #)This wil never ever ever happen...But when it does. The following information will help in debuggingresponse_complete:
  expected_latency_timing = -16615
  now = 1711906673977747
  expected_latency_start = 1711906673994362
  c->thread_start = 1711906673214362
  c->complete = 79
  throughput = 0.0001
  latest_should_send_time = 1711906673994534
  latest_expected_start = 1711906673994527
  latest_connect = 1711906673214513
  latest_write = 1711906673994534
  next expected_latency_start = 1711906674004362
Assertion failed: bucket_index < h->bucket_count (src/hdr_histogram.c: counts_index: 54)
```
