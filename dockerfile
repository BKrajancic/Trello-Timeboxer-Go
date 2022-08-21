FROM golang

ARG workdir=/TrelloTimeboxerGo
ADD ./ ${workdir}
RUN go install ${workdir}
RUN go build -x ${workdir}/app.go

WORKDIR ${workdir}/src
CMD ${workdir}/app
