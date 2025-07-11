FROM ubuntu:latest

WORKDIR /app

COPY web /app/web
COPY scheduler /app/scheduler

ENV TODO_DBFILE=/data/scheduler.db
ENV TODO_PORT=7540
ENV TODO_PASSWORD=1234 

CMD ["./scheduler"]