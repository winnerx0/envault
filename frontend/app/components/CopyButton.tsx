"use client";

import { useState } from "react";

export default function CopyButton({ text }: { text: string }) {
  const [copied, setCopied] = useState(false);

  const handleCopy = () => {
    navigator.clipboard.writeText(text);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <button
      onClick={handleCopy}
      className="text-xs font-mono text-[#71717a] hover:text-[#a78bfa] transition-colors px-2 py-1 rounded hover:bg-[#a78bfa]/10"
    >
      {copied ? "copied ✓" : "copy"}
    </button>
  );
}
