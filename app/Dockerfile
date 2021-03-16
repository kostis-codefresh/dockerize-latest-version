FROM alpine:3.12
ARG FILENAME
COPY $FILENAME /app/$FILENAME


RUN apk add --no-cache jq ca-certificates bash

WORKDIR /app
RUN chmod +x $FILENAME
RUN ln -s $FILENAME dockerized-app



CMD ["./dockerized-app"]
