import { useState, useEffect } from 'react';

function App() {

  const [data, setData] = useState<{ message: string; cpuInfo: string } | null>(null);

  useEffect(() => {
    fetch('/api/hello')
      .then((res) => res.json())
      .then((data) => setData(data))
      .catch((error) => console.error('Error fetching API:', error));
  }, []);

  return (
    <div>
      <h1>React + Go App</h1>
      {data ? (
        <div>
          <p>Message: {data.message}</p>
          <p>CPU Info: {data.cpuInfo}</p>
        </div>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
}

export default App;
