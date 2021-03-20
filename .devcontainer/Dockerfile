#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------
ARG VERSION="20.04"

FROM ubuntu:${VERSION}

ARG DEBIAN_FRONTEND=noninteractive
ARG TERRAFORM_VERSION=0.13.4
ARG TFLINT_VERSION=0.18.0
ARG GO_VERSION=1.15.1
ARG GOTESTSUM_VERSION=0.5.3
ARG GOLANGCILINT_VERSION=1.31.0

# Install common tools and packages
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    software-properties-common \
    gnupg \
    apt-transport-https \
    lsb-release \
    ca-certificates \
    jq \
    curl \
    git \
    unzip

RUN mkdir -p /tmp/docker-downloads \
    # Install Terraform, tflint
    && curl -sSL -o /tmp/docker-downloads/terraform.zip https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    && unzip /tmp/docker-downloads/terraform.zip \
    && mv terraform /usr/local/bin \
    && cd ~ \ 
    && rm -rf /tmp/docker-downloads \
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Make shared download folder
RUN mkdir -p /tmp/docker-downloads

# TFLint
RUN https://raw.githubusercontent.com/terraform-linters/tflint/master/install_linux.sh | bash

# Install Go into /usr/local/go
RUN curl -fsS https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz -o /tmp/docker-downloads/golang.tar.gz \
    && tar -C /usr/local -xvf /tmp/docker-downloads/golang.tar.gz 

# Install gotestsum into /usr/bin
RUN curl -fsSL https://github.com/gotestyourself/gotestsum/releases/download/v${GOTESTSUM_VERSION}/gotestsum_${GOTESTSUM_VERSION}_linux_amd64.tar.gz -o /tmp/docker-downloads/gotestsum.tar.gz \
    && tar -C /tmp/docker-downloads -xzf /tmp/docker-downloads/gotestsum.tar.gz \
    && mv /tmp/docker-downloads/gotestsum /usr/bin

# Install golangci-lint into /usr/bin
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/bin v${GOLANGCILINT_VERSION}

# Golang environmentals and path
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV PATH=$PATH:/usr/local/go/bin

# clean up the download folder
RUN rm -rf /tmp/docker-downloads

ARG DEBIAN_FRONTEND=interactive

# Set up a non root user called 'vscode'
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user
# Note. We don't switch to it, so the pipelines still run as root
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    # Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME