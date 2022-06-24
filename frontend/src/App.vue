<template>
    <div id="app" class="container-fluid">
        <div class="row">
            <div class="col-md-12 py-2">
                <router-link :to="{ name: 'collection_list' }">
                    <h1 class="text-center">Itinerary</h1>
                </router-link>
            </div>
        </div>
        <router-view />
    </div>
</template>
