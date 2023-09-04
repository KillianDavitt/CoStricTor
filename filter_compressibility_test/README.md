Tiny application to test size of submission and download for CoStricTor.

Provides a web server with a few endpoints:
- `/compressed` and `/uncompressed` to download the shared filters.
  - take a `size` parameter to specify the size of the filter to download
- `/uncompressedsubmit` and `/compressedsubmit` to download a random submission (to replicate the submission process)
  - take a `size` parameter to specify the size of the filter to download
  - take `p` and `q` parameters -- Rappor parameters
  - take `h` parameter -- number of hash functions
  - take `t` parameter -- which specifies the encoding: if `bit` then the submission is encoded as a bit vector, otherwise it is a json list.

Tests
-----
Test methods are provided to more easily determine size post-compression and the time required to compress.

Run with:
```bash
 pytest test.py::<test_function_name> --durations 0 -vv -s
```