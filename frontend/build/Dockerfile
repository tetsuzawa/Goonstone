FROM node:12.16-alpine

RUN mkdir -p /app
WORKDIR /app
ADD . /app

RUN rm -rf node_modeles && npm install

RUN npm run build

CMD ["npm","run","start"]
