FROM node:12.16-alpine

RUN mkdir -p /src/app
WORKDIR /src/app
ADD . /src/app

RUN npm install

CMD ["npm","run","dev"]
