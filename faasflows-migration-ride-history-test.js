import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    scenarios: {
        ride_history_load: {
            executor: 'constant-arrival-rate',
            rate: 200,  // 200 requests per second for Ride History node
            timeUnit: '1s',
            duration: '2m',  // 2 minutes duration
            preAllocatedVUs: 40,
            maxVUs: 250,     // Allow scaling up if needed
        },
    },
    thresholds: {
        http_req_duration: ['p(95)<3000'], // 95% of requests should complete within 3s
    },
};

export default function () {
    const userId = 10;
    const origin = {
        lat: 10.10,
        lon: 40.40
    };

    // Ride History request
    const url = 'http://127.0.0.1:8080/flow/ride-history';
    const payload = JSON.stringify({
        user_id: userId,
        origin: origin
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
        tags: { name: 'ride-history' }, // For better metrics reporting
    };

    const response = http.post(url, payload, params);

    check(response, {
        'status is 200': (r) => r.status === 200,
        'response body has content': (r) => r.body.length > 0,
    });

    // Small pause to ensure proper request distribution
    sleep(0.01);
}