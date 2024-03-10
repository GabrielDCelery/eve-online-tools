FROM ubuntu:focal-20240216

ARG USER_NAME=local-user
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Update system
RUN apt update \
    && apt install wget -y \
    && apt install curl -y

# Create user
RUN groupadd --gid $USER_GID $USER_NAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USER_NAME \
    && apt install sudo -y \
    && echo $USER_NAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USER_NAME \
    && chmod 0440 /etc/sudoers.d/$USER_NAME

# Switch to user
USER $USER_NAME

# Install Python
RUN sudo apt install python3 -y \
    && sudo apt install python3-pip -y

# Install pre-commit
RUN sudo pip3 install pre-commit

# Install Git
RUN sudo apt install git -y

# Install Go
ENV GO_VERSION="1.22.1"
RUN sudo wget https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz \
    && sudo rm -rf /usr/local/go \
    && sudo tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz
ENV PATH="${PATH}:/usr/local/go/bin"
ENV GOPATH="/home/$USER_NAME/go"
ENV PATH="${PATH}:${GOPATH}/bin"
RUN go version


# Install pre-commit dependencies
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin

WORKDIR /home/$USER_NAME/eve-online-tools-cli

ENTRYPOINT ["tail", "-f", "/dev/null"]
