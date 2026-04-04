const TERMINAL_LINES = [
  { type: "prompt", text: "envault login -p mysecret -t ghp_xxxx" },
  { type: "out",    text: "Project root: myapp" },
  { type: "out",    text: "Created private backup repo: alice/3f9a2b..." },
  { type: "out",    text: "Config saved to ~/.envault/config.yaml" },
  { type: "prompt", text: "envault init" },
  { type: "out",    text: "Initialized envault project in /projects/myapp" },
  { type: "prompt", text: "envault backup" },
  { type: "out",    text: "Encrypted env files backed up to alice/3f9a2b..." },
];



export default function Home() {
  return (
    <div className="min-h-screen bg-[#09090b] text-[#fafafa] font-sans">

      {/* ── Nav ──────────────────────────────────────────────── */}
      <nav className="sticky top-0 z-50 border-b border-[#27272a] bg-[#09090b]/85 backdrop-blur-xl">
        <div className="max-w-6xl mx-auto px-6 h-14 flex items-center justify-between">
          <span className="font-semibold text-sm tracking-widest text-white">ENVAULT</span>
          <a
            href="https://github.com/winnerx0/envault"
            target="_blank"
            rel="noopener noreferrer"
            className="text-sm text-[#71717a] hover:text-white transition-colors"
          >
            github ↗
          </a>
        </div>
      </nav>

      {/* ── Hero ─────────────────────────────────────────────── */}
      <section className="max-w-4xl mx-auto px-6 pt-24 pb-20 text-center">

        <div className="fade-up fade-up-1 inline-flex items-center gap-2 bg-[#18181b] border border-[#27272a] rounded-full px-4 py-1.5 text-xs text-[#71717a] mb-8 font-mono">
          <span className="w-1.5 h-1.5 rounded-full bg-[#8b5cf6] inline-block animate-pulse" />
          Open source · AES-256-GCM · Argon2id
        </div>

        <h1 className="fade-up fade-up-2 text-[clamp(36px,6.5vw,68px)] font-bold leading-[1.08] tracking-tight mb-6">
          The{" "}
          <code className="font-mono text-[#a78bfa] not-italic">.env</code>{" "}
          backup tool<br />
          you&apos;ll actually trust.
        </h1>

        <p className="fade-up fade-up-3 text-[#71717a] text-lg leading-relaxed max-w-xl mx-auto mb-10">
          Encrypt and back up your .env files to a private GitHub repo.
          AES-256-GCM encryption — secrets never leave your machine in plaintext.
        </p>

        <div className="fade-up fade-up-4 flex flex-col items-center gap-2 mb-16">
          <code className="font-mono text-xs text-[#a78bfa] bg-[#111117] border border-[#27272a] rounded-lg px-4 py-2.5 select-all">
            curl -fsSL https://raw.githubusercontent.com/winnerx0/envault/main/install.sh | sh
          </code>
          <p className="font-mono text-xs text-[#3f3f46]">
            or: go install github.com/winnerx0/envault/cmd/envault@latest
          </p>
        </div>

        <div className="fade-up fade-up-5 terminal text-left max-w-2xl mx-auto">
          <div className="terminal-chrome">
            <div className="w-3 h-3 rounded-full bg-[#ff5f57]" />
            <div className="w-3 h-3 rounded-full bg-[#febc2e]" />
            <div className="w-3 h-3 rounded-full bg-[#28c840]" />
            <span className="ml-2 font-mono text-xs text-[#71717a]">zsh — envault demo</span>
          </div>
          <div className="terminal-body space-y-0.5 min-h-[180px]">
            {TERMINAL_LINES.map((line, i) => (
              <div
                key={i}
                className={`term-line flex gap-2 ${
                  line.type === "out" ? "text-[#71717a]" : "text-[#fafafa]"
                }`}
              >
                <span
                  className={`select-none shrink-0 ${
                    line.type === "prompt" ? "text-[#a78bfa]" : "text-[#3f3f46]"
                  }`}
                >
                  {line.type === "prompt" ? "$" : "›"}
                </span>
                <span>{line.text}</span>
              </div>
            ))}
            <div className="term-line flex gap-2">
              <span className="text-[#a78bfa] select-none">$</span>
              <span className="cursor-blink" />
            </div>
          </div>
        </div>
      </section>

      {/* ── Feature Bento Rows ───────────────────────────────── */}
      <section className="max-w-6xl mx-auto px-6 py-20">
        <div className="rounded-2xl overflow-hidden border border-[#27272a] divide-y divide-[#27272a]">

          {/* Row 1 — AES-256-GCM */}
          <div className="grid grid-cols-1 lg:grid-cols-2 min-h-[380px] divide-y lg:divide-y-0 lg:divide-x divide-[#27272a]">
            <div className="bg-[#4c1d95] p-10 lg:p-14 flex flex-col justify-center">
              <span className="inline-block bg-white/10 rounded-full px-3 py-1 text-[11px] text-white/60 mb-6 tracking-wider uppercase w-fit">
                Encryption
              </span>
              <h3 className="text-[clamp(24px,3vw,34px)] font-bold text-white leading-tight mb-3">
                AES-256-GCM<br />encryption
              </h3>
              <p className="text-white/60 text-sm leading-relaxed max-w-xs">
                Military-grade authenticated encryption. Each .env file is independently encrypted with a unique 16-byte nonce before touching the network.
              </p>
            </div>
            <div className="bg-[#0f0f14] p-10 lg:p-14 flex items-center justify-center">
              <div className="w-full max-w-xs">
                <div className="bg-[#111117] border border-[#27272a] rounded-xl p-5 font-mono text-xs">
                  <div className="flex items-center gap-2 mb-4">
                    <span className="w-2 h-2 rounded-full bg-[#8b5cf6]" />
                    <span className="text-[#71717a] text-[10px] tracking-wider">myapp/.env.enc</span>
                  </div>
                  <div className="space-y-2">
                    <div>
                      <p className="text-[#52525b] text-[9px] tracking-widest uppercase mb-1">salt · 16 bytes</p>
                      <p className="text-[#6d28d9] text-[10px] break-all leading-5">3a8f2b91c4d7e6f0a1b2c3d4e5f6a7b8</p>
                    </div>
                    <div>
                      <p className="text-[#52525b] text-[9px] tracking-widest uppercase mb-1">nonce · 12 bytes</p>
                      <p className="text-[#7c3aed] text-[10px] break-all leading-5">9e1f3c7d2a5b8e4fc0d1e2f3</p>
                    </div>
                    <div>
                      <p className="text-[#52525b] text-[9px] tracking-widest uppercase mb-1">ciphertext</p>
                      <p className="text-[#a78bfa] text-[10px] break-all leading-5">f2c1b3d4e5a6f7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2…</p>
                    </div>
                  </div>
                </div>
                <p className="text-[#3f3f46] text-[10px] text-center mt-3 font-mono">plaintext never stored</p>
              </div>
            </div>
          </div>

          {/* Row 2 — One command backup */}
          <div className="grid grid-cols-1 lg:grid-cols-2 min-h-[380px] divide-y lg:divide-y-0 lg:divide-x divide-[#27272a]">
            <div className="bg-[#0f0f14] p-10 lg:p-14 flex items-center justify-center">
              <div className="w-full max-w-xs space-y-3">
                <div className="bg-[#111117] border border-[#27272a] rounded-xl p-4 font-mono text-xs">
                  <div className="flex items-center gap-2 mb-3">
                    <span className="w-2 h-2 rounded-full bg-green-400" />
                    <span className="text-[#71717a] text-[10px]">envault.json</span>
                  </div>
                  <pre className="text-[#a78bfa] text-[11px] leading-relaxed">{`{
  "name": "myapp",
  "version": "1.0"
}`}</pre>
                </div>
                <div className="bg-[#111117] border border-[#27272a] rounded-xl p-4 font-mono text-xs">
                  <div className="flex items-center gap-2 mb-3">
                    <span className="w-2 h-2 rounded-full bg-blue-400" />
                    <span className="text-[#71717a] text-[10px]">~/.envault/config.yaml</span>
                  </div>
                  <div className="space-y-1.5 text-[11px]">
                    <div className="flex gap-2">
                      <span className="text-[#71717a]">token:</span>
                      <span className="text-[#fafafa]">ghp_xxxx</span>
                    </div>
                    <div className="flex gap-2">
                      <span className="text-[#71717a]">passphrase:</span>
                      <span className="text-[#52525b]">••••••••</span>
                    </div>
                    <div className="flex gap-2">
                      <span className="text-[#71717a]">repo:</span>
                      <span className="text-[#fafafa]">alice/3f9a2b</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div className="bg-[#312e81] p-10 lg:p-14 flex flex-col justify-center">
              <span className="inline-block bg-white/10 rounded-full px-3 py-1 text-[11px] text-white/60 mb-6 tracking-wider uppercase w-fit">
                Backup
              </span>
              <h3 className="text-[clamp(24px,3vw,34px)] font-bold text-white leading-tight mb-3">
                One command<br />backup
              </h3>
              <p className="text-white/60 text-sm leading-relaxed max-w-xs">
                Run{" "}
                <code className="font-mono bg-white/10 px-1.5 py-0.5 rounded text-white/80 text-[11px]">
                  envault backup
                </code>{" "}
                and every .env file is encrypted and pushed to your private GitHub repo instantly.
              </p>
            </div>
          </div>

          {/* Row 3 — Argon2id */}
          <div className="grid grid-cols-1 lg:grid-cols-2 min-h-[380px] divide-y lg:divide-y-0 lg:divide-x divide-[#27272a]">
            <div className="bg-[#5b21b6] p-10 lg:p-14 flex flex-col justify-center">
              <span className="inline-block bg-white/10 rounded-full px-3 py-1 text-[11px] text-white/60 mb-6 tracking-wider uppercase w-fit">
                Key Derivation
              </span>
              <h3 className="text-[clamp(24px,3vw,34px)] font-bold text-white leading-tight mb-3">
                Argon2id<br />key derivation
              </h3>
              <p className="text-white/60 text-sm leading-relaxed max-w-xs">
                Memory-hard KDF with 3 iterations and 32 MB of memory. Brute force attacks become computationally infeasible even with dedicated hardware.
              </p>
            </div>
            <div className="bg-[#0f0f14] p-10 lg:p-14 flex items-center justify-center">
              <div className="flex flex-col items-center gap-2 w-full max-w-[220px]">
                <div className="w-full bg-[#111117] border border-[#27272a] rounded-xl px-5 py-3 text-center">
                  <p className="text-[#52525b] text-[9px] tracking-widest uppercase mb-1">INPUT</p>
                  <p className="text-[#fafafa] font-mono text-xs">passphrase</p>
                </div>
                <div className="flex flex-col items-center gap-0.5">
                  <div className="w-px h-4 bg-[#27272a]" />
                  <div className="bg-[#1a1040] border border-[#7c3aed] rounded-full px-4 py-1.5 text-[11px] text-[#a78bfa] font-mono whitespace-nowrap">
                    Argon2id · 3 iters · 32 MB
                  </div>
                  <div className="w-px h-4 bg-[#27272a]" />
                </div>
                <div className="w-full bg-[#111117] border border-[#8b5cf6]/40 rounded-xl px-5 py-3 text-center">
                  <p className="text-[#52525b] text-[9px] tracking-widest uppercase mb-1">KEY</p>
                  <p className="text-[#a78bfa] font-mono text-xs">256-bit AES key</p>
                </div>
                <div className="flex flex-col items-center gap-0.5">
                  <div className="w-px h-4 bg-[#27272a]" />
                  <div className="bg-[#111117] border border-[#27272a] rounded-full px-4 py-1.5 text-[11px] text-[#71717a] font-mono whitespace-nowrap">
                    AES-256-GCM + nonce
                  </div>
                  <div className="w-px h-4 bg-[#27272a]" />
                </div>
                <div className="w-full bg-[#111117] border border-[#27272a] rounded-xl px-5 py-3 text-center">
                  <p className="text-[#52525b] text-[9px] tracking-widest uppercase mb-1">OUTPUT</p>
                  <p className="text-[#6d28d9] font-mono text-xs">encrypted .env</p>
                </div>
              </div>
            </div>
          </div>

          {/* Row 4 — Recover anywhere */}
          <div className="grid grid-cols-1 lg:grid-cols-2 min-h-[380px] divide-y lg:divide-y-0 lg:divide-x divide-[#27272a]">
            <div className="bg-[#0f0f14] p-10 lg:p-14 flex items-center justify-center">
              <div className="w-full max-w-xs">
                <div className="bg-[#111117] border border-[#27272a] rounded-xl overflow-hidden font-mono text-xs">
                  <div className="bg-[#18181b] px-4 py-2.5 flex items-center gap-2 border-b border-[#27272a]">
                    <div className="w-2.5 h-2.5 rounded-full bg-[#ff5f57]" />
                    <div className="w-2.5 h-2.5 rounded-full bg-[#febc2e]" />
                    <div className="w-2.5 h-2.5 rounded-full bg-[#28c840]" />
                    <span className="ml-1 text-[#71717a] text-[10px]">new machine</span>
                  </div>
                  <div className="p-4 space-y-2 text-[11px] leading-relaxed">
                    <div className="flex gap-2">
                      <span className="text-[#a78bfa]">$</span>
                      <span className="text-[#fafafa]">envault recover</span>
                    </div>
                    <div className="flex gap-2">
                      <span className="text-[#3f3f46]">›</span>
                      <span className="text-[#71717a]">Fetching alice/3f9a2b...</span>
                    </div>
                    <div className="flex gap-2">
                      <span className="text-[#3f3f46]">›</span>
                      <span className="text-[#71717a]">Decrypting myapp/.env.enc</span>
                    </div>
                    <div className="flex gap-2">
                      <span className="text-[#3f3f46]">›</span>
                      <span className="text-[#71717a]">Decrypting myapp/.env.local.enc</span>
                    </div>
                    <div className="flex gap-2">
                      <span className="text-green-400">✓</span>
                      <span className="text-green-400">Restored 2 files</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div className="bg-[#2e1065] p-10 lg:p-14 flex flex-col justify-center">
              <span className="inline-block bg-white/10 rounded-full px-3 py-1 text-[11px] text-white/60 mb-6 tracking-wider uppercase w-fit">
                Recovery
              </span>
              <h3 className="text-[clamp(24px,3vw,34px)] font-bold text-white leading-tight mb-3">
                Recover<br />anywhere
              </h3>
              <p className="text-white/60 text-sm leading-relaxed max-w-xs">
                New machine, zero setup. Run{" "}
                <code className="font-mono bg-white/10 px-1.5 py-0.5 rounded text-white/80 text-[11px]">
                  envault recover
                </code>{" "}
                and your .env files are back in seconds. No git install required.
              </p>
            </div>
          </div>

        </div>
      </section>

      {/* ── Quick Start ──────────────────────────────────────── */}
      <section className="max-w-6xl mx-auto px-6 pb-24 border-t border-[#27272a] pt-20">
        <div className="text-center mb-20">
          <h2 className="text-4xl font-bold mb-4">Four commands. That&apos;s it.</h2>
          <p className="text-[#71717a] text-base">No configuration files. No setup wizards. Just simple CLI commands.</p>
        </div>

        <div className="space-y-20">

          {/* Step 01 — Login: text left, terminal right */}
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
            <div>
              <div className="flex items-center gap-5 mb-6">
                <span className="text-7xl font-bold text-[#1c1c1f] leading-none select-none">01</span>
                <div className="flex-1 h-px bg-[#27272a]" />
              </div>
              <h3 className="text-xl font-bold mb-2">Login</h3>
              <p className="text-[#71717a] text-sm leading-relaxed max-w-sm">
                Set your encryption passphrase and GitHub token. envault creates a private backup repo and saves credentials to <code className="font-mono text-[#a78bfa] text-xs">~/.envault/config.yaml</code>.
              </p>
            </div>
            <div className="terminal">
              <div className="terminal-chrome">
                <div className="w-3 h-3 rounded-full bg-[#ff5f57]" />
                <div className="w-3 h-3 rounded-full bg-[#febc2e]" />
                <div className="w-3 h-3 rounded-full bg-[#28c840]" />
                <span className="ml-2 font-mono text-xs text-[#71717a]">terminal</span>
              </div>
              <div className="terminal-body space-y-1 font-mono text-xs leading-relaxed">
                <div><span className="text-[#a78bfa]">$</span> <span className="text-[#fafafa]">envault login <span className="text-[#f97316]">-p</span> <span className="text-[#c4b5fd]">mysecret</span> <span className="text-[#f97316]">-t</span> <span className="text-[#c4b5fd]">ghp_xxxx</span></span></div>
                <div className="text-[#71717a]">› Generating repo name...</div>
                <div className="text-green-400">✓ Created private repo: alice/3f9a2b4d</div>
                <div className="text-green-400">✓ Config saved → ~/.envault/config.yaml</div>
              </div>
            </div>
          </div>

          {/* Step 02 — Init: terminal left, text right */}
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
            <div className="terminal lg:order-first order-last">
              <div className="terminal-chrome">
                <div className="w-3 h-3 rounded-full bg-[#ff5f57]" />
                <div className="w-3 h-3 rounded-full bg-[#febc2e]" />
                <div className="w-3 h-3 rounded-full bg-[#28c840]" />
                <span className="ml-2 font-mono text-xs text-[#71717a]">terminal</span>
              </div>
              <div className="terminal-body space-y-1 font-mono text-xs leading-relaxed">
                <div><span className="text-[#a78bfa]">$</span> <span className="text-[#fafafa]">envault init</span></div>
                <div className="text-[#71717a]">› Found project root: /projects/myapp</div>
                <div className="text-green-400">✓ Created envault.json</div>
              </div>
            </div>
            <div className="lg:order-last order-first">
              <div className="flex items-center gap-5 mb-6">
                <span className="text-7xl font-bold text-[#1c1c1f] leading-none select-none">02</span>
                <div className="flex-1 h-px bg-[#27272a]" />
              </div>
              <h3 className="text-xl font-bold mb-2">Init</h3>
              <p className="text-[#71717a] text-sm leading-relaxed max-w-sm">
                Run once in your project root. Creates <code className="font-mono text-[#a78bfa] text-xs">envault.json</code> to mark the directory as the backup source.
              </p>
            </div>
          </div>

          {/* Step 03 — Backup: text left, terminal right */}
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
            <div>
              <div className="flex items-center gap-5 mb-6">
                <span className="text-7xl font-bold text-[#1c1c1f] leading-none select-none">03</span>
                <div className="flex-1 h-px bg-[#27272a]" />
              </div>
              <h3 className="text-xl font-bold mb-2">Backup</h3>
              <p className="text-[#71717a] text-sm leading-relaxed max-w-sm">
                Finds all <code className="font-mono text-[#a78bfa] text-xs">.env*</code> files, encrypts each one with AES-256-GCM, and pushes them to your private GitHub repo. Plaintext never leaves your machine.
              </p>
            </div>
            <div className="terminal">
              <div className="terminal-chrome">
                <div className="w-3 h-3 rounded-full bg-[#ff5f57]" />
                <div className="w-3 h-3 rounded-full bg-[#febc2e]" />
                <div className="w-3 h-3 rounded-full bg-[#28c840]" />
                <span className="ml-2 font-mono text-xs text-[#71717a]">terminal</span>
              </div>
              <div className="terminal-body space-y-1 font-mono text-xs leading-relaxed">
                <div><span className="text-[#a78bfa]">$</span> <span className="text-[#fafafa]">envault backup</span></div>
                <div className="text-[#71717a]">› Scanning for .env files...</div>
                <div className="text-[#71717a] pl-4">✓ .env</div>
                <div className="text-[#71717a] pl-4">✓ .env.local</div>
                <div className="text-[#71717a] pl-4">✓ config/.env</div>
                <div className="text-green-400">✓ Encrypted 3 files</div>
                <div className="text-green-400">✓ Pushed to alice/3f9a2b4d</div>
              </div>
            </div>
          </div>

          {/* Step 04 — Recover: terminal left, text right */}
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
            <div className="terminal lg:order-first order-last">
              <div className="terminal-chrome">
                <div className="w-3 h-3 rounded-full bg-[#ff5f57]" />
                <div className="w-3 h-3 rounded-full bg-[#febc2e]" />
                <div className="w-3 h-3 rounded-full bg-[#28c840]" />
                <span className="ml-2 font-mono text-xs text-[#71717a]">terminal</span>
              </div>
              <div className="terminal-body space-y-1 font-mono text-xs leading-relaxed">
                <div><span className="text-[#a78bfa]">$</span> <span className="text-[#fafafa]">envault recover</span></div>
                <div className="text-[#71717a]">› Fetching alice/3f9a2b4d...</div>
                <div className="text-green-400">✓ Decrypted .env</div>
                <div className="text-green-400">✓ Decrypted .env.local</div>
                <div className="text-green-400">✓ Decrypted config/.env</div>
                <div className="text-green-400">✓ Restored 3 files</div>
              </div>
            </div>
            <div className="lg:order-last order-first">
              <div className="flex items-center gap-5 mb-6">
                <span className="text-7xl font-bold text-[#1c1c1f] leading-none select-none">04</span>
                <div className="flex-1 h-px bg-[#27272a]" />
              </div>
              <h3 className="text-xl font-bold mb-2">Recover</h3>
              <p className="text-[#71717a] text-sm leading-relaxed max-w-sm">
                On any new machine, run recover to pull and decrypt your .env files. No git install needed — just your passphrase.
              </p>
            </div>
          </div>

        </div>
      </section>

      {/* ── Footer ───────────────────────────────────────────── */}
      <footer className="border-t border-[#27272a]">
        <div className="max-w-6xl mx-auto px-6 py-8 flex items-center justify-between">
          <span className="font-semibold text-sm text-[#3f3f46] tracking-widest">ENVAULT</span>
          <a
            href="https://github.com/winnerx0/envault"
            target="_blank"
            rel="noopener noreferrer"
            className="font-mono text-xs text-[#71717a] hover:text-[#a78bfa] transition-colors"
          >
            github.com/winnerx0/envault ↗
          </a>
        </div>
      </footer>

    </div>
  );
}
