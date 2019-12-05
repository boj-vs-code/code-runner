FROM python:3.7.4
WORKDIR /app

COPY ./runner /runner




ENV RUNNER_SOURCE_FILENAME="Main.py"

ENV RUNNER_SCRIPT="python3 Main.py"

ENV RUNNER_COMPILE="python3 -c \"import py_compile; py_compile.compile(r'Main.py')\""


RUN chmod +x /runner
EXPOSE 5000
ENTRYPOINT /runner