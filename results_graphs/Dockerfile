# syntax=docker/dockerfile:1

FROM python:3.9

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

WORKDIR /code/changing_filter_size/
# Run
CMD [ "bash", "-c", "python3 generate_graph.py && cd ../p_versus_q && python3 generate_graph.py && cd ../best_sites && python3 graph_optimal_filter_size.py"]
