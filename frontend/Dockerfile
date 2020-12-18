FROM ubuntu

RUN apt update && \
    apt install -y curl && \
    apt install -y gnupg2

RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
    echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list && \
    apt update && \
    apt install -y yarn

CMD cd /usr/src/frontend && yarn && yarn serve