FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates tzdata
ADD wallet /app/wallet
EXPOSE 80
ENV ENDPOINT https://calibration.node.glif.io
VOLUME /app/data
CMD ["/app/wallet"]