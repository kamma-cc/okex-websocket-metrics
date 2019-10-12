FROM openjdk:11-jdk-slim AS build-env
ADD . /app
WORKDIR /app
ENV GRADLE_OPTS -Dorg.gradle.daemon=false
RUN /app/gradlew -Dskip.tests=true build

FROM gcr.io/distroless/java:11
COPY --from=build-env /app/build/libs /app
WORKDIR /app
CMD ["java", "-jar", "okex-websocket-metrics.jar"]