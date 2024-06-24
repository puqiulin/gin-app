/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{js,ts,jsx,tsx,mdx}"],
    theme: {
        extend: {
            backgroundImage: {
                'hot_air_balloon': "url('./public/images/hot_air_balloon.png')",
            },
            colors: {
                primary: "#345",
                secondary: {
                    100: "#564",
                    200: "#567",
                }
            },
            fontFamily: {
                poetsen: ["Poetsen One"]
            }
        },
    },
    plugins: [require('daisyui'),],
    daisyui: {
        themes: ["light", "dark", "cupcake"],
    },
}

