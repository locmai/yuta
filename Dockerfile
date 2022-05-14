FROM alpine:3.15.4

ARG component

COPY build/$component .

CMD ["/${component}"]
