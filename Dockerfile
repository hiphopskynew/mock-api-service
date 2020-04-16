FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o mock-api-service

FROM alpine
WORKDIR /services
COPY --from=build-env /src/mock-api-service /services/mock-api-service
COPY --from=build-env /src/config/config.yml /services/config/config.yml
CMD ./mock-api-service