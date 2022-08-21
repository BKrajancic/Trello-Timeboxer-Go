FROM golang

ARG workdir=/TrelloTimeboxerGo
ADD ./ ${workdir}
WORKDIR ${workdir}
RUN go build -o app ${workdir}
WORKDIR ${workdir}/src

CMD ${workdir}/app