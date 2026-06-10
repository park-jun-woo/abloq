# abloqd image — builds the yongol projection in backend/arts/backend.
#
# Build context is the abloq repository root (see docker-compose.yaml):
# backend/arts/ is a disposable projection, so run `yongol generate
# backend/specs backend/arts` before building. pkg/content is consumed via a
# local replace because github.com/park-jun-woo/abloq@v0.0.1 is not published
# yet — the replace happens only inside the image, never in the working tree.

FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
COPY pkg ./pkg
COPY backend/arts/backend ./backend/arts/backend
WORKDIR /src/backend/arts/backend
RUN go mod edit -replace github.com/park-jun-woo/abloq=/src \
    && CGO_ENABLED=0 go build -o /out/abloqd ./cmd

FROM alpine:3.21
# addgroup/adduser are busybox builtins; CA certs come from the builder image.
# git + openssh-client are required at runtime by the queue exporter
# (Phase009): pkg/queueio shells out to git for the work-clone pull → commit →
# deploy-key push cycle — without them Hurl (host-run) passes but production
# dies on the first POST /queue/export.
RUN apk add --no-cache git openssh-client \
    && addgroup -S abloqd && adduser -S abloqd -G abloqd \
    # exporter work clone home — the only writable path abloqd needs. The
    # directory exists in the image so the compose named volume inherits
    # abloqd ownership on first use.
    && mkdir -p /var/lib/abloqd/queue-export && chown -R abloqd:abloqd /var/lib/abloqd
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /out/abloqd /usr/local/bin/abloqd
COPY backend/arts/backend/policy /etc/abloqd/policy
USER abloqd
EXPOSE 8080
ENV OPA_POLICY_PATH=/etc/abloqd/policy/authz.rego
ENV QUEUE_EXPORT_WORKDIR=/var/lib/abloqd/queue-export
ENTRYPOINT ["/usr/local/bin/abloqd"]
