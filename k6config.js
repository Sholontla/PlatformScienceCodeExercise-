import { Options } from 'k6/options';
import { options } from './k6config.js';
import { Trend } from 'k6/metrics';

export const options = new Options({
  vus: 10,
  duration: '30s',
  thresholds: {
    http_req_duration: ['p(95)<500'],
  },
  ext: {
    loadimpact: {
      projectID: 123456,
      name: 'My Test',
    },
  },
  http: {
    retries: 3,
    timeout: '10s',
    headers: {
      'User-Agent': 'k6',
    },
  },
});

export const myTrend = new Trend('my_trend');
``