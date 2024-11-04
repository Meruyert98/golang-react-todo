# Use Node.js 16.x as the base image
FROM node:18 as build-stage

WORKDIR /app

COPY package.json ./

# Install dependencies
RUN npm install

COPY . .

# Set environment variable for the API endpoint
ARG REACT_APP_API_ENDPOINT
ENV REACT_APP_API_ENDPOINT=${REACT_APP_API_ENDPOINT}

# Build the project
RUN npm run build

# Production environment
FROM nginx:alpine

COPY --from=build-stage /app/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
