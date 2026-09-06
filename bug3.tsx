function StatusBadge({ score }: { score: number }) {
  if (score >= 50) {
    return <span className="pass">Pass</span>;
  } else {
    return <span className="fail">Fail</span>;
  }
  console.log("this line is dead code");
}

export default StatusBadge;
