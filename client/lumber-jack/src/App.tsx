import { useEffect, useState } from "react";
import "./App.css";
import Sidebar from "./components/Sidebar";
import DataContent from "./components/DataContent";

type Tab = "home" | "network" | "cpu" | "storage" | "user";

function App() {
  const [data, setData] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState<Tab>("home");

  const fetchData = (tab: string) => {
    setLoading(true);
    fetch(`/api/${tab}`)
      .then((res) => {
        if (!res.ok) throw new Error(`Failed to fetch data for ${tab}`);
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

  return (
    <div className="dashboard-container">
      <Sidebar onSelectTab={setActiveTab} />
      <main className="main-content">
        <DataContent data={data} loading={loading} activeTab={activeTab} />
      </main>
    </div>
  );
}

export default App;

