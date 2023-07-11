pogreb-bench
============
pogreb-bench is a key-value store benchmarking tool. Currently it supports pogreb, goleveldb, bolt and badger and rosedb.
==============================================
engine: pogreb
keys: 100000
key size: 16-64
value size 128-512
concurrency: 1

put: 0.722s     138546 ops/s
get: 0.052s     1919173 ops/s

put + get: 0.774s
file size: 105.44MB
peak sys mem: 35.19MB
==============================================
engine: bbolt
keys: 100000
key size: 16-64
value size 128-512
concurrency: 1

put: 2.552s     39192 ops/s
get: 0.159s     630411 ops/s

put + get: 2.710s
file size: 64.05MB
peak sys mem: 35.63MB
==============================================
engine: goleveldb
keys: 100000
key size: 16-64
value size 128-512
concurrency: 1

put: 0.447s     223657 ops/s
get: 0.397s     251893 ops/s

put + get: 0.844s
file size: 99.68MB
peak sys mem: 60.31MB
==============================================
badger 2023/07/11 13:10:38 INFO: All 0 tables opened in 0s
badger 2023/07/11 13:10:39 INFO: Discard stats nextEmptySlot: 0
badger 2023/07/11 13:10:39 INFO: Set nextTxnTs to 0
engine: badger
keys: 100000
key size: 16-64
value size 128-512
concurrency: 1

put: 0.753s     132880 ops/s
badger 2023/07/11 13:10:40 INFO: Lifetime L0 stalled for: 0s
badger 2023/07/11 13:10:40 INFO: 
Level 0 [ ]: NumTables: 01. Size: 12 MiB of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
badger 2023/07/11 13:10:40 INFO: All 1 tables opened in 2ms
badger 2023/07/11 13:10:40 INFO: Discard stats nextEmptySlot: 0
badger 2023/07/11 13:10:40 INFO: Set nextTxnTs to 100000
get: 0.317s     315521 ops/s

put + get: 1.069s
badger 2023/07/11 13:10:41 INFO: Lifetime L0 stalled for: 0s
badger 2023/07/11 13:10:41 INFO: 
Level 0 [ ]: NumTables: 01. Size: 12 MiB of 0 B. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 StaleData: 0 B Target FileSize: 2.0 MiB
Level Done
file size: 113.11MB
peak sys mem: 406.32MB
==============================================
engine: rosedb
keys: 100000
key size: 16-64
value size 128-512
concurrency: 1

put: 0.720s     138867 ops/s
get: 0.153s     653056 ops/s

put + get: 0.873s
file size: 113.11MB
peak sys mem: 180.99MB