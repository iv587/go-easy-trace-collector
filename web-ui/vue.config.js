module.exports = {
    devServer: {
        proxy: {
            '^/api': {
                target: 'http://log.hoolihome.com',
                ws: true,
                changeOrigin: true
            },
        }
    }
}
