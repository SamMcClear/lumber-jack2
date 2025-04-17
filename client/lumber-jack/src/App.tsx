import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [data, setData] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState("home");

  const fetchData = (tab: string) => {
    setLoading(true);
    fetch(`/api/${tab}`)
      .then((res) => {
        if (!res.ok) {
          throw new Error(`Failed to fetch data for ${tab}`);
        }
        return res.json();
      })
      .then((fetchedData) => {
        setData(fetchedData);
        setLoading(false);
      })
      .catch((error) => {
        console.error(`Error fetching ${tab} data:`, error);
        setLoading(false);
      });
  };

  useEffect(() => {
    fetchData(activeTab); // Fetch data for the active tab
  }, [activeTab]);

  const renderContent = () => {
    if (loading) {
      return <p className="loading-text">Retrieving data<span className="dots">...</span></p>;
    }

    switch (activeTab) {
      case "home":
        return (
          <div className="data-container">
            <div className="data-card">
              <h2>Message</h2>
              <p>{data?.message || "Welcome to LumberJack!"}</p>
            </div>
            <div className="data-card">
              <h2>CPU Info</h2>
              <p>{data?.cpuInfo || "No data available"}</p>
            </div>
            <div className="data-card">
              <h2>User IP</h2>
              <p>{data?.UserIP || "No data available"}</p>
            </div>
            <div className="data-card">
              <h2>User Info</h2>
              <p>{data?.userInfo || "No data available"}</p>
            </div>
          </div>
        );
      case "network":
        return (
          <div className="data-container">
            <div className="data-card">
              <h2>Network Usage</h2>
              <p>{data?.networkUsage || "No data available"}</p>
            </div>
          </div>
        );
      case "cpu":
        return (
          <div className="data-container">
            <div className="data-card">
              <h2>CPU Analytics</h2>
              <p>{data?.cpuInfo || "No data available"}</p>
            </div>
          </div>
        );
      case "storage":
        return (
          <div className="data-container">
            <div className="data-card">
              <h2>Storage Analytics</h2>
              <p>{data?.storageInfo || "No data available"}</p>
            </div>
          </div>
        );
      case "user":
        return (
          <div className="data-container">
            <div className="data-card">
              <h2>User and Server Info</h2>
              <p>{data?.userInfo || "No data available"}</p>
            </div>
          </div>
        );
      default:
        return <p>Welcome to the Dashboard</p>;
    }
  };

  return (
    <div className="dashboard-container">
      <aside className="sidebar">
        <h2>LumberJack Monitoring</h2>
        <ul>
          <li onClick={() => setActiveTab("home")}>Home</li>
          <li onClick={() => setActiveTab("network")}>Network</li>
          <li onClick={() => setActiveTab("cpu")}>CPU</li>
          <li onClick={() => setActiveTab("storage")}>Storage</li>
          <li onClick={() => setActiveTab("user")}>User Info</li>
        </ul>
      </aside>
      <main className="main-content">{renderContent()}</main>
    </div>
  );
}

export default App;
