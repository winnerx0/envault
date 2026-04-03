import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "envault — Encrypted .env Backups",
  description:
    "Encrypt your .env files with AES-256-GCM and back them up to a private GitHub repository. Your secrets never leave your machine unencrypted.",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
