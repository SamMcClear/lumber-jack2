type Tab = "home" | "network" | "cpu" | "storage" | "user";

type SidebarProps = {
  onSelectTab: (tab: Tab) => void;
};

const Sidebar = ({ onSelectTab }: SidebarProps) => (
  <aside className="sidebar">
    <h2>LumberJack Monitoring</h2>
    <ul>
      <li onClick={() => onSelectTab("home")}>Home</li>
      <li onClick={() => onSelectTab("network")}>Network</li>
      <li onClick={() => onSelectTab("cpu")}>CPU</li>
      <li onClick={() => onSelectTab("storage")}>Storage</li>
      <li onClick={() => onSelectTab("user")}>User Info</li>
    </ul>
  </aside>
);

export default Sidebar;

