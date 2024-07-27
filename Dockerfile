FROM silkeh/clang:18

RUN apt update && apt-get install curl -y
ARG TARGETARCH
RUN curl -L https://dl.google.com/go/go1.22.5.linux-"$TARGETARCH".tar.gz | tar -xz -C /usr/local;
COPY . /project/candice
RUN mkdir /candice && cd /project/candice/cmd/build/ && /usr/local/go/bin/go build . && mv /project/candice/cmd/build/build /candice/candice
ENV PATH="$PATH:/candice"