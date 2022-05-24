FROM alpine:3.16.0

ARG component

COPY build/$component .

CMD ["/${component}"]
