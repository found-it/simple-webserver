FROM python:3.6-alpine

RUN apk update && \
    apk add gcc musl-dev python3-dev libffi-dev openssl-dev

# Flask 0.12.2 has a couple CVEs in it
RUN pip install flask
#==0.12.2

ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8
ENV FLASK_APP=/app/app.py

WORKDIR /app
COPY . .

ENTRYPOINT ["flask", "run", "--host", "0.0.0.0"]
