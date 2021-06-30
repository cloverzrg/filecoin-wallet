FROM ubuntu:20.04
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get --allow-insecure-repositories update
RUN apt-get install -y ca-certificates tzdata hwloc libhwloc-dev mesa-opencl-icd ocl-icd-opencl-dev libc6-dev libc-dev
ADD go-bin /app/fil-wallet
EXPOSE 80
ENV ENDPOINT https://calibration.node.glif.io
ENV TOKEN ''
VOLUME /app/data
CMD ["/app/fil-wallet"]