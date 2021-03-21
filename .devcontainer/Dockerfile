FROM mcr.microsoft.com/vscode/devcontainers/base:buster

ARG INSTALL_ZSH="true"
ARG UPGRADE_PACKAGES="true"
ARG ENABLE_NONROOT_DOCKER="true"
ARG USE_MOBY="true"

ARG USERNAME=automatic
ARG USER_UID=1000
ARG USER_GID=$USER_UID
COPY library-scripts/*.sh /tmp/library-scripts/
RUN apt-get update \
    && /bin/bash /tmp/library-scripts/common-debian.sh "${INSTALL_ZSH}" "${USERNAME}" "${USER_UID}" "${USER_GID}" "${UPGRADE_PACKAGES}" "true" "true" \
    && /bin/bash /tmp/library-scripts/docker-debian.sh "${ENABLE_NONROOT_DOCKER}" "/var/run/docker-host.sock" "/var/run/docker.sock" "${USERNAME}" "${USE_MOBY}" \
    && apt-get install -y make \
    && apt-get autoremove -y && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts/

ARG TERRAFORM_VERSION=0.14.8
ARG TFLINT_VERSION=0.25.0
ARG TERRAFORM_DOC_VERSION=0.11.2
ARG GO_VERSION=1.16.1
ARG GOTESTSUM_VERSION=1.6.2
ARG GOLANGCILINT_VERSION=1.38.0
ARG KUBECTL_VERSION=1.18.12

# Make shared download folder
RUN mkdir -p /tmp/downloads

# Install Terraform
RUN curl -sSL -o /tmp/downloads/terraform.zip https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    && unzip /tmp/downloads/terraform.zip \
    && mv terraform /usr/local/bin

# TFLint
RUN curl -L -o /tmp/downloads/tflint.zip https://github.com/terraform-linters/tflint/releases/download/v${TFLINT_VERSION}/tflint_$(uname)_amd64.zip \
    && unzip /tmp/downloads/tflint.zip \
    && mv tflint /usr/local/bin

# Install Terraform-docs
RUN curl -Lo /tmp/downloads/terraform-docs.tar.gz https://github.com/terraform-docs/terraform-docs/releases/download/v${TERRAFORM_DOC_VERSION}/terraform-docs-v${TERRAFORM_DOC_VERSION}-$(uname)-amd64.tar.gz \
    && tar -xzf /tmp/downloads/terraform-docs.tar.gz \
    && chmod +x terraform-docs \
    && mv terraform-docs /usr/local/bin

# Go
RUN curl -fsS https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz -o /tmp/downloads/golang.tar.gz \
    && tar -C /usr/local -xvf /tmp/downloads/golang.tar.gz 

# gotestsum
RUN curl -fsSL https://github.com/gotestyourself/gotestsum/releases/download/v${GOTESTSUM_VERSION}/gotestsum_${GOTESTSUM_VERSION}_linux_amd64.tar.gz -o /tmp/downloads/gotestsum.tar.gz \
    && tar -C /tmp/downloads -xzf /tmp/downloads/gotestsum.tar.gz \
    && mv /tmp/downloads/gotestsum /usr/bin

# golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/bin v${GOLANGCILINT_VERSION}

# Install k3d
RUN curl -s https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash

# Install kubectl
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl \
    && chmod +x ./kubectl \
    && mv ./kubectl /usr/local/bin

# Cleanup
RUN rm -rf /tmp/downloads    

# Golang environmentals and path
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV PATH=$PATH:/usr/local/go/bin

ENTRYPOINT [ "/usr/local/share/docker-init.sh" ]
CMD [ "sleep", "infinity" ]