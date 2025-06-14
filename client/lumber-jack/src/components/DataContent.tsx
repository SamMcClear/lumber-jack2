import DataCard from "./DataCard";
import Loader from "./Loader";

type Tab = "home" | "network" | "cpu" | "storage" | "user";

type DataContentProps = {
  data: any;
  loading: boolean;
  activeTab: Tab;
};

const DataContent = ({ data, loading, activeTab }: DataContentProps) => {
  if (loading) return <Loader />;

  const cards: Record<Tab, { title: string; content?: string; key: string }[]> = {
    home: [
      { title: "Message", content: data?.message || "Welcome to LumberJack!", key: "message" },
      { title: "CPU Info", content: data?.cpuInfo, key: "cpu" },
      { title: "User IP", content: data?.UserIP, key: "ip" },
      { title: "User Info", content: data?.userInfo, key: "user" },
    ],
    network: [{ title: "Network Usage", content: data?.networkUsage, key: "network" }],
    cpu: [{ title: "CPU Analytics", content: data?.cpuInfo, key: "cpu" }],
    storage: [{ title: "Storage Analytics", content: data?.storageInfo, key: "storage" }],
    user: [{ title: "User and Server Info", content: data?.userInfo, key: "userInfo" }],
  };

  return (
    <div className="data-container">
      {cards[activeTab].map(({ title, content, key }) => (
        <DataCard key={key} title={title} content={content} />
      ))}
    </div>
  );
};

export default DataContent;

