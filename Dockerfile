FROM dqneo/ubuntu-build-essential:go

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
RUN apt-get update && apt install -y automake
RUN cd ctags
RUN apt-get install -y libjansson-dev && cd ctags && ./autogen.sh && ./configure && make && make install
# RUN ./autogen.sh
# RUN ./configure
# RUN make
# RUN make install

RUN git clone https://github.com/go-telegram-bot-api/telegram-bot-api
RUN cd telegram-bot-api/
RUN ctags -R  --output-format=json --languages=go .

CMD ["/bin/bash"]
