FROM ghcr.io/charmbracelet/vhs:v0.9.0

RUN apt-get -y install --no-install-recommends nvim

ENTRYPOINT ["/usr/bin/vhs"]
