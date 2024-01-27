/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./**/*.{ts,tsx}"],
  theme: {
    extend: {
      colors: {
        "main-50": "#f8fafc",
        "main-100": "#f1f5f9",
        "main-200": "#e2e8f0",
        "main-300": "#cbd5e1",
        "main-400": "#94a3b8",
        "main-500": "#64748b",
        "main-600": "#475569",
        "main-700": "#334155",
        "main-800": "#1E2429",
        "main-900": "#0f172a",
        "main-950": "#171B20",

        "primary-50": "#f5f3ff",
        "primary-100": "#ede9fe",
        "primary-200": "#ddd6fe",
        "primary-300": "#c4b5fd",
        "primary-400": "#a78bfa",
        "primary-500": "#8b5cf6",
        "primary-600": "#7c3aed",
        "primary-700": "#6d28d9",
        "primary-800": "#5b21b6",
        "primary-900": "#4c1d95",
        "primary-950": "#2e1065",

        "danger-50": "#fff1f2",
        "danger-100": "#ffe4e6",
        "danger-200": "#fecdd3",
        "danger-300": "#fda4af",
        "danger-400": "#fb7185",
        "danger-500": "#f43f5e",
        "danger-600": "#e11d48",
        "danger-700": "#be123c",
        "danger-800": "#9f1239",
        "danger-900": "#881337",
        "danger-950": "#4c0519",

        "success-050": "#ecfdf5",
        "success-100": "#d1fae5",
        "success-200": "#a7f3d0",
        "success-300": "#6ee7b7",
        "success-400": "#34d399",
        "success-500": "#10b981",
        "success-600": "#059669",
        "success-700": "#047857",
        "success-800": "#065f46",
        "success-900": "#064e3b",
        "success-950": "#022c22",
      },
    },
  },
  plugins: [require("tailwind-scrollbar")],
};
