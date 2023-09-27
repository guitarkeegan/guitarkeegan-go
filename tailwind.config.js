/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.tmpl"],
  theme: {
    backgroundImage: {
      'home': "url('/images/trebleOnLeftGrass-fs8.png')",
    },
      "colors": {
       "brand": {
        "100": "#c8e7c3",
        "200": "#44e929",
        "300": "#77fe61",
        "400": "#158362",
        "500": "#1b4538"
       },
       "Grays": {
        "100": "#cbe1c9",
        "200": "#818e80",
        "300": "#606a5f",
        "400": "#3f483e",
        "500": "#1a1e19"
       },
       "greens": {
        "700": "#124708",
        "600": "#124708",
        "500": "#158362",
        "400": "#44e929",
        "300": "#77fe61",
        "200": "#c8e7c3",
        "100": "#def3ed"
       },
      "fontSize": {
       "base": "0.75rem",
       "lg": "1rem",
       "xl": "1.5rem",
       "2xl": "2.5rem",
       "3xl": "4rem"
      },
      "fontFamily": {
       "serif": "Cormorant",
       "sans": "Roboto"
      },
      "borderRadius": {
       "none": "0",
       "xs": "0.3125rem"
      }
     }
  },
  plugins: [
    require("tailwindcss-3d"),
  ],
}

