FROM node:18.12.1-alpine3.16 as build

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn

COPY src ./src
COPY public ./public
COPY tsconfig.json next.config.js .eslintrc.json ./

RUN yarn build

FROM node:18.12.1-alpine3.16 as prod
WORKDIR /app
ENV NODE_ENV=production

COPY package.json yarn.lock ./

RUN yarn install --production

COPY --from=build /app/.next .next

CMD ["yarn", "start"]