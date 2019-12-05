FROM gcc:8.3.0
WORKDIR /app
COPY ./runner /runner

ENV RUNNER_SOURCE_FILENAME="Main.cc"
ENV RUNNER_COMPILE="g++ Main.cc -o Main -O2 -Wall -lm -static -std=gnu++14 -DONLINE_JUDGE -DBOJ"
ENV RUNNER_SCRIPT="./Main"

RUN chmod +x /runner
EXPOSE 5000
ENTRYPOINT /runner
