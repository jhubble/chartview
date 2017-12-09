module.exports = {
    files: {
        javascripts: {
            joinTo: {
                'vendor.js': /^(?!app)/, // Files that are not in `app` dir.
                'app.js': /^app/
            }
        },
        stylesheets: {
            joinTo: {
                'css/app.css': /^app/
            }
        },
        templates: {
            joinTo: 'app.js'
        }
    },
    plugins: {
        babel: {
            presets: ['latest', 'stage-3']
        },
        vue: {
            extractCSS: true,
            out: 'public/css/components.css'
        },
        copyfilemon:{
            'css': 'node_modules/vuetify/dist/vuetify.css'
        }
    },
    npm: {
        aliases: {
            'vue': 'vue/dist/vue.common.js'
        }
    }
}
