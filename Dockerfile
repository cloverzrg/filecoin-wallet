FROM ubuntu:18.04
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get --allow-insecure-repositories update
RUN apt-get install -y ca-certificates tzdata hwloc libhwloc-dev mesa-opencl-icd ocl-icd-opencl-dev libc6-dev libc-dev
RUN apt-get install -y libc6-dev libc-dev make mesa-opencl-icd ocl-icd-opencl-dev gcc git bzr jq pkg-config curl clang build-essential hwloc libhwloc-dev wget
ADD go-bin /app/fil-wallet
ADD templates /app/templates
EXPOSE 80
ENV ENDPOINT https://calibration.node.glif.io
ENV TOKEN ''
VOLUME /app/data
CMD ["/app/fil-wallet"]