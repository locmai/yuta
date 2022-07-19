FROM alpine:3.16.1

ARG component

COPY build/$component .

CMD ["/${component}"]
