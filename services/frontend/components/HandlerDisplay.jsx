"use client";
import { useEffect, useState } from "react";

function HandlerDisplay({ data }) {
  const [myData, setMyData] = useState(data);
  useEffect(() => {
    fetch("http://localhost:8000/handler").then((x) => x.json()).then((x) =>
      setMyData(x)
    );
  }, []);

  return (
    <div>
      <div>Data is: {JSON.stringify(myData)}</div>
    </div>
  );
}

export default HandlerDisplay;
