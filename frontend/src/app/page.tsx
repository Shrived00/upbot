"use client";
import { useEffect, useState } from "react";

export default function Home() {
  const [message, setMessage] = useState("");
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/health`)
      .then((res) => res.json())
      .then((data) => setMessage(data.message))
      .catch((err) => setError(err.message || String(err)));
  }, []);

  if (error) {
    return (
      <h1>
        Error connecting to backend:
        <br />
        {error}
      </h1>
    );
  }

  return <h1>{message || "Loading..."}</h1>;
}
