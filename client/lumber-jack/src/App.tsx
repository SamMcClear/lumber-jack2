import { useState, useEffect } from 'react';

function App() {
  const [data, setData] = useState<string | null>(null);

  useEffect(() => {
    fetch('/api/hello')
      .then((res) => res.json())
      .then((data) => setData(data.message))
      .catch((error) => console.error('Error fetching API:', error));
  }, []);

  return (
    <div>
      <h1>React + Go App</h1>
      <p>{data ? data : 'Loading...'}</p>
    </div>
  );
}

export default App;
