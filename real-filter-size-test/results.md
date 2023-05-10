#### submission

| filter size (bits) | filter size (bytes) |                                     | submission size (bytes) (100000 samples) |                                    |         |
|--------------------|---------------------|-------------------------------------|------------------------------------------|------------------------------------|---------|
|                    | uncompressed json   | compressed json (brotli level 4/11) | uncompressed binary                      | compressed binary (brotli level 4) ||
|                    |                     |                                     | mean                                     | std                                ||
| 20000              | 100010              | 32577/26936                         | 2500                                     | 1412.21465                         | 74.422  |
| 40000              | 200016              | 65046/53467                         | 5000                                     | 2818.77                            | 141.415 |

Higher compression level can be used on server as it can be computed rarely and cached (or fixed for read-only filter).
Lower compression level should be used on client to not tax computational resources.