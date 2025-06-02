// app/page.tsx
"use client";

import { authClient } from "@/lib/auth-client";

export default function Home() {
  const handleGoogleLogin = async () => {
    await authClient.signIn.social({
      provider: "google",
      callbackURL: "/dashboard",
      errorCallbackURL: "/login?error=oauth",
    });
  };

  return (
    <main className="p-10">
      <button
        className="bg-blue-600 text-white px-4 py-2 rounded"
        onClick={handleGoogleLogin}
      >
        Sign in with Google
      </button>
    </main>
  );
}
