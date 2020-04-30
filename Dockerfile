FROM python:3.6-alpine

RUN apk update && \
    apk add gcc musl-dev python3-dev libffi-dev openssl-dev

# Flask 0.12.2 has a couple CVEs in it
RUN pip install flask
#==0.12.2

RUN addgroup --gid 2323 "flasky" && \
    adduser --disabled-password \
            --home "/home/flasky" \
            --ingroup "flasky" \
            --no-create-home \
            --uid 2324 \
            "flasky"

USER flasky
WORKDIR /home/flasky

ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8
ENV FLASK_APP=/app/app.py

WORKDIR /app
COPY . .

# EXPOSE 22

ENTRYPOINT ["flask", "run", "--host", "0.0.0.0"]
