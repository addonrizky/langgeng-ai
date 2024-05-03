# BUILDER GOLANG APPS
FROM golang:1.20-bullseye as builder
WORKDIR /app9

COPY go.mod ./
COPY main.go ./
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


# BUILDER PYTHON LANGCHAIN
FROM debian:bookworm

RUN apt update -qq \
    && apt install -qq -y --no-install-recommends \
      wget \
      zlib1g-dev \
      libncurses5-dev \
      libgdbm-dev \ 
      libnss3-dev \
      libssl-dev \
      libreadline-dev \
      libffi-dev \
      libsqlite3-dev \
      libbz2-dev \
      python3-pip \
      curl \
      git \
      gnupg \
      libgconf-2-4 \
      libxss1 \
      libxtst6 \
      g++ \
      build-essential \
      python3.11-venv \ 
      vim

RUN mkdir -p $HOME/.venvs \
    && python3 -m venv $HOME/.venvs/MyEnv \
    && $HOME/.venvs/MyEnv/bin/python -m pip install chromadb \
    && $HOME/.venvs/MyEnv/bin/python -m pip install openai \
    && $HOME/.venvs/MyEnv/bin/python -m pip install langchain \
    && $HOME/.venvs/MyEnv/bin/python -m pip install tiktoken

RUN echo "alias python=$HOME/.venvs/MyEnv/bin/python" >> ~/.bashrc

WORKDIR /app9

COPY . .

COPY --from=builder /app9/main .