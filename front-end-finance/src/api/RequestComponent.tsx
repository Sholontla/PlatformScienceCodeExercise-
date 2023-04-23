import React, { useState } from 'react';
import axios from 'axios';
import DailyRevenue from '@/content/Dashboards/Analysis/DailyRevenue';



const RequestComponent = () => {
  // Define state to hold the response data
  const [responseData, setResponseData] = useState<any>(null);

  // Define a function to make the request
  const makeRequest = async () => {
    try {
      const response = await axios.get('http://localhost:1003/service/daily/revenue');
      setResponseData(response.data);
      
    } catch (error) {
      console.error(error);
    }
  };

  // Render a button that triggers the request when clicked
  // and display the response data if it exists
  return (
    <div>
      <button onClick={makeRequest}>
        Make Request
      </button>
      {responseData ? (
        <div>
          <h2>Response Data:</h2>
          <p>{JSON.stringify(responseData)}</p>
          <DailyRevenue dataChart={responseData} />
        </div>
        
      ) : (
        <p>No response data yet.</p>
      )}
    </div>
  );
};

export default RequestComponent;

