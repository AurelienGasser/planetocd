#!/bin/sh

set -e

gem install bundler
sed -i "s/gem 'tzinfo-data', platforms: \[:mingw, :mswin, :x64_mingw, :jruby\]/gem 'tzinfo-data'/" ./Gemfile
rm Gemfile.lock
bundle install
bundle update tzinfo
bundle exec rails server