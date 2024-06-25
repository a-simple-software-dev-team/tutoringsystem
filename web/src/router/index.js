import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from "../components/Login.vue"
import register from "../components/register.vue"
import tutors from "../components/tutors.vue"
import students_profile from "../components/students_profile.vue"
import tutors_schedule from "../components/tutors_schedule.vue"
import tutors_profile from "../components/tutors_profile.vue"
import tutors_certification from "../components/tutors_certification.vue"
import matching from '@/components/matching.vue'
import matching_student from "../components/matching_student.vue"
import notifications from "../components/notifications.vue"
import student from "../components/student.vue"
import notifications_student from '../components/notifications_student.vue'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/home',
    name: 'home',
    component: HomeView
  },
  
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/register',
    name: 'register',
    component: register
  },
  {
    path: '/tutors',
    name: 'tutors',
    component: tutors,
    children:[{
      
        path: '/tutors/schedule',
        component: tutors_schedule
      
     },
     {
      
      path: '/tutors/profile',
      component: tutors_profile
    
   },  
   {
      
    path: '/notifications',
    component: notifications
  
    },  
   {
      
    path: '/tutors/certification',
    component: tutors_certification
  
 },
 {
      
  path: '/matching',
  component: matching

},
   
    ]
  },

  {
    path: '/student',
    name: 'student',
    component: student,
    children:[
   {
    
    path: '/students/profile',
    component: students_profile
  
 },  
 {
    
  path: '/notifications_student',
  component: notifications_student

},  

{
    
path: '/matching_student',
component: matching_student

},
 
  ]

  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
