FROM node:lts as base

WORKDIR /home/workspace

ADD . .

RUN bash -c "yarn install"
RUN bash -c "yarn build"

# 

FROM node:alpine

RUN apk add bash

RUN bash -c "yarn global add serve"
COPY --from=base /home/workspace/build ./build

EXPOSE 3000

CMD ["bash", "-c", "serve -l 3000 -s build"]