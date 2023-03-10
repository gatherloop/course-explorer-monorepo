FROM node:16

WORKDIR /app

COPY dist/ .

RUN npm install --force

CMD ["npm", "start"]