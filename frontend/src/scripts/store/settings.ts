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

export default settingsStore;