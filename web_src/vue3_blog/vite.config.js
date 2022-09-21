const path = require('path')
module.exports ={
    alias:{
        '/@/': path.resolve(__dirname, './src')
    },
    hostname: 'localhost', //默认是localhost
    port:'8000', //默认是3000端口
    open:false,
    https:false,
    ssr:false,
    base: './', //生产环境下的公共路径
    outDir: 'dist', 
    proxy: {
        '/api': {
            target: 'http://localhost:8080', //后端服务实际地址
            changeOrigin: true,
            rewrite: path => path.replace(/^\/api/, '')
        }
    },
    plugins:[
        
    ]
}