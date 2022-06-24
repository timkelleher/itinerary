import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(),
  mode: 'history',
  routes: [
        {
            name: "home",
            alias: ['/index.html'],
            path: '/',
            component: () => import('./src/views/CollectionParentPage.vue'),
            children: [
                {
                    name: "collection_list",
                    path: '/',
                    component: () => import('./src/views/CollectionListPage.vue'),
                },
                {
                    name: "collection_detail",
                    path: '/collection/:id',
                    component: () => import('./src/views/CollectionDetailPage.vue'),
                },
                {
                    name: "collection_add",
                    path: '/collection/add',
                    component: () => import('./src/views/CollectionAddPage.vue'),
                },
                {
                    name: "collection_edit",
                    path: '/collection/edit/:id',
                    component: () => import('./src/views/CollectionEditPage.vue'),
                },
                {
                    name: "endpoint_detail",
                    path: '/endpoint/view/:id',
                    component: () => import('./src/views/EndpointDetailPage.vue'),
                    props: { default: true, mode: "" },
                },
                {
                    name: "endpoint_add",
                    path: '/endpoint/add/:collectionID',
                    component: () => import('./src/views/EndpointAddPage.vue'),
                },
                {
                    name: "endpoint_edit",
                    path: '/endpoint/edit/:id',
                    component: () => import('./src/views/EndpointEditPage.vue'),
                },
            ],
        },
        {
            name: "error",
            path: '/error',
            component: () => import('./src/views/ErrorPage.vue'),
        },
    ]
});

export default router;