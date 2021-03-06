FROM alpine:latest

LABEL app="angago"
LABEL maintainer="manigandan.jeff@gmail.com"
LABEL version="1.0.3"
LABEL description="angago localhost proxy tunnel."

RUN mkdir -p /app && apk update && apk add --no-cache ca-certificates
WORKDIR /app
# This require the project to be built first before copying,
# else docker build will fail
COPY angago /app/
COPY stub /app/stub

# Load ENV
ENV ENV prod
ENV PORT 80
ENV ANGAGO_CONFIG_PATH "/root/.angago/config.yaml"

EXPOSE 80

# include the conf file location on load?
CMD /app/angago
