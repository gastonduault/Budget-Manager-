import {
    createRouter,
    createWebHistory
} from 'vue-router'
// import DisplayBudget from '../views/DisplayBudget.vue'
import Login from '@/components/Login.vue'
import DisplayBudget from '@/views/DisplayBudget.vue'

const routes = [{
        path: '/budget',
        name: 'budget',
        component: DisplayBudget
    },
    {
        path: '/',
        name: 'login',
        component: Login
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router