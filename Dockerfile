# Use an official Bazel image as a base
#FROM gcr.io/bazel-public/bazel:7.3.1
FROM ubuntu:24.04

# Install wget and other packages required by bazel
RUN apt-get update && apt-get install -y wget g++ unzip zip apt-transport-https curl gnupg
USER root
RUN curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor >bazel-archive-keyring.gpg
RUN mv bazel-archive-keyring.gpg /usr/share/keyrings
RUN echo "deb [arch=amd64 signed-by=/usr/share/keyrings/bazel-archive-keyring.gpg] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list
RUN apt-get install -y bazel

# Install bazelisk, and then bazel
#RUN wget https://github.com/bazelbuild/bazelisk/releases/download/v1.20.0/bazelisk-linux-amd64 && \
#    chmod 755 bazelisk-linux-amd64 && \
#    mv bazelisk-linux-amd64 /usr/bin/bazelisk

# Install Golang
RUN wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz  && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz

# Set the working directory inside the container
WORKDIR /workspace

# Copy your project files into the container
COPY . /workspace

# Set environment variables, if necessary
ENV GO111MODULE=on

# Run Bazel tests
CMD bazel test //...