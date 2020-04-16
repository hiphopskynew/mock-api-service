FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o mock-api-service

FROM alpine
WORKDIR /services
COPY --from=build-env /src/mock-api-service /services/.
CMD ./mock-api-service