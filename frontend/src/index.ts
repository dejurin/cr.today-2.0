import "./styles/index.css";
import settingsStore from "./scripts/store/settings"
import darkModeToggle from "./scripts/data/darkModeToggle"
import currencyDropdown from "./scripts/data/currencyDropdown"

import Alpine from "alpinejs";

// init

window.Alpine = Alpine;

// store

Alpine.store("settings", settingsStore);

// data

Alpine.data("currencyDropdown", currencyDropdown);
Alpine.data("darkModeToggle", darkModeToggle);

// start

Alpine.start();
