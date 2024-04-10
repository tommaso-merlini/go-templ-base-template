/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./view/**/*.templ}", "./**/*.templ"],
    safelist: [],
    plugins: [require("daisyui")],
    daisyui: {
        themes: ["dark"]
    },
    theme: {
        extend: {
            fontFamily: {
                'work-sans': ['"Work Sans"', 'sans-serif'],
            },
        },
    },
}
