FROM alpine:3.15.0

ARG component

COPY build/$component .

CMD ["/${component}"]
