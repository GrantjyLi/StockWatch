package com.stockwatch.alertsevaluator;

import com.rabbitmq.client.Channel;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.amqp.core.Message;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.amqp.support.AmqpHeaders;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.messaging.handler.annotation.Header;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.math.BigDecimal;
import java.util.List;

@Service
public class AlertEvaluationService {

    // Central logger for the alert evaluation flow.
    private static final Logger log = LoggerFactory.getLogger(AlertEvaluationService.class);
    // RabbitMQ exchange used to publish triggered alerts.
    private static final String ALERTS_EXCHANGE = "new_alerts";

    // Data access for alert/watchlist lookups.
    private final JdbcTemplate jdbcTemplate;
    // Messaging client used to publish triggered alerts.
    private final RabbitTemplate rabbitTemplate;

    @Autowired
    public AlertEvaluationService(JdbcTemplate jdbcTemplate, RabbitTemplate rabbitTemplate) {
        this.jdbcTemplate = jdbcTemplate;
        this.rabbitTemplate = rabbitTemplate;
    }

    // Checks whether the current price satisfies the alert's comparison operator.
    public boolean matchesOperator(String operator, double currentPrice, double targetPrice) {
        if (operator == null) {
            return false;
        }

        return switch (operator) {
            case ">=" -> currentPrice >= targetPrice;
            case "<=" -> currentPrice <= targetPrice;
            case "=" -> Double.compare(currentPrice, targetPrice) == 0;
            default -> false;
        };
    }

    // Queries all active alerts that should fire for the incoming price update.
    public List<TriggeredAlert> findTriggeredAlerts(PriceUpdate update) {
        String sql = """
            select
                a.id as alert_id,
                a.target_price,
                a.operator,
                u.email as user_email
            from alerts a
            join watchlists w on a.watchlist_id = w.id
            join users u on w.user_id = u.id
            where a.ticker = ?
              and a.triggered = false
              and (
                  (a.operator = '>=' and ? >= a.target_price) or
                  (a.operator = '<=' and ? <= a.target_price) or
                  (a.operator = '=' and ? = a.target_price)
              )
            """;

        return jdbcTemplate.query(sql, (rs, rowNum) -> {
            TriggeredAlert alert = new TriggeredAlert();
            alert.setAlertId(rs.getString("alert_id"));
            alert.setTicker(update.getTicker());
            alert.setTargetPrice(rs.getBigDecimal("target_price"));
            alert.setOperator(rs.getString("operator"));
            alert.setUserEmail(rs.getString("user_email"));
            return alert;
        }, update.getTicker(), update.getPrice(), update.getPrice(), update.getPrice());
    }

    // Consumes price update messages from RabbitMQ and acknowledges only on successful processing.
    @RabbitListener(queues = "alerts_evaluator", ackMode = "MANUAL")
    public void onPriceUpdate(@Payload PriceUpdate update,
                             @Header(AmqpHeaders.DELIVERY_TAG) long tag,
                             Channel channel) throws IOException {
        try {
            processPriceUpdate(update);
            channel.basicAck(tag, false);
        } catch (Exception e) {
            log.error("Failed to evaluate price update {}", update, e);
            channel.basicNack(tag, false, false);
        }
    }

    // Evaluates a price update against all matching alerts and publishes each triggered alert.
    public void processPriceUpdate(PriceUpdate update) {
        List<TriggeredAlert> triggeredAlerts = findTriggeredAlerts(update);

        for (TriggeredAlert alert : triggeredAlerts) {
            alert.setTicker(update.getTicker());
            log.info("Alert triggered: {}", alert.getAlertId());
            rabbitTemplate.convertAndSend(ALERTS_EXCHANGE, "alerts." + alert.getAlertId(), alert);
        }
    }
}
