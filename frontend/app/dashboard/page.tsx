"use client";

import { useEffect, useState } from "react";
import { authClient } from "@/lib/auth-client";

export default function Dashboard() {
  const [user, setUser] = useState<any>(null);

  useEffect(() => {
    authClient.getSession().then((session) => {
      setUser(session?.user);
    });
  }, []);

  const handleLogout = async () => {
    await authClient.signOut();
    window.location.href = "/";
  };

  return (
    <main className="p-10">
      <h1 className="mb-4">Welcome, {user?.email || "User"}</h1>
      <button
        className="bg-red-600 text-white px-4 py-2 rounded"
        onClick={handleLogout}
      >
        Logout
      </button>
    </main>
  );
}
