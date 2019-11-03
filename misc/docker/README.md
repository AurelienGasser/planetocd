Run the following commands:

```bash
$ docker build -t planetocd .
$ docker run -it planetocd /bin/bash
root$ ./bootstrap.sh
root$ cd planetocd
root$ bundle exec rails server &
root$ curl localhost:4242 # tada!
```
