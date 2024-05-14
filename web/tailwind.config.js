/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{js,ts,jsx,tsx,mdx}"],
    theme: {
        extend: {
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
    plugins: [],
}

