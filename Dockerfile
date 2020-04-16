FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o service

FROM alpine
WORKDIR /app
COPY --from=build-env /src/service /app/
CMD ./service