Run Planet OCD on your machine
------------------------------

In order to run Planet OCD locally, you'll need:

- ruby version 2.6.5
- a postgresql server

### Install ruby 2.6.5

The easiest way is to use [rbenv](https://github.com/rbenv/rbenv#installation). 

After installing `rbenv`, you can install ruby v2.6.5 with:

```bash
$ eval "$(rbenv init -)" # you'll need to run this for each new shell,
                         # or add it to your ~/.bashrc
$ rbenv install 2.6.5
```

Check that everything is fine:

```bash
$ cd planetocd
$ ruby -v
ruby 2.6.5p114 (2019-10-01 revision 67812) [x86_64-linux]
```

### Connect to a postgresql server

#### Option 1: Elephant SQL

An easy way to get a postgres server running is to use [Elephant SQL](elephantsql.com). 
It only takes a minute to sign-up and get a free instance running.

Once you have created an instance, the instance URL should be visible from the Elephant SQL UI (e.g. `postgres://bvbrywcs:yWUhJbxjQSB1rEtcbyfCWpAgnAgKSwf6@salt.db.elephantsql.com:5432/bvbrywcs`)

Simply run:

```
$ export DATABASE_URL="<your_instance_url>"
```

#### Option 2: local postgresql server

If you already have a postgresql server running, you can simply create a new `planetocd` user/password:

```
$ su - postgres
postgres$ psql -c "CREATE USER planetocd with encrypted password 'planetocd';"
postgres$ psql -c "ALTER USER planetocd CREATEDB;"
```

### Run Planet OCD

Once you have ruby v2.6.5 and a postgresql ready, you can bootstrap the app with:

```bash
$ gem install bundler:2.0.2
$ bundle install --path vendor/bundle --binstubs vendor/bundle/bin -j4 --deployment
$ rake assets:precompile
```

And then start the server:

```bash
$ bundle exec rails server
```

Your local instance of Planet OCD should now be running on http://localhost:4242/