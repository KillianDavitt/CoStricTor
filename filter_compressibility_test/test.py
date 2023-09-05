import pytest
from statistics import mean, stdev

from app import app


@pytest.mark.parametrize("compress_level", [4, 11])
@pytest.mark.parametrize("size", [20, 40])
def test_submission(compress_level, size):
    app.config["COMPRESS_BR_LEVEL"] = compress_level
    client = app.test_client()
    headers = [('Accept-Encoding', 'br')]
    lengths = []
    for i in range(1000):
        response = client.get(f"/compressedsubmit?size={size}", headers=headers)
        lengths.append(response.calculate_content_length())

    print()
    print("compression level, size, mean content length, stdev content length")
    print(compress_level, size, mean(lengths), stdev(lengths))


@pytest.mark.parametrize("compress_level", [4, 11])
@pytest.mark.parametrize("size", [20, 40])
def test_filters(compress_level, size):
    app.config["COMPRESS_BR_LEVEL"] = compress_level
    # app.config["COMPRESS_BR_WINDOW"] = 24
    # app.config["COMPRESS_BR_BLOCK"] = 24
    client = app.test_client()
    headers = [('Accept-Encoding', 'br')]
    response = client.get(f"/compressed?size={size}", headers=headers)
    print()
    print("compression level, size, content length")
    print(compress_level, size, response.calculate_content_length())
