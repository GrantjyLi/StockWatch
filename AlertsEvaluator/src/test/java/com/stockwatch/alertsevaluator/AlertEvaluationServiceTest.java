package com.stockwatch.alertsevaluator;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class AlertEvaluationServiceTest {

    private final AlertEvaluationService service = new AlertEvaluationService();

    @Test
    void shouldMatchThresholdOperators() {
        assertTrue(service.matchesOperator(">=", 100.0, 100.0));
        assertTrue(service.matchesOperator(">=", 101.0, 100.0));
        assertTrue(service.matchesOperator("<=", 99.0, 100.0));
        assertTrue(service.matchesOperator("=", 100.0, 100.0));

        assertFalse(service.matchesOperator(">=", 99.0, 100.0));
        assertFalse(service.matchesOperator("<=", 101.0, 100.0));
        assertFalse(service.matchesOperator("=", 101.0, 100.0));
    }
}
