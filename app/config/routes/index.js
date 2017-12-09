/**
 * Import your routes from different modules here. Eg:
 *
 * import authRoutes from auth
 */
import error from '../../components/error.vue'
import chart from '../../components/Chart.vue'
import home from '../../components/Home.vue'

let routes = [
    {
        path: '/chart',
        name: 'chart',
        component: chart
    },
    {
        path: '/',
        name: 'home',
        component: home
    },
    {
        path: '*',
        component: error,
        name: 'error'

    }
]

/**
 * Add your sub-routes here. One way to do so is:
 *
 * routes = [...routes, ...authRoutes]
 */

export default routes
