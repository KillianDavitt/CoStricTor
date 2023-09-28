# syntax=docker/dockerfile:1

FROM python:3.9

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

WORKDIR /code/changing_filter_size/
# Run
CMD [ "python3", "generate_graph.py"]

WORKDIR /code/p_versus_q/
# Run
CMD [ "python3", "generate_graph.py"]

WORKDIR /code/best_sites/
# Run
CMD [ "python3", "graph_optimal_filter_size.py"]
