FROM golang

ARG workdir=/TrelloTimeboxerGo
ADD ./ ${workdir}
CMD go install ${workdir}
CMD go build -x ${workdir}/app.go

workdir ${workdir}/src
RUN ${workdir}/app
