FROM node:16.13.0-buster as build

WORKDIR /Family-Tree

COPY package.json package.json
COPY yarn.lock yarn.lock

RUN yarn install --frozen-lockfile

COPY . .

RUN npm run build

# NGINX Web Server
FROM nginx:stable as prod

COPY --from=build /Family-Tree/build /var/www
COPY --from=build /Family-Tree/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD [ "nginx", "-g", "daemon off;" ]
