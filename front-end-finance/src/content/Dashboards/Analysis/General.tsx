import { useState } from 'react';
import { Timeline, TimelineEvent } from 'react-event-timeline';

function General({dataSource}) {
  const [period, setPeriod] = useState('last_year');

  const periods = [
    {
      value: 'today',
      text: 'Today'
    },
    {
      value: 'yesterday',
      text: 'Yesterday'
    },
    {
      value: 'last_month',
      text: 'Last month'
    },
    {
      value: 'last_year',
      text: 'Last year'
    }
  ];

  const timelineData = [
    {
      time: new Date(2023, 3, 16, 13, 53, 41, 144496800),
      value: 51.25
    },
    {
      time: new Date(2023, 3, 16, 13, 53, 51, 888529900),
      value: 51.25
    },
    {
      time: new Date(2023, 3, 16, 14, 4, 56, 196813600),
      value: 51.25
    }
  ];

  return (
    <Timeline>
      {timelineData.map((dataPoint, index) => (
        <TimelineEvent
          key={index}
          title={dataPoint.value.toString()}
          createdAt={dataPoint.time.toString()}
          bubbleStyle={{ background: '#00C853' }}
          icon={<i className="fa fa-dollar"></i>}
        />
      ))}
    </Timeline>
  );
}

export default General;