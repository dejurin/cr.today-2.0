/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class',
  content: ['../webapp/views/**/*.{html,jet,jet.html}'],
  theme: {
    extend: {},
  },
  plugins: ["prettier-plugin-tailwindcss"],
}


