package com.stockwatch.alertsevaluator;

import org.springframework.amqp.core.Binding;
import org.springframework.amqp.core.BindingBuilder;
import org.springframework.amqp.core.Queue;
import org.springframework.amqp.core.TopicExchange;
import org.springframework.amqp.rabbit.connection.CachingConnectionFactory;
import org.springframework.amqp.rabbit.connection.ConnectionFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class RabbitMqConfig {

    public static final String PRICES_EXCHANGE = "new_prices";
    public static final String ALERTS_EXCHANGE = "new_alerts";
    public static final String EVALUATOR_QUEUE = "alerts_evaluator";

    @Bean
    public ConnectionFactory rabbitConnectionFactory(@Value("${RMQ_ADDR_URL:amqp://guest:guest@localhost:5672/}") String rmqAddress) {
        CachingConnectionFactory factory = new CachingConnectionFactory();
        factory.setUri(rmqAddress);
        return factory;
    }

    @Bean
    public TopicExchange pricesExchange() {
        return new TopicExchange(PRICES_EXCHANGE, true, false);
    }

    @Bean
    public Queue evaluatorQueue() {
        return new Queue(EVALUATOR_QUEUE, true);
    }

    @Bean
    public Binding evaluatorBinding() {
        return BindingBuilder.bind(evaluatorQueue()).to(pricesExchange()).with("stocks.#");
    }

    @Bean
    public TopicExchange alertsExchange() {
        return new TopicExchange(ALERTS_EXCHANGE, true, false);
    }
}
