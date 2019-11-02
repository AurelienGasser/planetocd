Run locally
-----------

Thanks for contributing to Planet OCD!

To run the ruby-on-rails application locally, please follow these steps:

- https://github.com/rbenv/rbenv#installation

docker run -it ubuntu /bin/bash

apt-get install -y postgresql

adduser postgres
mkdir /usr/local/pgsql/data
chown postgres /usr/local/pgsql/data
su - postgres
# or pg_ctlcluster 11 main start

su - postgres
/usr/lib/postgresql/10/bin/pg_ctl -D /etc/postgresql/10/main -l logfile start
create user planetocd with encrypted password 'planetocd';
ALTER USER planetocd CREATEDB
exit

export PLANETOCD_DATABASE_PASSWORD='planetocd'
rake db:create

bundle exec rails server