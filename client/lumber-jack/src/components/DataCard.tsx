type DataCardProps = {
  title: string;
  content?: string;
};

const DataCard = ({ title, content }: DataCardProps) => (
  <div className="data-card">
    <h2>{title}</h2>
    <p>{content || "No data available"}</p>
  </div>
);

export default DataCard;

