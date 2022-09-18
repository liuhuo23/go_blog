module.exports = {
  pluginOptions: {
    'style-resources-loader': {
      preProcessor: 'sass',
      patterns: [

      ]
    }
  },
  css: {
    loaderOptions: {
      sass: {
        data: `
          @import "@/assets/scss/style.scss"
        `
      }
    }
  }
}
