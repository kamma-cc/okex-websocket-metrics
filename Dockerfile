FROM openjdk:11-jdk-slim AS build-env
ADD . /app
WORKDIR /app
RUN gradlew -Dskip.tests=true build

FROM gcr.io/distroless/java:11
COPY --from=build-env /app/build/libs /app
WORKDIR /app
CMD ["java", "-jar", "okex-websocket-metrics.jar"]