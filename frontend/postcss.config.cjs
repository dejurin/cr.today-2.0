module.exports = {
    plugins: {
      autoprefixer: {},
      "postcss-import": {},
      "tailwindcss/nesting": {},
      tailwindcss: {},
      ...(process.env.NODE_ENV === "production" ? { cssnano: {} } : {}),
    },
  };
  