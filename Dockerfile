FROM ubuntu:20.04
RUN apt-get update && apt-get install -y --no-cache ca-certificates tzdata
ADD go-bin /app/fil-wallet
EXPOSE 80
ENV ENDPOINT https://calibration.node.glif.io
VOLUME /app/data
CMD ["/app/fil-wallet"]