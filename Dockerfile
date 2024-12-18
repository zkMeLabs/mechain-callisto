FROM golang:1.20-bookworm AS builder

WORKDIR /go/src/github.com/forbole/bdjuno

ENV CGO_CFLAGS="-O -D__BLST_PORTABLE__"
ENV CGO_CFLAGS_ALLOW="-O -D__BLST_PORTABLE__"

COPY . .

RUN make build
ADD https://github.com/hasura/graphql-engine/releases/download/v2.43.0/cli-hasura-linux-amd64 ./build/hasura
RUN chmod +x ./build/hasura


FROM golang:1.20-bookworm

WORKDIR /root

COPY --from=builder /go/src/github.com/forbole/bdjuno/build/hasura /usr/bin/hasura
COPY --from=builder /go/src/github.com/forbole/bdjuno/hasura ./hasura
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/bin/bdjuno

CMD [ "bdjuno" ]