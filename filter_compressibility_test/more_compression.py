import lzma
import zlib
import bz2
import lz4framed
import snappy

from app import data20k, data20k_alt, data40k, data40k_alt

LZMA = "LZMA"
ZLIB = "ZLIB"
BZ2 = "BZ2"
LZ4 = "LZ4"
SNAPPY = "SNAPPY"

COMPRESSION = [LZMA, ZLIB, BZ2, LZ4, SNAPPY]

COMPRESSION_MAP = {
    LZMA: lzma.compress,
    ZLIB: zlib.compress,
    BZ2: bz2.compress,
    LZ4: lz4framed.compress,
    SNAPPY: snappy.compress,
}


def compression(size=40):
    if size == 20:
        data = data20k + data20k_alt
    else:
        data = data40k_alt + data40k
    data_bytes = map(lambda x: x.to_bytes(4, byteorder='little'), data)
    data_joined = b''.join(data_bytes)

    for comp in COMPRESSION:
        print(comp, len(data_joined), len(COMPRESSION_MAP[comp](data_joined)))


if __name__ == '__main__':
    compression(20)
