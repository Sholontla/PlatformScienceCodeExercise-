import { CartesianGrid, Legend, Line, LineChart, Tooltip, XAxis, YAxis } from "recharts";

const DailyRevenue = ({ dataChart }) => {
    return (
        <LineChart width={1000} height={400} data={dataChart}>
            <Line type ="monotone" dataKey="daily_revenue" stroke="#78f7ff" strokeWidth={3}/>
            <CartesianGrid stroke="#ccc"/>
            <XAxis dataKey="time"/>
            <YAxis/>
            <Tooltip/>
            <Legend/>
        </LineChart>
    );
};

export default DailyRevenue;