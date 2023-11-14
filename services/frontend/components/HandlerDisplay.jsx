"use client";
import { useEffect, useState } from "react";

function HandlerDisplay({ data }) {
  const [myData, setMyData] = useState(data);
  useEffect(() => {
    fetch("http://localhost:8000/handler").then((x) => x.json()).then((x) =>
      setMyData(x)
    ).catch(e => {
      console.error(' failed to fetch: ', e)
    });
  }, []);

  return (
    <div>
      <h2> handler tester</h2>
      <div>Data is: {JSON.stringify(myData)}</div>
    </div>
  );
}

export default HandlerDisplay;
