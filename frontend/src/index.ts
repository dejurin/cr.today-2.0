import "./styles/index.css";

import Alpine from "alpinejs";
import mask from '@alpinejs/mask';
 
Alpine.plugin(mask)

window.Alpine = Alpine;

const settingsStore = {
  amount: localStorage.getItem("amount") || "1",
  base: localStorage.getItem("base") || "USD",
  quote: localStorage.getItem("quote") || "EUR",
  timeZone: localStorage.getItem("timeZone") || "UTC",
  decimals: parseInt(localStorage.getItem("decimals") || "2", 10),


  setAmount(value: string) {
    this.amount = value;
    localStorage.setItem("amount", value);
  },

  setBase(value: string) {
    this.base = value;
    localStorage.setItem("base", value);
  },

  setQuote(value: string) {
    this.quote = value;
    localStorage.setItem("quote", value);
  },

  setTimeZone(value: string) {
    this.timeZone = value;
    localStorage.setItem("timeZone", value);
  },

  setDecimals(value: number) {
    this.decimals = value;
    localStorage.setItem("decimals", value.toString());
  },
};

Alpine.store("settings", settingsStore);

Alpine.data("darkModeToggle", () => ({
  darkModeSwitch: window.matchMedia("(prefers-color-scheme: dark)").matches,

  darkModeToggle() {
    this.darkModeSwitch = !this.darkModeSwitch;
    localStorage.setItem("dark-mode-switch", this.darkModeSwitch ? "true" : "false");
    document.documentElement.classList.toggle("dark", this.darkModeSwitch);
  },

  init() {
    const storedPreference = localStorage.getItem("dark-mode-switch");

    console.log(storedPreference);

    if (storedPreference !== null) {
      this.darkModeSwitch = storedPreference === "true";
    }

    document.documentElement.classList.toggle("dark", this.darkModeSwitch);
  },
}));

Alpine.start();
