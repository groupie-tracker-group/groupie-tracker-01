# THIS IS A DOCKERFILE FOR THE COUPIE-TRACKER APPLICATION
# IT WILL BUILD THE APPLICATION AND RUN IT ON PORT 8080
# THE APPLICATION WILL BE ACCESSIBLE ON THE HOST MACHINE VIA THE URL: http://localhost:8080
# ____________________________________________________________________________________________  #


# Use the official golang image to create a build artifact.
FROM  golang:1.20

# Create and change to the app directory.
WORKDIR /COUPIE-TRACKER

# Copy local code to the container image.
COPY  /COUPIE-TRACKER .


EXPOSE 3000
# Run the web service on container startup.
CMD ["go", "run", "main.go"]