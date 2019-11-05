# wee-go-image
The smallest go docker image on the planet! From [this](https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324) blog post.

# usage
```bash
# build
$ docker build -t wee-go-image:v0.1.0 .
# run
$ docker run --rm -p 8989:8989 --name wee-go wee-go-image:v0.1.0
```

~5.4MB - not bad! Granted we did not include any 3rd party dependencies but we do have a working http server.

# license
MIT
