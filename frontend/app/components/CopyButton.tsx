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
      className="text-xs font-mono text-[#5a7d99] hover:text-[#4db8ff] transition-colors px-2 py-1 rounded hover:bg-[#1c3d5c]"
    >
      {copied ? "copied ✓" : "copy"}
    </button>
  );
}
