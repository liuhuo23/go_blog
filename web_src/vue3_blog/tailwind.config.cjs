/** @type {import('tailwindcss').Config} */
module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  content: [],
  theme: {
    extend: {},
    maxWidth: {
      '1/10': '10%',
      '3/10': '30%',
      '4/10': '40%',
      '5/10': '50%',
    },
    minWidth: {
      '1/10': '10%',
      '2/10': '20%',
      '3/10': '30%',
      '4/10': '40%',
      '1/24': '4.1666666%',
      '0': '0',
      '1/4': '25%',
      '1/2': '50%',
      '3/4': '75%',
      'full': '100%',
      '4/5': '80%',
      '35/100': '35%',
      '30/100': '30%'
    }
  },
  variants: {
    extend: {
      backgroundColor: ['active'],
      border: ['active'],
    }
  },
  plugins: [],
}
