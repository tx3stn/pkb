FROM ghcr.io/charmbracelet/vhs:v0.10.0

RUN apt-get -y install --no-install-recommends neovim xclip

ENTRYPOINT ["/usr/bin/vhs"]
