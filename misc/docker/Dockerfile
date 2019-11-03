FROM ubuntu 

CMD /bin/bash

RUN apt-get update
RUN apt-get install -y git gcc curl build-essential libssl-dev libreadline-dev zlib1g-dev sudo
# ubuntu only
RUN apt-get install -y libpq-dev

# install rbenv/ruby
RUN curl -fsSL https://github.com/rbenv/rbenv-installer/raw/master/bin/rbenv-installer | bash; exit 0
ENV PATH="/root/.rbenv/bin/:${PATH}"
RUN echo 'eval "$(rbenv init -)"' >> /root/.bashrc
RUN rbenv install 2.6.5

# install postgresql
RUN ln -fs /usr/share/zoneinfo/America/New_York /etc/localtime
RUN apt-get install -y postgresql postgresql-client
RUN dpkg-reconfigure --frontend noninteractive tzdata
RUN mkdir -p /usr/local/pgsql/data
RUN chown postgres /usr/local/pgsql/data

WORKDIR /

COPY ./bootstrap.sh ./
