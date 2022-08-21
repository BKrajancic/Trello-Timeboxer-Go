FROM golang

ARG workdir=/TrelloTimeboxerGo
ENV binary_filepath ${workdir}/app
ADD ./ ${workdir}
WORKDIR ${workdir}
RUN go build -o ${binary_filepath} ${workdir}
WORKDIR ${workdir}/src

CMD ${binary_filepath}