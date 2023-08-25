# syntax=docker/dockerfile:1

FROM python:latest

#WORKDIR /app
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY . .

# Run
CMD [ "python3", "changing_filter_size/crews.py"]
