FROM node:12.3 as build
WORKDIR /app
COPY app/package*.json ./
RUN npm install
COPY app/ .
RUN npm run build

FROM node:12.3-alpine as run
WORKDIR /app
COPY --from=build /app/dist ./dist/
RUN apk add dumb-init
COPY app/package*.json ./
RUN npm install --production
CMD ["dumb-init", "node", "dist/index.js"]
