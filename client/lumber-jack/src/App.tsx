import { useEffect, useState } from "react";
import "./App.css";

type Tab = "home" | "network" | "cpu" | "storage" | "user";

function App() {
  const [data, setData] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState<Tab>("home");

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
    fetchData(activeTab);
  }, [activeTab]);

  const renderContent = () => {
    if (loading) {
      return (
        <p className="loading-text">
          Retrieving data<span className="dots">...</span>
        </p>
      );
    }

    const cards: Record<Tab, { title: string; content?: string; key: string }[]> = {
      home: [
        { title: "Message", content: data?.message || "Welcome to LumberJack!", key: "message" },
        { title: "CPU Info", content: data?.cpuInfo, key: "cpu" },
        { title: "User IP", content: data?.UserIP, key: "ip" },
        { title: "User Info", content: data?.userInfo, key: "user" },
      ],
      network: [
        { title: "Network Usage", content: data?.networkUsage, key: "network" },
      ],
      cpu: [
        { title: "CPU Analytics", content: data?.cpuInfo, key: "cpu" },
      ],
      storage: [
        { title: "Storage Analytics", content: data?.storageInfo, key: "storage" },
      ],
      user: [
        { title: "User and Server Info", content: data?.userInfo, key: "userInfo" },
      ],
    };

    return (
      <div className="data-container">
        {cards[activeTab].map(({ title, content, key }) => (
          <div key={key} className="data-card">
            <h2>{title}</h2>
            <p>{content || "No data available"}</p>
          </div>
        ))}
      </div>
    );
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
