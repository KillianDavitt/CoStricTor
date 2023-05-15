#### submission

| filter size (bits) | filter size (bytes) |                                 |                                  | time to compress |                 | submission size (bytes) (100000 samples) |                                    |        |                                     |        | time to compress (100000 samples) |                 |
|--------------------|---------------------|---------------------------------|----------------------------------|------------------|-----------------|------------------------------------------|------------------------------------|--------|-------------------------------------|--------|-----------------------------------|-----------------|
|                    | uncompressed bin    | compressed bin (brotli level 4) | compressed bin (brotli level 11) | brotli level 4   | brotli level 11 | uncompressed binary                      | compressed binary (brotli level 4) |        | compressed binary (brotli level 11) |        | brotli level 4                    | brotli level 11 |
|                    |                     |                                 |                                  |                  |                 |                                          | mean                               | std    | mean                                | std    |                                   |                 |
| 20000              | 80180               | 28609                           | 23076                            | <0.01s           | 0.07s           | 2500                                     | 1412.21                            | 74.42  | 1433.46                             | 158.23 | 142.86s                           | 329.20s         |
| 40000              | 160180              | 56519                           | 45038                            | <0.01s           | 0.22s           | 5000                                     | 2784.55                            | 103.11 | 2839.12                             | 353.46 | 270.73s                           | 610.52s         |

Higher compression level can be used on server as it can be computed rarely and cached (or fixed for read-only filter).
Lower compression level should be used on client to not tax computational resources? 

TODO:
- 4 x filter
