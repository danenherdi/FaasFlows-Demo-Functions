import http from 'k6/http';
import { sleep } from 'k6';

// Constants for RPS (Requests Per Second) for each endpoint
const HOMEPAGE_RPS = 100;
const RIDE_HISTORY_RPS = 200;
const FRIENDS_RPS = 50;

// Total RPS
const TOTAL_RPS = HOMEPAGE_RPS + RIDE_HISTORY_RPS + FRIENDS_RPS;

// Probabilities for each endpoint based on RPS
const HOMEPAGE_PROBABILITY = HOMEPAGE_RPS / TOTAL_RPS;
const RIDE_HISTORY_PROBABILITY = RIDE_HISTORY_RPS / TOTAL_RPS;
// Friends probability = 1 - (homepage + ride_history)

export const options = {
    // Define the stages for ramping up the load
    scenarios: {
        constant_load: {
            executor: 'constant-arrival-rate',
            rate: TOTAL_RPS,
            timeUnit: '1s',
            duration: '5m',
            preAllocatedVUs: 50,
            maxVUs: 500,
        },
    },
};

export default function () {
    const userId = 10;

    const origin = {
        lat: 10.10,
        lon: 40.40
    };

    // Randomly select an endpoint based on the defined probabilities
    const random = Math.random();
    let url, payload;

    if (random < HOMEPAGE_PROBABILITY) {
        // Homepage request (100 RPS)
        url = 'http://127.0.0.1:8080/flow/homepage';
        payload = JSON.stringify({
            user_id: userId,
            origin: origin
        });
    } else if (random < HOMEPAGE_PROBABILITY + RIDE_HISTORY_PROBABILITY) {
        // Ride History request (200 RPS)
        url = 'http://127.0.0.1:8080/flow/ride-history';
        payload = JSON.stringify({
            user_id: userId,
            origin: origin
        });
    } else {
        // Friends request (50 RPS)
        url = 'http://127.0.0.1:8080/flow/friends';
        payload = JSON.stringify({
            user_id: userId,
            origin: origin
        });
    }

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const response = http.post(url, payload, params);

    check(response, {
        'status is 200': (r) => r.status === 200,
        'response body has content': (r) => r.body.length > 0,
    });

    // Sleep for a short duration to control the request rate
    sleep(0.01);
}