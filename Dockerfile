FROM ubuntu:20.04
RUN apt-get --allow-insecure-repositories update
RUN apt-get install -y --no-cache ca-certificates tzdata
ADD go-bin /app/fil-wallet
EXPOSE 80
ENV ENDPOINT https://calibration.node.glif.io
ENV TOKEN ''
VOLUME /app/data
CMD ["/app/fil-wallet"]