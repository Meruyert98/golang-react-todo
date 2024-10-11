# Use Node.js 16.x as the base image
FROM node:18 as build-stage

WORKDIR /app

COPY package.json ./

# Install dependencies
RUN npm install

COPY . .

# Build the project
RUN npm run build

# Production environment
FROM nginx:alpine

COPY --from=build-stage /app/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
