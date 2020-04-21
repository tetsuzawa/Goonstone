FROM node:12.16-alpine

RUN mkdir -p /app
WORKDIR /app
ADD . /app

RUN npm install && \
    npm run build

CMD ["npm","run","start"]
