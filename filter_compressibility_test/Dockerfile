# syntax=docker/dockerfile:1

FROM python:latest

#WORKDIR /app
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY . .

# Run
CMD [ "flask", "run"]
CMD ["pytest","test.py", "-vv", "-s"]
