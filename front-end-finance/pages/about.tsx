import DailyRevenue from '@/content/Dashboards/Analysis/DailyRevenue';



function About() {

    const data = [
        {
          name: "Day 1",
          daily_revenue: 100
        },
        {
          name: "Day 2",
          daily_revenue: 150
        },
        {
          name: "Day 3",
          daily_revenue: 200
        }
      ];
  return (
    <>
      <div>
      <DailyRevenue dataChart={data}/>

      </div>
      
    </>
  );
}

export default About;
