FROM node:17-alpine AS frontend

WORKDIR /fe

ADD frontend ./

ENV PATH /app/node_modules/.bin:$PATH
RUN npm ci --silent
# RUN npm install react-scripts@3.4.1 -g --silent

RUN npm run build

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

ADD src ./src
COPY Wep_App_Sec.go ./

COPY --from=frontend /fe/build ./build
RUN go build -o /Web_App_Sec


EXPOSE 8080

CMD ["/Web_App_Sec"]