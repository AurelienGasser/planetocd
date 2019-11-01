FROM ubuntu 

CMD /bin/bash

RUN apt-get update
RUN apt-get install -y git gcc curl build-essential libssl-dev libreadline-dev zlib1g-dev 
# ubuntu only
RUN apt-get install -y libpq-dev
RUN curl -fsSL https://github.com/rbenv/rbenv-installer/raw/master/bin/rbenv-installer | bash; exit 0

ENV PATH="/root/.rbenv/bin/:${PATH}"

RUN rbenv install 2.6.5
RUN git clone https://github.com/AurelienGasser/planetocd.git

WORKDIR /planetocd

RUN echo 'eval "$(rbenv init -)"' >> /root/.bashrc
COPY ./make-ready.sh .
