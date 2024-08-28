import "./styles/index.css";

import Alpine from "alpinejs";

window.Alpine = Alpine;

Alpine.data("darkModeToggle", () => ({
  darkMode: window.matchMedia("(prefers-color-scheme: dark)").matches,

  toggleDarkMode() {
    this.darkMode = !this.darkMode;
    localStorage.setItem("dark-mode", this.darkMode ? "true" : "false");
    document.documentElement.classList.toggle("dark", this.darkMode);
  },

  init() {
    const storedPreference = localStorage.getItem("dark-mode");

    console.log(storedPreference);

    if (storedPreference !== null) {
      this.darkMode = storedPreference === "true";
    }

    document.documentElement.classList.toggle("dark", this.darkMode);
  },
}));

Alpine.start();
