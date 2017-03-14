FROM golang
WORKDIR ${GOPATH}/src/github.com/tily/sdk-go
RUN go get -u github.com/golang/dep/...
ADD Makefile ${GOPATH}/src/github.com/tily/sdk-go/Makefile
RUN make get-deps-tests && make get-deps-verify
RUN apt-get update && \
    apt-get install -y vim && \
    mkdir -p ~/.vim/autoload ~/.vim/bundle && \
    curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim && \
    echo 'execute pathogen#infect()' >> ~/.vimrc && \
    echo 'syntax on' >> ~/.vimrc && \
    echo 'filetype plugin indent on' >> ~/.vimrc && \
    git clone https://github.com/fatih/vim-go.git ~/.vim/bundle/vim-go
RUN echo 'set encoding=utf-8' >> ~/.vimrc && \
    echo 'set fileencodings=iso-2022-jp,euc-jp,sjis,utf-8' >> ~/.vimrc && \
    echo 'set fileformats=unix,dos,mac' >> ~/.vimrc
