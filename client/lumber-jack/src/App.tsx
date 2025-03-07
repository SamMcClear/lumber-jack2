import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [data, setData] = useState<
    { message: string; cpuInfo: string; UserIP: string } | null
  >(null);

  console.log(data);
  useEffect(() => {
    fetch("/api/hello")
      .then((res) => res.json())
      .then((data) => setData(data))
      .catch((error) => console.error("Error fetching API:", error));
  }, []);

  return (
    <div className="app-container">
      <h1 className="app-title">React + Go App</h1>
      {data ? (
        <div className="data-container">
          <div className="data-card">
            <h2>Message</h2>
            <p>{data.message}</p>
          </div>
          <div className="data-card">
            <h2>CPU Info</h2>
            <p>{data.cpuInfo}</p>
          </div>
          <div className="data-card">
            <h2>User IP</h2>
            <p>{data.UserIP}</p>
          </div>
        </div>
      ) : (
        <p className="loading-text">Loading...</p>
      )}
    </div>
  );
}

export default App;
