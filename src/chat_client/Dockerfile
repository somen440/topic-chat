FROM python:3-slim
ADD /dist/main.js /dist/main.js
ADD index.html index.html
CMD ["python", "-m", "http.server", "8000"]
