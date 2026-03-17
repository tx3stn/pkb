FROM ghcr.io/charmbracelet/vhs:v0.11.0

RUN apt-get -y install --no-install-recommends neovim xclip

ENTRYPOINT ["/usr/bin/vhs"]
