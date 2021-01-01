FROM debian:bullseye-slim
WORKDIR /app
ADD go-todo-linux .
CMD ["/app/go-todo-linux"]


