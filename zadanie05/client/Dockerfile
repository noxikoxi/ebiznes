FROM node:23-slim

WORKDIR /app
COPY . .

RUN npm install
RUN npm install -g serve
RUN npm run build

EXPOSE 5137

CMD ["serve", "-s", "dist", "-l", "5137"]