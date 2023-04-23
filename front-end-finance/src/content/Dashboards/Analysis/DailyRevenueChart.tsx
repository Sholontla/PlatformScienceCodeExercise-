import React from "react";
import ChartistTooltip from "chartist-plugin-tooltips-updated";
import dynamic from 'next/dynamic';

const Chartist = dynamic(() => import('react-chartist'), {
  ssr: false,
});

const DailyRevenueChart = () => {
  const [chart, setChart] = React.useState(null);

  const data = {
    labels: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
    series: [[1, 2, 2, 3, 3, 4, 3]]
  };

  const options = {
    low: 0,
    showArea: true,
    fullWidth: true,
    axisX: {
      position: "end",
      showGrid: true
    },
    axisY: {
      // On the y-axis start means left and end means right
      showGrid: false,
      showLabel: false,
      labelInterpolationFnc: value => `$${value / 1}k`
    }
  };

  const plugins = [ChartistTooltip()];

  React.useEffect(() => {
    if (typeof window !== "undefined") {
      setChart(<Chartist
        data={data}
        options={{ ...options, plugins }}
        type="Line"
        className="ct-series-g ct-double-octave"
      />);
    }
  }, []);

  return chart;
};

export default DailyRevenueChart;