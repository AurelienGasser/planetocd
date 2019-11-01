Planet OCD
=======

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![Donate](https://img.shields.io/badge/Paypal-Donate-green.svg?logo=paypal&style=flat)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=23LG7JTZSCA54&source=url)

About
----------

[**Planet OCD**](https://www.planetocd.org/) is a website featuring articles about Obsessive Compulsive Disorder (OCD) in foreign languages. 

The goal of Planet OCD is to spread quality and up-to-date information to non-English speakers, for free. 

The articles come from external sources and are translated from their original English version. Currently, only articles translated into _French_ are available. But other languages can be introduced (maybe by you? see the [Contribution Guide](#contribute)).

Contribute
----------

> **Working on your first Pull Request?** You can learn how from this *free* series [How to Contribute to an Open Source Project on GitHub](https://egghead.io/series/how-to-contribute-to-an-open-source-project-on-github)

Your contributions are essential. See the sections below to see how you can help.

### Suggest new articles

You can suggest new articles that you would like to see on Planet OCD by sumitting a [new Issue](https://github.com/AurelienGasser/planetocd/issues).

### Correct an existing article

You can make corrections to existing articles by opening a Pull Request.

### Translate

Are you bilingual in English and a foreign language? You can help by translating new articles.

You can submit Pull Requests with new translations. See the [Translation Guide](./translation_guide.md) for more information.

### Donate

Your donation will help finance the costs related to the hosting server. I can also be used to hire professional translation services, and enrich Planet OCD with brand new articles. [Click here](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=23LG7JTZSCA54&source=url) to donate now!

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