FROM golang:1.20

WORKDIR /wrk

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

# install chromedriver
#RUN apt-get -y update
#RUN apt-get install -yqq unzip
#RUN wget -O /tmp/chromedriver.zip http://chromedriver.storage.googleapis.com/`curl -sS chromedriver.storage.googleapis.com/LATEST_RELEASE`/chromedriver_linux64.zip
#RUN unzip /tmp/chromedriver.zip chromedriver -d /usr/local/bin/

COPY ./driver/chromedriver /usr/local/bin/

RUN go build -o /parser

CMD [ "/parser" ]
