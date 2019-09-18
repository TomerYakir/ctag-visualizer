# FROM dqneo/ubuntu-build-essential:go
FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y wget git gcc

RUN wget -P /tmp https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf /tmp/go1.11.5.linux-amd64.tar.gz
RUN rm /tmp/go1.11.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"


# Default Packages
RUN buildDeps=' \
          bison \
          libgdbm-dev \
          ruby \
          wget \
          build-essential \
          xz-utils \
          curl \
          autoconf \
          zlib1g-dev \
          libcurl4-openssl-dev \
          libsasl2-dev \
          libgmp3-dev \
          libpq-dev \
          openssh-client \
          git \
          git-core \
          libsnappy-dev \
          libreadline-dev \
          vim \
    ' \
    && apt-get update \
    && apt-get install -y --no-install-recommends $buildDeps \
    && rm -rf /var/lib/apt/lists/*

# Install universal ctags
RUN git clone https://github.com/universal-ctags/ctags.git
RUN cd ctags
RUN apt-get update && apt install -y automake pkg-config
RUN apt-get install -y libjansson-dev && cd ctags && ./autogen.sh && ./configure && make && make install

RUN git clone https://github.com/go-telegram-bot-api/telegram-bot-api
# RUN cd telegram-bot-api/
# RUN ctags -R  --output-format=json --languages=go .

WORKDIR /src

# Import the code from the context.
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app server.go bindata.go

CMD ["/bin/bash"]

# # Final stage: the running container.
# FROM scratch AS final

# # Import the compiled executable from the first stage.
# COPY --from=builder /app /app

# # Declare the port on which the webserver will be exposed.
# # As we're going to run the executable as an unprivileged user, we can't bind
# # to ports below 1024.
# EXPOSE 8080

# # Run the compiled binary.
# ENTRYPOINT ["/app"]
