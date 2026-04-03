import CopyButton from "./components/CopyButton";

const STEPS = [
  {
    num: "01",
    label: "Login",
    cmd: "envault login -p <password> -t <token>",
    desc: "Set your encryption passphrase and GitHub token. A private backup repo is created automatically and credentials are saved to ~/.envault/config.yaml.",
  },
  {
    num: "02",
    label: "Init",
    cmd: "envault init",
    desc: "Creates envault.json in your project root, marking it as the source of truth for all encrypted backups in that project.",
  },
  {
    num: "03",
    label: "Backup",
    cmd: "envault backup",
    desc: "All .env files are encrypted locally and pushed to your private GitHub repo via the Git Data API. Plaintext never leaves your machine.",
  },
  {
    num: "04",
    label: "Recover",
    cmd: "envault recover",
    desc: "On any machine, download and decrypt your .env files using your passphrase. Back in seconds, no git required.",
  },
];

const FEATURES = [
  {
    glyph: "❏",
    title: "AES-256-GCM",
    desc: "Military-grade authenticated encryption. Each file is independently encrypted.",
  },
  {
    glyph: "◈",
    title: "Argon2id KDF",
    desc: "Memory-hard key derivation makes brute force attacks computationally infeasible.",
  },
  {
    glyph: "⬡",
    title: "One repo, all projects",
    desc: "All encrypted backups land in one private repo, organized under each project's folder name.",
  },
  {
    glyph: "◎",
    title: "Zero plaintext upload",
    desc: "Encryption happens entirely on your machine. GitHub only ever sees ciphertext.",
  },
  {
    glyph: "◻",
    title: "Per-file unique salt",
    desc: "Every file gets a fresh 16-byte random salt and nonce. No two ciphertexts are alike.",
  },
  {
    glyph: "△",
    title: "Git Data API",
    desc: "Pushes directly via the GitHub API. No local git binary needed on the target machine.",
  },
];

const TERMINAL_LINES = [
  { type: "prompt", text: "envault login -p mysecret -t ghp_xxxx" },
  { type: "out", text: "Project root: myapp" },
  { type: "out", text: "Created private backup repo: alice/3f9a2b..." },
  { type: "out", text: "Config saved to ~/.envault/config.yaml" },
  { type: "prompt", text: "envault init" },
  { type: "out", text: "Initialized envault project in /projects/myapp" },
  { type: "prompt", text: "envault backup" },
  { type: "out", text: "Encrypted env files backed up to alice/3f9a2b..." },
];

const QUICK_START = `# 1. Login — set password, token, create private repo
envault login -p <password> -t <github_token>

# 2. Init — mark your project root
envault init

# 3. Backup — encrypt and push .env files
envault backup

# 4. Recover — pull and decrypt on any machine
envault recover`;

