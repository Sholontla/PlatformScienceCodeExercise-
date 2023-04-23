import http from 'k6/http';
import { options, myTrend } from './k6config.js';

export default function () {
  for (let i = 0; i < 10; i++) {
    const res = http.get('http://localhost:1003/service/daily/revenue');
    myTrend.add(res.timings.duration);
  }
}

export function handleSummary(data) {
  myTrend.add(data.metrics.http_req_duration_percentiles['95']);
}
