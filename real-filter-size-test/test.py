import pytest
from statistics import mean, stdev

from app import app


@pytest.fixture()
def client():
    return app.test_client()


def test_request_example(client):
    headers = [('Accept-Encoding', 'br')]
    lengths = []
    for size in [20, 40]:
        for i in range(100000):
            response = client.get(f"/compressedsubmit?size={size}", headers=headers)
            lengths.append(response.calculate_content_length())
        print()
        print(size)
        print(mean(lengths))
        print(stdev(lengths))