export default function Home() {
  return (
    <div className="min-h-screen relative z-10">

      {/* ── Nav ─────────────────────────────────────────────── */}
      <nav className="sticky top-0 z-50 border-b border-[#1c3d5c] bg-[#091520]/92 backdrop-blur-md">
        <div className="max-w-6xl mx-auto px-6 h-14 flex items-center justify-between">
          <span className="font-display text-base tracking-widest text-[#4db8ff]">
            ENVAULT
          </span>
          <a
            href="https://github.com/winnerx0/envault"
            target="_blank"
            rel="noopener noreferrer"
            className="font-mono text-xs text-[#5a7d99] hover:text-[#4db8ff] transition-colors"
          >
            github ↗
          </a>
        </div>
      </nav>

      {/* ── Hero ────────────────────────────────────────────── */}
      <section className="max-w-6xl mx-auto px-6 pt-20 pb-20">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-14 items-start">

          {/* Left */}
          <div>
            <div className="fade-up fade-up-1 inline-flex items-center gap-2 bg-[#0d1e2e] border border-[#1c3d5c] rounded-sm px-3 py-1 text-xs font-mono text-[#5a7d99] mb-8">
              <span className="w-1.5 h-1.5 rounded-full bg-[#4db8ff] inline-block animate-pulse" />
              AES-256-GCM · Argon2id · Open source
            </div>

            <h1 className="fade-up fade-up-2 font-display text-[clamp(40px,7.5vw,72px)] leading-[1.1] tracking-wide mb-6">
              SECURE<br />
              YOUR<br />
              <span className="text-[#4db8ff]">.ENV</span><br />
              FILES
            </h1>

            <p className="fade-up fade-up-3 text-[#5a7d99] text-lg leading-relaxed max-w-sm mb-10 font-sans">
              Encrypt and back up your{" "}
              <code className="font-mono text-sm text-[#c8dff0] bg-[#0d1e2e] px-1.5 py-0.5 rounded-sm border border-[#1c3d5c]">
                .env
              </code>{" "}
              files to a private GitHub repo. Your secrets never leave your machine in plaintext.
            </p>

            <div className="fade-up fade-up-4 flex flex-wrap gap-3">
              <code className="font-mono text-xs bg-[#0d1e2e] border border-[#1c3d5c] text-[#4db8ff] px-4 py-2.5 rounded-sm select-all">
                go install github.com/winnerx0/envault/cmd/envault@latest
              </code>
            </div>
          </div>

          {/* Right — animated terminal */}
          <div className="terminal lg:mt-2">
            <div className="terminal-chrome">
              <div className="w-3 h-3 rounded-full bg-[#ff5f57]" />
              <div className="w-3 h-3 rounded-full bg-[#febc2e]" />
              <div className="w-3 h-3 rounded-full bg-[#28c840]" />
              <span className="ml-2 font-mono text-xs text-[#5a7d99]">zsh — envault demo</span>
            </div>
            <div className="terminal-body space-y-0.5 min-h-[200px]">
              {TERMINAL_LINES.map((line, i) => (
                <div
                  key={i}
                  className={`term-line flex gap-2 ${
                    line.type === "out" ? "text-[#5a7d99]" : "text-[#c8dff0]"
                  }`}
                >
                  <span
                    className={`select-none shrink-0 ${
                      line.type === "prompt" ? "text-[#4db8ff]" : "text-[#1c3d5c]"
                    }`}
                  >
                    {line.type === "prompt" ? "$" : "›"}
                  </span>
                  <span>{line.text}</span>
                </div>
              ))}
              <div className="term-line flex gap-2">
                <span className="text-[#4db8ff] select-none">$</span>
                <span className="cursor-blink" />
              </div>
            </div>
          </div>

        </div>
      </section>

      {/* ── Divider ─────────────────────────────────────────── */}
      <div className="max-w-6xl mx-auto px-6">
        <div className="border-t border-[#1c3d5c]" />
      </div>

      {/* ── How it works ────────────────────────────────────── */}
      <section className="max-w-6xl mx-auto px-6 py-20">
        <p className="blueprint-label">How it works</p>
        <h2 className="font-display text-4xl tracking-wide mb-14 mt-3">
          FOUR COMMANDS
        </h2>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-px bg-[#1c3d5c]">
          {STEPS.map((step) => (
            <div
              key={step.num}
              className="bg-[#091520] p-6 hover:bg-[#0c1f30] transition-colors group"
            >
              <div className="font-display text-5xl text-[#122333] group-hover:text-[#4db8ff]/10 transition-colors mb-3 leading-none">
                {step.num}
              </div>
              <p className="font-mono text-[#4db8ff] text-xs uppercase tracking-widest mb-2">
                {step.label}
              </p>
              <code className="block font-mono text-xs text-[#c8dff0] bg-[#0d1e2e] border border-[#1c3d5c] px-3 py-2 rounded-sm mb-3 break-all leading-relaxed">
                $ {step.cmd}
              </code>
              <p className="text-[#5a7d99] text-sm leading-relaxed font-sans">
                {step.desc}
              </p>
            </div>
          ))}
        </div>
      </section>

      {/* ── Divider ─────────────────────────────────────────── */}
      <div className="max-w-6xl mx-auto px-6">
        <div className="border-t border-[#1c3d5c]" />
      </div>

      {/* ── Features ────────────────────────────────────────── */}
      <section className="max-w-6xl mx-auto px-6 py-20">
        <p className="blueprint-label">Features</p>
        <h2 className="font-display text-4xl tracking-wide mb-14 mt-3">
          BUILT SECURE
        </h2>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-px bg-[#1c3d5c]">
          {FEATURES.map((f, i) => (
            <div
              key={i}
              className="bg-[#091520] p-6 hover:bg-[#0c1f30] transition-colors group"
            >
              <span className="font-mono text-lg text-[#1c3d5c] group-hover:text-[#4db8ff] transition-colors block mb-3">
                {f.glyph}
              </span>
              <h3 className="font-mono text-sm text-[#c8dff0] mb-1.5">{f.title}</h3>
              <p className="text-[#5a7d99] text-sm leading-relaxed font-sans">{f.desc}</p>
            </div>
          ))}
        </div>
      </section>

      {/* ── Divider ─────────────────────────────────────────── */}
      <div className="max-w-6xl mx-auto px-6">
        <div className="border-t border-[#1c3d5c]" />
      </div>

      {/* ── Quick Start ─────────────────────────────────────── */}
      <section className="max-w-6xl mx-auto px-6 py-20">
        <p className="blueprint-label">Quick Start</p>
        <h2 className="font-display text-4xl tracking-wide mb-10 mt-3">
          GET RUNNING
        </h2>

        <div className="terminal max-w-2xl">
          <div className="terminal-chrome justify-between">
            <div className="flex items-center gap-2">
              <div className="w-3 h-3 rounded-full bg-[#ff5f57]" />
              <div className="w-3 h-3 rounded-full bg-[#febc2e]" />
              <div className="w-3 h-3 rounded-full bg-[#28c840]" />
            </div>
            <CopyButton text={QUICK_START} />
          </div>
          <div className="terminal-body">
            <pre className="text-sm whitespace-pre-wrap leading-relaxed">
              <span className="text-[#1c3d5c]"># 1. Login — set password, token, create private repo{"\n"}</span>
              <span className="text-[#c8dff0]">
                <span className="text-[#4db8ff]">$</span>{" "}
                envault login{" "}
                <span className="text-[#ff7043]">-p</span>{" "}
                <span className="text-[#7ec8e3]">&lt;password&gt;</span>{" "}
                <span className="text-[#ff7043]">-t</span>{" "}
                <span className="text-[#7ec8e3]">&lt;github_token&gt;</span>
              </span>{"\n\n"}
              <span className="text-[#1c3d5c]"># 2. Init — mark your project root{"\n"}</span>
              <span className="text-[#c8dff0]"><span className="text-[#4db8ff]">$</span> envault init</span>{"\n\n"}
              <span className="text-[#1c3d5c]"># 3. Backup — encrypt and push .env files{"\n"}</span>
              <span className="text-[#c8dff0]"><span className="text-[#4db8ff]">$</span> envault backup</span>{"\n\n"}
              <span className="text-[#1c3d5c]"># 4. Recover — pull and decrypt on any machine{"\n"}</span>
              <span className="text-[#c8dff0]"><span className="text-[#4db8ff]">$</span> envault recover</span>
            </pre>
          </div>
        </div>
      </section>

      {/* ── Divider ─────────────────────────────────────────── */}
      <div className="max-w-6xl mx-auto px-6">
        <div className="border-t border-[#1c3d5c]" />
      </div>

      {/* ── Coming Soon ─────────────────────────────────────── */}
      <section className="max-w-6xl mx-auto px-6 py-20">
        <p className="blueprint-label">Coming Soon</p>
        <h2 className="font-display text-4xl tracking-wide mb-10 mt-3">
          WHAT&apos;S NEXT
        </h2>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 max-w-2xl">
          <div className="border border-[#1c3d5c] rounded-sm p-5 flex items-start gap-3 hover:border-[#2a5878] hover:bg-[#0c1f30] transition-colors">
            <span className="font-mono text-xs text-[#1c3d5c] mt-0.5 shrink-0">[ ]</span>
            <div>
              <p className="font-mono text-sm text-[#c8dff0] mb-1.5">Selective backup</p>
              <p className="text-[#5a7d99] text-xs leading-relaxed font-sans">
                <code className="text-[#4db8ff]/60 font-mono">envault backup .env.local</code>
                {" "}— target specific files instead of all{" "}
                <code className="font-mono text-[#5a7d99]">.env*</code>
              </p>
            </div>
          </div>
          <div className="border border-[#1c3d5c] rounded-sm p-5 flex items-start gap-3 hover:border-[#2a5878] hover:bg-[#0c1f30] transition-colors">
            <span className="font-mono text-xs text-[#1c3d5c] mt-0.5 shrink-0">[ ]</span>
            <div>
              <p className="font-mono text-sm text-[#c8dff0] mb-1.5">Selective recover</p>
              <p className="text-[#5a7d99] text-xs leading-relaxed font-sans">
                <code className="text-[#4db8ff]/60 font-mono">envault recover .env.production</code>
                {" "}— restore a single file on demand
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* ── Footer ──────────────────────────────────────────── */}
      <footer className="border-t border-[#1c3d5c]">
        <div className="max-w-6xl mx-auto px-6 py-8 flex items-center justify-between">
          <span className="font-display text-sm text-[#122333] tracking-widest">
            ENVAULT
          </span>
          <a
            href="https://github.com/winnerx0/envault"
            target="_blank"
            rel="noopener noreferrer"
            className="font-mono text-xs text-[#5a7d99] hover:text-[#4db8ff] transition-colors"
          >
            github.com/winnerx0/envault ↗
          </a>
        </div>
      </footer>

    </div>
  );
}
