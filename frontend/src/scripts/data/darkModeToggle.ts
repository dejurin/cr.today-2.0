const darkModeToggle = () => ({
    darkModeSwitch: window.matchMedia("(prefers-color-scheme: dark)").matches,

    darkModeToggle() {
        this.darkModeSwitch = !this.darkModeSwitch;
        localStorage.setItem("dark-mode-switch", this.darkModeSwitch ? "true" : "false");
        document.documentElement.classList.toggle("dark", this.darkModeSwitch);
    },

    init() {
        const storedPreference = localStorage.getItem("dark-mode-switch");

        if (storedPreference !== null) {
            this.darkModeSwitch = storedPreference === "true";
        }

        document.documentElement.classList.toggle("dark", this.darkModeSwitch);
    },
})

export default darkModeToggle;