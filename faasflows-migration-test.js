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
    const validUserIds = [10, 20, 30, 40];
    const userId = validUserIds[Math.floor(Math.random() * validUserIds.length)];
    const origin = `${(Math.random() * 90).toFixed(6)},${(Math.random() * 90).toFixed(6)}`;

    // Randomly select an endpoint based on the defined probabilities
    const random = Math.random();
    let url, payload;

    if (random < HOMEPAGE_PROBABILITY) {
        // Homepage request (100 RPS)
        url = 'http://gateway.openfaas:8080/flow/homepage';
        payload = JSON.stringify({
            user_id: userId,
            origin: origin
        });
    } else if (random < HOMEPAGE_PROBABILITY + RIDE_HISTORY_PROBABILITY) {
        // Ride History request (200 RPS)
        url = 'http://gateway.openfaas:8080/flow/ride-history';
        payload = JSON.stringify({
            user_id: userId
        });
    } else {
        // Friends request (50 RPS)
        url = 'http://gateway.openfaas:8080/flow/friends';
        payload = JSON.stringify({
            user_id: userId
        });
    }

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);

    // Sleep for a short duration to control the request rate
    sleep(0.01);
}