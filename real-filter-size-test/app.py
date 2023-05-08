import random
from bitarray import bitarray
from bitarray.util import serialize

from flask import Flask, request, jsonify

from flask_compress import Compress

app = Flask(__name__)
app.config["COMPRESS_REGISTER"] = False  # disable default compression of all eligible requests
compress = Compress()
compress.init_app(app)


def read_parse(filename):
    with open(filename, "r") as f:
        data = [int(x) for x in f.readline()[1:-2].split(" ")]
        return data


data20k = read_parse("crews_output_1.csv")
data40k = read_parse("crews_output_2.csv")


@app.route("/uncompressed")
def uncompressed():
    size = request.args.get("size", default=20, type=int)
    data = data20k if size == 20 else data40k
    return jsonify([data])


@app.route("/compressed")
@compress.compressed()
def compressed():
    return uncompressed()


@app.route("/uncompressedsubmit")
def submituncompressed():
    size = request.args.get("size", default=20, type=int)
    p = request.args.get("p", default=0.2, type=float)
    q = request.args.get("q", default=0.9, type=float)
    h = request.args.get("h", default=2, type=int)
    t = request.args.get("type", default="bit", type=str)
    data = bitarray(size * 1000) if t == "bit" else [0] * (size * 1000)

    for i in range(h):
        data[random.randint(0, size * 1000)] = 1

    for i in range(size * 1000):
        if data[i]:
            data[i] = 1 if random.random() < q else 0
        else:
            data[i] = 0 if random.random() < p else 1
    return serialize(data) if t == "bit" else jsonify(data)


@app.route("/compressedsubmit")
@compress.compressed()
def submitcompressed():
    return submituncompressed()
