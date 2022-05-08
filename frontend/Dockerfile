FROM node:14-alpine3.14

WORKDIR /app


COPY package.json ./

ENV NODE_ENV production

RUN npm install
RUN npm install react-scripts -g
RUN npm install serve -g


COPY . ./

ENV REACT_APP_URL http://localhost:3000
ENV REACT_APP_PORT 8000

ENV REACT_APP_API_URL simple-backend.io
ENV REACT_APP_API_PORT 80

RUN npm run build

EXPOSE 3000

CMD ["serve", "-s", "build"]