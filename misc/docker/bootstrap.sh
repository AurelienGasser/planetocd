#!/bin/bash

set -e

# postgres stuff
sudo -H -u postgres bash -c "/usr/lib/postgresql/10/bin/pg_ctl -D /etc/postgresql/10/main -l /var/lib/postgresql/logfile start"
sudo -H -u postgres bash -c $'psql -c "create user planetocd with encrypted password \'planetocd\';"'
sudo -H -u postgres bash -c $'psql -c "ALTER USER planetocd CREATEDB;"'

# clone the repo
git clone https://github.com/AurelienGasser/planetocd.git
cd planetocd

# rails stuff
gem install bundler:2.0.2
bundle install --path vendor/bundle --binstubs vendor/bundle/bin -j4 --deployment
bundle exec rake db:create
bundle exec rake assets:precompile
