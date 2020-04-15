FROM ubuntu:18.04

RUN apt-get update && \
    apt-get install -y python3 python3-dev python3-pip nginx && \
    pip3 install uwsgi

WORKDIR /app
COPY . .

RUN pip3 install -r requirements.txt

ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8

ENTRYPOINT ["env", "FLASK_APP=/app/app.py", "flask", "run"]
