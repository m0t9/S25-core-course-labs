# Setting base image as Python 3.11 slim
FROM python:3.11-slim AS build-env

COPY /static /app/static
COPY /templates /app/templates
COPY main.py /app/main.py
COPY requirements.txt /app/requirements.txt

WORKDIR /app

RUN pip install --no-cache-dir -r requirements.txt && cp $(which uvicorn) /app

FROM gcr.io/distroless/python3 as final

COPY --from=build-env /app /app
COPY --from=build-env /usr/local/lib/python3.11/site-packages /usr/local/lib/python3.11/site-packages
ENV PYTHONPATH=/usr/local/lib/python3.11/site-packages

WORKDIR /app

CMD ["./uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
