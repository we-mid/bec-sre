# How to add plugin to RabbitMQ docker image?
# https://stackoverflow.com/questions/52819237/how-to-add-plugin-to-rabbitmq-docker-image
# https://blog.rabbitmq.com/posts/2015/04/scheduling-messages-with-rabbitmq/

# FROM rabbitmq:3.9-management
FROM rabbitmq:3-management-alpine

# COPY rabbitmq.conf /etc/rabbitmq/rabbitmq.conf

# RUN apt-get -o Acquire::Check-Date=false update && apt-get install -y curl
RUN apk update && apk add curl

# RUN curl -L https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/3.9.0/rabbitmq_delayed_message_exchange-3.9.0.ez\
#   > $RABBITMQ_HOME/plugins/rabbitmq_delayed_message_exchange-3.9.0.ez
RUN curl -L https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/v3.12.0/rabbitmq_delayed_message_exchange-3.12.0.ez \
  > $RABBITMQ_HOME/plugins/rabbitmq_delayed_message_exchange-3.12.0.ez

# RUN chown rabbitmq:rabbitmq $RABBITMQ_HOME/plugins/rabbitmq_delayed_message_exchange-3.9.0.ez
RUN chown rabbitmq:rabbitmq $RABBITMQ_HOME/plugins/rabbitmq_delayed_message_exchange-3.12.0.ez

RUN rabbitmq-plugins enable rabbitmq_delayed_message_exchange

# How to add initial users when starting a RabbitMQ Docker container?
# https://stackoverflow.com/questions/30747469/how-to-add-initial-users-when-starting-a-rabbitmq-docker-container
# ENV RABBITMQ_USER user
# ENV RABBITMQ_PASSWORD user
ENV RABBITMQ_PID_FILE /var/lib/rabbitmq/mnesia/rabbitmq
ADD init.sh /init.sh
RUN chmod +x /init.sh
CMD ["/init.sh"]

LABEL author="uxfritz@163.com" \
  description="RabbitMQ w/ delayed_message_exchange plugin"
