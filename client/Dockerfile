#Starting base image
FROM node:alpine

#Setting up working directory
WORKDIR /app

#Copying and downloading dependencies
# COPY package.json .
# RUN npm install

#Copying everything else
COPY . .

#Setting up starting command
CMD ["npm", "run", "dev"]