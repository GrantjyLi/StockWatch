package com.stockwatch.alertsevaluator;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.math.BigDecimal;

public class TriggeredAlert {

    @JsonProperty("alert_ID")
    private String alertId;

    @JsonProperty("ticker")
    private String ticker;

    @JsonProperty("target_price")
    private BigDecimal targetPrice;

    @JsonProperty("operator")
    private String operator;

    @JsonProperty("user_email")
    private String userEmail;

    public String getAlertId() {
        return alertId;
    }

    public void setAlertId(String alertId) {
        this.alertId = alertId;
    }

    public String getTicker() {
        return ticker;
    }

    public void setTicker(String ticker) {
        this.ticker = ticker;
    }

    public BigDecimal getTargetPrice() {
        return targetPrice;
    }

    public void setTargetPrice(BigDecimal targetPrice) {
        this.targetPrice = targetPrice;
    }

    public String getOperator() {
        return operator;
    }

    public void setOperator(String operator) {
        this.operator = operator;
    }

    public String getUserEmail() {
        return userEmail;
    }

    public void setUserEmail(String userEmail) {
        this.userEmail = userEmail;
    }
}
